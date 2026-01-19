package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/hightemp/wapp_scanimage_server/internal/config"
	"github.com/hightemp/wapp_scanimage_server/internal/database"
	"github.com/hightemp/wapp_scanimage_server/internal/handlers"
	"github.com/hightemp/wapp_scanimage_server/internal/services"
	"github.com/joho/godotenv"
)

//go:embed dist/*
var frontendFS embed.FS

func main() {
	// Load .env file if exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize configuration
	cfg := config.New()

	// Initialize database
	db, err := database.New(cfg.DatabasePath())
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize services
	scannerService := services.NewScannerService(cfg)
	fileManagerService := services.NewFileManagerService(cfg, db)
	converterService := services.NewConverterService(cfg)

	// Initialize handlers
	scannerHandler := handlers.NewScannerHandler(scannerService)
	filesHandler := handlers.NewFilesHandler(fileManagerService, converterService)
	settingsHandler := handlers.NewSettingsHandler(db)

	// Setup Gin
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// CORS middleware
	router.Use(corsMiddleware())

	// API routes
	api := router.Group("/api")
	{
		// Scanner routes
		api.GET("/scanners", scannerHandler.GetScanners)
		api.GET("/scanners/raw", scannerHandler.GetScannersRaw)
		api.POST("/scan", scannerHandler.Scan)

		// Scanned files routes
		api.GET("/scans", filesHandler.GetScannedFiles)
		api.DELETE("/scans", filesHandler.DeleteAllScannedFiles)
		api.DELETE("/scans/:name", filesHandler.DeleteScannedFile)
		api.PUT("/scans/:name", filesHandler.RenameScannedFile)
		api.POST("/scans/batch", filesHandler.BatchOperation)
		api.PUT("/scans/order", filesHandler.UpdateFileOrder)

		// Archive routes
		api.GET("/archives", filesHandler.GetArchives)
		api.POST("/archives", filesHandler.CreateArchive)
		api.POST("/archives/all", filesHandler.CreateArchiveAll)
		api.DELETE("/archives", filesHandler.DeleteAllArchives)
		api.DELETE("/archives/:name", filesHandler.DeleteArchive)

		// PDF routes
		api.GET("/pdfs", filesHandler.GetPDFs)
		api.POST("/pdfs", filesHandler.CreatePDF)
		api.POST("/pdfs/all", filesHandler.CreatePDFAll)
		api.DELETE("/pdfs", filesHandler.DeleteAllPDFs)
		api.DELETE("/pdfs/:name", filesHandler.DeletePDF)

		// Settings routes
		api.GET("/settings", settingsHandler.GetSettings)
		api.PUT("/settings", settingsHandler.UpdateSettings)
		api.GET("/settings/:key", settingsHandler.GetSetting)
		api.PUT("/settings/:key", settingsHandler.SetSetting)
	}

	// Serve static files (scanned images, archives, pdfs)
	router.Static("/files/scanned", cfg.ScannedPath)
	router.Static("/files/archives", cfg.ArchivesPath)
	router.Static("/files/pdf", cfg.PDFPath)

	// Serve frontend
	setupFrontend(router)

	// Start server
	log.Printf("Starting server on %s", cfg.Address())
	log.Printf("Debug mode: %v", cfg.Debug)

	// Graceful shutdown
	go func() {
		if err := router.Run(cfg.Address()); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func setupFrontend(router *gin.Engine) {
	// Try to use embedded frontend first
	distFS, err := fs.Sub(frontendFS, "dist")
	if err != nil {
		log.Println("No embedded frontend found, will serve from ./frontend/dist if available")
		// Fallback to filesystem
		if _, err := os.Stat("./frontend/dist"); err == nil {
			router.Static("/assets", "./frontend/dist/assets")
			router.StaticFile("/", "./frontend/dist/index.html")
			router.NoRoute(func(c *gin.Context) {
				c.File("./frontend/dist/index.html")
			})
		}
		return
	}

	// Serve embedded frontend
	router.StaticFS("/assets", http.FS(distFS))

	router.GET("/", func(c *gin.Context) {
		data, err := fs.ReadFile(distFS, "index.html")
		if err != nil {
			c.String(http.StatusNotFound, "Frontend not found")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})

	router.NoRoute(func(c *gin.Context) {
		data, err := fs.ReadFile(distFS, "index.html")
		if err != nil {
			c.String(http.StatusNotFound, "Page not found")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})
}
