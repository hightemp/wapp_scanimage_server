package services

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/hightemp/wapp_scanimage_server/internal/config"
	"github.com/hightemp/wapp_scanimage_server/internal/models"
)

type ConverterService struct {
	cfg *config.Config
}

func NewConverterService(cfg *config.Config) *ConverterService {
	return &ConverterService{cfg: cfg}
}

// ConvertToPDF converts images to a PDF file using img2pdf
func (c *ConverterService) ConvertToPDF(fileNames []string, outputName string) (string, error) {
	if len(fileNames) == 0 {
		return "", fmt.Errorf("no files to convert")
	}

	// Generate output name if not provided
	if outputName == "" {
		outputName = fmt.Sprintf("%d.pdf", time.Now().Unix())
	}

	// Ensure .pdf extension
	if filepath.Ext(outputName) != ".pdf" {
		outputName += ".pdf"
	}

	outputPath := filepath.Join(c.cfg.PDFPath, outputName)

	// Build list of full paths
	filePaths := make([]string, len(fileNames))
	for i, name := range fileNames {
		filePaths[i] = filepath.Join(c.cfg.ScannedPath, filepath.Base(name))

		// Verify file exists
		if _, err := os.Stat(filePaths[i]); os.IsNotExist(err) {
			return "", fmt.Errorf("file not found: %s", name)
		}
	}

	// Build command: img2pdf file1 file2 ... -o output.pdf
	args := append(filePaths, "-o", outputPath)
	cmd := exec.Command("img2pdf", args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("img2pdf failed: %w, output: %s", err, string(output))
	}

	return outputName, nil
}

// ConvertToPDFWithProgress converts images to PDF with progress reporting
func (c *ConverterService) ConvertToPDFWithProgress(fileNames []string, outputName string, progress chan<- models.ProgressInfo) (string, error) {
	total := len(fileNames)
	if total == 0 {
		return "", fmt.Errorf("no files to convert")
	}

	// Report starting
	progress <- models.ProgressInfo{
		Total:   total,
		Current: 0,
		Percent: 0,
		Status:  "preparing",
	}

	// Generate output name if not provided
	if outputName == "" {
		outputName = fmt.Sprintf("%d.pdf", time.Now().Unix())
	}

	if filepath.Ext(outputName) != ".pdf" {
		outputName += ".pdf"
	}

	outputPath := filepath.Join(c.cfg.PDFPath, outputName)

	// Build list of full paths
	filePaths := make([]string, len(fileNames))
	for i, name := range fileNames {
		filePaths[i] = filepath.Join(c.cfg.ScannedPath, filepath.Base(name))

		if _, err := os.Stat(filePaths[i]); os.IsNotExist(err) {
			return "", fmt.Errorf("file not found: %s", name)
		}

		progress <- models.ProgressInfo{
			Total:    total,
			Current:  i + 1,
			Percent:  float64(i+1) / float64(total) * 50, // First 50% is preparation
			Status:   "preparing",
			FileName: name,
		}
	}

	// Report converting
	progress <- models.ProgressInfo{
		Total:   total,
		Current: total,
		Percent: 50,
		Status:  "converting",
	}

	// Execute img2pdf
	args := append(filePaths, "-o", outputPath)
	cmd := exec.Command("img2pdf", args...)

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("img2pdf failed: %w", err)
	}

	// Report completed
	progress <- models.ProgressInfo{
		Total:    total,
		Current:  total,
		Percent:  100,
		Status:   "completed",
		FileName: outputName,
	}

	return outputName, nil
}

// CreateArchive creates a ZIP archive from the specified files
func (c *ConverterService) CreateArchive(fileNames []string, outputName string) (string, error) {
	if len(fileNames) == 0 {
		return "", fmt.Errorf("no files to archive")
	}

	// Generate output name if not provided
	if outputName == "" {
		outputName = fmt.Sprintf("%d.zip", time.Now().Unix())
	}

	// Ensure .zip extension
	if filepath.Ext(outputName) != ".zip" {
		outputName += ".zip"
	}

	outputPath := filepath.Join(c.cfg.ArchivesPath, outputName)

	// Create the archive
	archive, err := os.Create(outputPath)
	if err != nil {
		return "", fmt.Errorf("failed to create archive: %w", err)
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	for _, name := range fileNames {
		filePath := filepath.Join(c.cfg.ScannedPath, filepath.Base(name))

		if err := c.addFileToArchive(zipWriter, filePath, filepath.Base(name)); err != nil {
			return "", fmt.Errorf("failed to add %s to archive: %w", name, err)
		}
	}

	return outputName, nil
}

// CreateArchiveWithProgress creates a ZIP archive with progress reporting
func (c *ConverterService) CreateArchiveWithProgress(fileNames []string, outputName string, progress chan<- models.ProgressInfo) (string, error) {
	total := len(fileNames)
	if total == 0 {
		return "", fmt.Errorf("no files to archive")
	}

	if outputName == "" {
		outputName = fmt.Sprintf("%d.zip", time.Now().Unix())
	}

	if filepath.Ext(outputName) != ".zip" {
		outputName += ".zip"
	}

	outputPath := filepath.Join(c.cfg.ArchivesPath, outputName)

	archive, err := os.Create(outputPath)
	if err != nil {
		return "", fmt.Errorf("failed to create archive: %w", err)
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	for i, name := range fileNames {
		filePath := filepath.Join(c.cfg.ScannedPath, filepath.Base(name))

		progress <- models.ProgressInfo{
			Total:    total,
			Current:  i + 1,
			Percent:  float64(i+1) / float64(total) * 100,
			Status:   "archiving",
			FileName: name,
		}

		if err := c.addFileToArchive(zipWriter, filePath, filepath.Base(name)); err != nil {
			return "", fmt.Errorf("failed to add %s to archive: %w", name, err)
		}
	}

	progress <- models.ProgressInfo{
		Total:    total,
		Current:  total,
		Percent:  100,
		Status:   "completed",
		FileName: outputName,
	}

	return outputName, nil
}

// addFileToArchive adds a single file to a zip archive
func (c *ConverterService) addFileToArchive(zipWriter *zip.Writer, filePath, archiveName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = archiveName
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	return err
}

// CreateArchiveFromAll creates an archive from all scanned files
func (c *ConverterService) CreateArchiveFromAll() (string, error) {
	pattern := filepath.Join(c.cfg.ScannedPath, c.cfg.ScannedFileMask)
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return "", err
	}

	fileNames := make([]string, len(matches))
	for i, path := range matches {
		fileNames[i] = filepath.Base(path)
	}

	return c.CreateArchive(fileNames, "")
}

// ConvertAllToPDF converts all scanned files to a single PDF
func (c *ConverterService) ConvertAllToPDF() (string, error) {
	pattern := filepath.Join(c.cfg.ScannedPath, c.cfg.ScannedFileMask)
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return "", err
	}

	fileNames := make([]string, len(matches))
	for i, path := range matches {
		fileNames[i] = filepath.Base(path)
	}

	return c.ConvertToPDF(fileNames, "")
}
