package services

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/hightemp/wapp_scanimage_server/internal/config"
	"github.com/hightemp/wapp_scanimage_server/internal/database"
	"github.com/hightemp/wapp_scanimage_server/internal/models"
)

type FileManagerService struct {
	cfg *config.Config
	db  *database.DB
}

func NewFileManagerService(cfg *config.Config, db *database.DB) *FileManagerService {
	return &FileManagerService{cfg: cfg, db: db}
}

// GetScannedFiles returns list of scanned files
func (f *FileManagerService) GetScannedFiles() ([]models.FileInfo, error) {
	return f.getFiles(f.cfg.ScannedPath, f.cfg.ScannedFileMask, f.cfg.ScannedRelPath)
}

// GetArchives returns list of archive files
func (f *FileManagerService) GetArchives() ([]models.FileInfo, error) {
	return f.getFiles(f.cfg.ArchivesPath, f.cfg.ArchivesFileMask, f.cfg.ArchivesRelPath)
}

// GetPDFs returns list of PDF files
func (f *FileManagerService) GetPDFs() ([]models.FileInfo, error) {
	return f.getFiles(f.cfg.PDFPath, f.cfg.PDFFileMask, f.cfg.PDFRelPath)
}

// getFiles retrieves files matching the pattern from the specified directory
func (f *FileManagerService) getFiles(dir, pattern, relPath string) ([]models.FileInfo, error) {
	fullPattern := filepath.Join(dir, pattern)
	matches, err := filepath.Glob(fullPattern)
	if err != nil {
		return nil, fmt.Errorf("failed to glob files: %w", err)
	}

	// Get file orders from database
	orders, _ := f.db.GetAllFileOrders()

	files := make([]models.FileInfo, 0, len(matches))
	for _, path := range matches {
		info, err := os.Stat(path)
		if err != nil {
			continue
		}

		fileName := filepath.Base(path)
		order := -1
		if pos, ok := orders[fileName]; ok {
			order = pos
		}

		files = append(files, models.FileInfo{
			Name:         fileName,
			Path:         path,
			RelativePath: relPath + "/" + fileName,
			Size:         info.Size(),
			SizeHuman:    humanFileSize(info.Size()),
			ModTime:      info.ModTime(),
			Order:        order,
		})
	}

	// Sort by order (if set) or by modification time
	sort.Slice(files, func(i, j int) bool {
		if files[i].Order >= 0 && files[j].Order >= 0 {
			return files[i].Order < files[j].Order
		}
		if files[i].Order >= 0 {
			return true
		}
		if files[j].Order >= 0 {
			return false
		}
		return files[i].ModTime.After(files[j].ModTime)
	})

	return files, nil
}

// DeleteScannedFile deletes a scanned file
func (f *FileManagerService) DeleteScannedFile(fileName string) error {
	return f.deleteFile(f.cfg.ScannedPath, fileName)
}

// DeleteArchive deletes an archive file
func (f *FileManagerService) DeleteArchive(fileName string) error {
	return f.deleteFile(f.cfg.ArchivesPath, fileName)
}

// DeletePDF deletes a PDF file
func (f *FileManagerService) DeletePDF(fileName string) error {
	return f.deleteFile(f.cfg.PDFPath, fileName)
}

// deleteFile deletes a file from the specified directory
func (f *FileManagerService) deleteFile(dir, fileName string) error {
	// Sanitize filename to prevent directory traversal
	fileName = filepath.Base(fileName)
	fullPath := filepath.Join(dir, fileName)

	// Verify the file is within the expected directory
	if filepath.Dir(fullPath) != dir {
		return fmt.Errorf("invalid file path")
	}

	if err := os.Remove(fullPath); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	// Remove from order database
	f.db.DeleteFileOrder(fileName)

	return nil
}

// DeleteAllScannedFiles deletes all scanned files
func (f *FileManagerService) DeleteAllScannedFiles() error {
	files, err := f.GetScannedFiles()
	if err != nil {
		return err
	}

	for _, file := range files {
		if err := os.Remove(file.Path); err != nil {
			return fmt.Errorf("failed to delete %s: %w", file.Name, err)
		}
		f.db.DeleteFileOrder(file.Name)
	}

	return nil
}

// DeleteAllArchives deletes all archive files
func (f *FileManagerService) DeleteAllArchives() error {
	files, err := f.GetArchives()
	if err != nil {
		return err
	}

	for _, file := range files {
		if err := os.Remove(file.Path); err != nil {
			return fmt.Errorf("failed to delete %s: %w", file.Name, err)
		}
	}

	return nil
}

// DeleteAllPDFs deletes all PDF files
func (f *FileManagerService) DeleteAllPDFs() error {
	files, err := f.GetPDFs()
	if err != nil {
		return err
	}

	for _, file := range files {
		if err := os.Remove(file.Path); err != nil {
			return fmt.Errorf("failed to delete %s: %w", file.Name, err)
		}
	}

	return nil
}

// RenameScannedFile renames a scanned file
func (f *FileManagerService) RenameScannedFile(oldName, newName string) error {
	return f.renameFile(f.cfg.ScannedPath, oldName, newName)
}

// RenameArchive renames an archive file
func (f *FileManagerService) RenameArchive(oldName, newName string) error {
	return f.renameFile(f.cfg.ArchivesPath, oldName, newName)
}

// RenamePDF renames a PDF file
func (f *FileManagerService) RenamePDF(oldName, newName string) error {
	return f.renameFile(f.cfg.PDFPath, oldName, newName)
}

// renameFile renames a file in the specified directory
func (f *FileManagerService) renameFile(dir, oldName, newName string) error {
	oldName = filepath.Base(oldName)
	newName = filepath.Base(newName)

	oldPath := filepath.Join(dir, oldName)
	newPath := filepath.Join(dir, newName)

	if err := os.Rename(oldPath, newPath); err != nil {
		return fmt.Errorf("failed to rename file: %w", err)
	}

	// Update order database
	if order, err := f.db.GetFileOrder(oldName); err == nil && order >= 0 {
		f.db.DeleteFileOrder(oldName)
		f.db.SetFileOrder(newName, order)
	}

	return nil
}

// UpdateFileOrder updates the order of files
func (f *FileManagerService) UpdateFileOrder(orders map[string]int) error {
	return f.db.UpdateFileOrders(orders)
}

// GetScannedFilePath returns the full path for a scanned file
func (f *FileManagerService) GetScannedFilePath(fileName string) string {
	return filepath.Join(f.cfg.ScannedPath, filepath.Base(fileName))
}

// GetArchivePath returns the full path for an archive file
func (f *FileManagerService) GetArchivePath(fileName string) string {
	return filepath.Join(f.cfg.ArchivesPath, filepath.Base(fileName))
}

// GetPDFPath returns the full path for a PDF file
func (f *FileManagerService) GetPDFPath(fileName string) string {
	return filepath.Join(f.cfg.PDFPath, filepath.Base(fileName))
}

// CopyFile copies a file from src to dst
func (f *FileManagerService) CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

// GenerateArchiveName generates a unique archive name
func (f *FileManagerService) GenerateArchiveName() string {
	return fmt.Sprintf("%d.zip", time.Now().Unix())
}

// GeneratePDFName generates a unique PDF name
func (f *FileManagerService) GeneratePDFName() string {
	return fmt.Sprintf("%d.pdf", time.Now().Unix())
}

// humanFileSize converts bytes to human readable format
func humanFileSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
