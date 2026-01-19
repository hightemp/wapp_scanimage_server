package models

import "time"

// FileInfo represents information about a scanned file
type FileInfo struct {
	Name         string    `json:"name"`
	Path         string    `json:"path"`
	RelativePath string    `json:"relativePath"`
	Size         int64     `json:"size"`
	SizeHuman    string    `json:"sizeHuman"`
	ModTime      time.Time `json:"modTime"`
	Order        int       `json:"order"`
}

// ScannerInfo represents information about a scanner device
type ScannerInfo struct {
	Device      string `json:"device"`
	Vendor      string `json:"vendor"`
	Model       string `json:"model"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// ScanSettings represents scan settings
type ScanSettings struct {
	Resolution int    `json:"resolution"`
	Quality    int    `json:"quality"`
	Format     string `json:"format"`
	Device     string `json:"device"`
}

// BatchOperation represents a batch operation request
type BatchOperation struct {
	Action string   `json:"action"`
	Files  []string `json:"files"`
}

// ConvertRequest represents a PDF conversion request
type ConvertRequest struct {
	Files      []string `json:"files"`
	OutputName string   `json:"outputName,omitempty"`
}

// RenameRequest represents a file rename request
type RenameRequest struct {
	OldName string `json:"oldName"`
	NewName string `json:"newName"`
}

// Settings represents application settings stored in DB
type Settings struct {
	ID        int64  `json:"id"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

// ScanOrder represents file ordering information
type ScanOrder struct {
	ID       int64  `json:"id"`
	FileName string `json:"fileName"`
	Position int    `json:"position"`
}

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// ProgressInfo represents operation progress
type ProgressInfo struct {
	Total    int     `json:"total"`
	Current  int     `json:"current"`
	Percent  float64 `json:"percent"`
	Status   string  `json:"status"`
	FileName string  `json:"fileName,omitempty"`
}
