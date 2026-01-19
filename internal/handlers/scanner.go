package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hightemp/wapp_scanimage_server/internal/models"
	"github.com/hightemp/wapp_scanimage_server/internal/services"
)

type ScannerHandler struct {
	scannerService *services.ScannerService
}

func NewScannerHandler(scannerService *services.ScannerService) *ScannerHandler {
	return &ScannerHandler{scannerService: scannerService}
}

// GetScanners returns the list of available scanners
// GET /api/scanners
func (h *ScannerHandler) GetScanners(c *gin.Context) {
	scanners, err := h.scannerService.GetScanners()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Data:    scanners,
	})
}

// GetScannersRaw returns raw scanner list
// GET /api/scanners/raw
func (h *ScannerHandler) GetScannersRaw(c *gin.Context) {
	output, err := h.scannerService.GetScannersRaw()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Data:    output,
	})
}

// Scan performs a scan operation
// POST /api/scan
func (h *ScannerHandler) Scan(c *gin.Context) {
	var settings models.ScanSettings
	if err := c.ShouldBindJSON(&settings); err != nil {
		// Use default settings if not provided
		settings = models.ScanSettings{
			Resolution: 300,
			Quality:    80,
			Format:     "jpeg",
		}
	}

	// Set defaults for missing values
	if settings.Resolution == 0 {
		settings.Resolution = 300
	}
	if settings.Quality == 0 {
		settings.Quality = 80
	}
	if settings.Format == "" {
		settings.Format = "jpeg"
	}

	fileName, err := h.scannerService.Scan(settings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Scan completed successfully",
		Data: map[string]string{
			"fileName": fileName,
		},
	})
}
