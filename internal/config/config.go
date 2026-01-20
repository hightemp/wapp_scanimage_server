package config

import (
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	// Server settings
	ServerHost string
	ServerPort int
	Debug      bool

	// Paths
	RootPath     string
	DataPath     string
	ScannedPath  string
	ArchivesPath string
	PDFPath      string
	CachePath    string

	// Relative paths for URLs
	ScannedRelPath  string
	ArchivesRelPath string
	PDFRelPath      string

	// File masks
	ScannedFileMask  string
	ArchivesFileMask string
	PDFFileMask      string
}

func New() *Config {
	rootPath, _ := os.Getwd()

	// Default paths based on rootPath
	defaultDataPath := filepath.Join(rootPath, "data")
	defaultScannedPath := filepath.Join(rootPath, "files", "scanned")
	defaultArchivesPath := filepath.Join(rootPath, "files", "archives")
	defaultPDFPath := filepath.Join(rootPath, "files", "pdf")
	defaultCachePath := filepath.Join(rootPath, "cache")

	cfg := &Config{
		ServerHost: getEnv("SERVER_HOST", "0.0.0.0"),
		ServerPort: getEnvInt("SERVER_PORT", 8080),
		Debug:      getEnvBool("DEBUG", false),

		RootPath:     rootPath,
		DataPath:     getEnv("DATA_PATH", defaultDataPath),
		ScannedPath:  getEnv("SCANNED_PATH", defaultScannedPath),
		ArchivesPath: getEnv("ARCHIVES_PATH", defaultArchivesPath),
		PDFPath:      getEnv("PDF_PATH", defaultPDFPath),
		CachePath:    getEnv("CACHE_PATH", defaultCachePath),

		ScannedRelPath:  "/files/scanned",
		ArchivesRelPath: "/files/archives",
		PDFRelPath:      "/files/pdf",

		ScannedFileMask:  "*.jpeg",
		ArchivesFileMask: "*.zip",
		PDFFileMask:      "*.pdf",
	}

	// Ensure directories exist
	os.MkdirAll(cfg.DataPath, 0755)
	os.MkdirAll(cfg.ScannedPath, 0755)
	os.MkdirAll(cfg.ArchivesPath, 0755)
	os.MkdirAll(cfg.PDFPath, 0755)
	os.MkdirAll(cfg.CachePath, 0755)

	return cfg
}

func (c *Config) DatabasePath() string {
	return filepath.Join(c.DataPath, "dbfile.db")
}

func (c *Config) Address() string {
	return c.ServerHost + ":" + strconv.Itoa(c.ServerPort)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}
