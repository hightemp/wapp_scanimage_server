package services

import (
	"bufio"
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hightemp/wapp_scanimage_server/internal/config"
	"github.com/hightemp/wapp_scanimage_server/internal/models"
)

type ScannerService struct {
	cfg *config.Config
}

func NewScannerService(cfg *config.Config) *ScannerService {
	return &ScannerService{cfg: cfg}
}

// GetScanners returns a list of available scanners
func (s *ScannerService) GetScanners() ([]models.ScannerInfo, error) {
	cmd := exec.Command("scanimage", "-L")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list scanners: %w", err)
	}

	return s.parseScannerList(string(output)), nil
}

// GetScannersRaw returns raw scanner list output
func (s *ScannerService) GetScannersRaw() (string, error) {
	cmd := exec.Command("scanimage", "-L")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to list scanners: %w", err)
	}
	return string(output), nil
}

// GetScannersCached returns cached scanner list
func (s *ScannerService) GetScannersCached() ([]models.ScannerInfo, error) {
	// Cache path would be: filepath.Join(s.cfg.CachePath, "scanners_list.txt")
	// For simplicity, we'll just get fresh data each time
	// In production, you'd want proper cache invalidation with time-based expiry

	return s.GetScanners()
}

// Scan performs a scan with the given settings
func (s *ScannerService) Scan(settings models.ScanSettings) (string, error) {
	timestamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d.jpeg", timestamp)
	outputPath := filepath.Join(s.cfg.ScannedPath, fileName)

	args := []string{
		"--format=jpeg",
		fmt.Sprintf("--output-file=%s", outputPath),
		fmt.Sprintf("--resolution=%d", settings.Resolution),
	}

	// Add device if specified
	if settings.Device != "" {
		args = append(args, fmt.Sprintf("--device-name=%s", settings.Device))
	}

	cmd := exec.Command("scanimage", args...)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("scan failed: %w", err)
	}

	return fileName, nil
}

// ScanAsync performs a scan asynchronously and reports progress via channel
func (s *ScannerService) ScanAsync(settings models.ScanSettings, progress chan<- models.ProgressInfo) (string, error) {
	timestamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d.jpeg", timestamp)
	outputPath := filepath.Join(s.cfg.ScannedPath, fileName)

	args := []string{
		"--format=jpeg",
		fmt.Sprintf("--output-file=%s", outputPath),
		fmt.Sprintf("--resolution=%d", settings.Resolution),
		"--progress",
	}

	if settings.Device != "" {
		args = append(args, fmt.Sprintf("--device-name=%s", settings.Device))
	}

	cmd := exec.Command("scanimage", args...)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", fmt.Errorf("failed to get stderr pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("failed to start scan: %w", err)
	}

	// Parse progress from stderr
	go func() {
		scanner := bufio.NewScanner(stderr)
		re := regexp.MustCompile(`(\d+)%`)
		for scanner.Scan() {
			line := scanner.Text()
			if matches := re.FindStringSubmatch(line); len(matches) > 1 {
				if percent, err := strconv.Atoi(matches[1]); err == nil {
					progress <- models.ProgressInfo{
						Total:    100,
						Current:  percent,
						Percent:  float64(percent),
						Status:   "scanning",
						FileName: fileName,
					}
				}
			}
		}
	}()

	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("scan failed: %w", err)
	}

	progress <- models.ProgressInfo{
		Total:    100,
		Current:  100,
		Percent:  100,
		Status:   "completed",
		FileName: fileName,
	}

	return fileName, nil
}

// parseScannerList parses the output of scanimage -L
func (s *ScannerService) parseScannerList(output string) []models.ScannerInfo {
	var scanners []models.ScannerInfo

	// Example line: device `escl:https://192.168.31.200:443' is a Canon MF110/910 platen scanner
	re := regexp.MustCompile(`device\s+` + "`" + `([^']+)'\s+is\s+a\s+(.+)`)

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		matches := re.FindStringSubmatch(line)
		if len(matches) >= 3 {
			device := matches[1]
			description := matches[2]

			// Try to parse vendor/model from description
			parts := strings.SplitN(description, " ", 2)
			vendor := ""
			model := description
			scanType := ""

			if len(parts) >= 2 {
				vendor = parts[0]
				model = parts[1]
			}

			// Extract scanner type if present
			if idx := strings.LastIndex(model, " "); idx != -1 {
				possibleType := model[idx+1:]
				if possibleType == "scanner" || possibleType == "flatbed" || possibleType == "platen" {
					scanType = possibleType
					model = strings.TrimSuffix(model, " "+possibleType)
					model = strings.TrimSuffix(model, " scanner")
				}
			}

			scanners = append(scanners, models.ScannerInfo{
				Device:      device,
				Vendor:      vendor,
				Model:       model,
				Type:        scanType,
				Description: description,
			})
		} else {
			// Fallback: just add the whole line as description
			scanners = append(scanners, models.ScannerInfo{
				Description: line,
			})
		}
	}

	return scanners
}
