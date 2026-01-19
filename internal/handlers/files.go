package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hightemp/wapp_scanimage_server/internal/models"
	"github.com/hightemp/wapp_scanimage_server/internal/services"
)

type FilesHandler struct {
	fileManager *services.FileManagerService
	converter   *services.ConverterService
}

func NewFilesHandler(fileManager *services.FileManagerService, converter *services.ConverterService) *FilesHandler {
	return &FilesHandler{
		fileManager: fileManager,
		converter:   converter,
	}
}

// GetScannedFiles returns the list of scanned files
// GET /api/scans
func (h *FilesHandler) GetScannedFiles(c *gin.Context) {
	files, err := h.fileManager.GetScannedFiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Data:    files,
	})
}

// DeleteScannedFile deletes a scanned file
// DELETE /api/scans/:name
func (h *FilesHandler) DeleteScannedFile(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "file name is required",
		})
		return
	}

	if err := h.fileManager.DeleteScannedFile(name); err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "File deleted successfully",
	})
}

// DeleteAllScannedFiles deletes all scanned files
// DELETE /api/scans
func (h *FilesHandler) DeleteAllScannedFiles(c *gin.Context) {
	if err := h.fileManager.DeleteAllScannedFiles(); err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "All files deleted successfully",
	})
}

// RenameScannedFile renames a scanned file
// PUT /api/scans/:name
func (h *FilesHandler) RenameScannedFile(c *gin.Context) {
	oldName := c.Param("name")
	if oldName == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "file name is required",
		})
		return
	}

	var req models.RenameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "invalid request body",
		})
		return
	}

	if err := h.fileManager.RenameScannedFile(oldName, req.NewName); err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "File renamed successfully",
	})
}

// BatchOperation performs batch operations on scanned files
// POST /api/scans/batch
func (h *FilesHandler) BatchOperation(c *gin.Context) {
	var req models.BatchOperation
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "invalid request body",
		})
		return
	}

	switch req.Action {
	case "delete":
		for _, file := range req.Files {
			if err := h.fileManager.DeleteScannedFile(file); err != nil {
				c.JSON(http.StatusInternalServerError, models.APIResponse{
					Success: false,
					Error:   err.Error(),
				})
				return
			}
		}
		c.JSON(http.StatusOK, models.APIResponse{
			Success: true,
			Message: "Files deleted successfully",
		})

	case "archive":
		archiveName, err := h.converter.CreateArchive(req.Files, "")
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse{
				Success: false,
				Error:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, models.APIResponse{
			Success: true,
			Message: "Archive created successfully",
			Data:    map[string]string{"fileName": archiveName},
		})

	case "pdf":
		pdfName, err := h.converter.ConvertToPDF(req.Files, "")
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse{
				Success: false,
				Error:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, models.APIResponse{
			Success: true,
			Message: "PDF created successfully",
			Data:    map[string]string{"fileName": pdfName},
		})

	default:
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "unknown action: " + req.Action,
		})
	}
}

// UpdateFileOrder updates the order of files
// PUT /api/scans/order
func (h *FilesHandler) UpdateFileOrder(c *gin.Context) {
	var orders map[string]int
	if err := c.ShouldBindJSON(&orders); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "invalid request body",
		})
		return
	}

	if err := h.fileManager.UpdateFileOrder(orders); err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "File order updated successfully",
	})
}

// GetArchives returns the list of archive files
// GET /api/archives
func (h *FilesHandler) GetArchives(c *gin.Context) {
	files, err := h.fileManager.GetArchives()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Data:    files,
	})
}

// DeleteArchive deletes an archive file
// DELETE /api/archives/:name
func (h *FilesHandler) DeleteArchive(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "file name is required",
		})
		return
	}

	if err := h.fileManager.DeleteArchive(name); err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Archive deleted successfully",
	})
}

// DeleteAllArchives deletes all archive files
// DELETE /api/archives
func (h *FilesHandler) DeleteAllArchives(c *gin.Context) {
	if err := h.fileManager.DeleteAllArchives(); err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "All archives deleted successfully",
	})
}

// CreateArchive creates a new archive from selected files
// POST /api/archives
func (h *FilesHandler) CreateArchive(c *gin.Context) {
	var req struct {
		Files      []string `json:"files"`
		OutputName string   `json:"outputName"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "invalid request body",
		})
		return
	}

	archiveName, err := h.converter.CreateArchive(req.Files, req.OutputName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Archive created successfully",
		Data:    map[string]string{"fileName": archiveName},
	})
}

// CreateArchiveAll creates an archive from all scanned files
// POST /api/archives/all
func (h *FilesHandler) CreateArchiveAll(c *gin.Context) {
	archiveName, err := h.converter.CreateArchiveFromAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Archive created successfully",
		Data:    map[string]string{"fileName": archiveName},
	})
}

// GetPDFs returns the list of PDF files
// GET /api/pdfs
func (h *FilesHandler) GetPDFs(c *gin.Context) {
	files, err := h.fileManager.GetPDFs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Data:    files,
	})
}

// DeletePDF deletes a PDF file
// DELETE /api/pdfs/:name
func (h *FilesHandler) DeletePDF(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "file name is required",
		})
		return
	}

	if err := h.fileManager.DeletePDF(name); err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "PDF deleted successfully",
	})
}

// DeleteAllPDFs deletes all PDF files
// DELETE /api/pdfs
func (h *FilesHandler) DeleteAllPDFs(c *gin.Context) {
	if err := h.fileManager.DeleteAllPDFs(); err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "All PDFs deleted successfully",
	})
}

// CreatePDF converts selected files to PDF
// POST /api/pdfs
func (h *FilesHandler) CreatePDF(c *gin.Context) {
	var req models.ConvertRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "invalid request body",
		})
		return
	}

	pdfName, err := h.converter.ConvertToPDF(req.Files, req.OutputName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "PDF created successfully",
		Data:    map[string]string{"fileName": pdfName},
	})
}

// CreatePDFAll converts all scanned files to a single PDF
// POST /api/pdfs/all
func (h *FilesHandler) CreatePDFAll(c *gin.Context) {
	pdfName, err := h.converter.ConvertAllToPDF()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "PDF created successfully",
		Data:    map[string]string{"fileName": pdfName},
	})
}
