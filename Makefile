.PHONY: all build dev frontend backend clean deps install uninstall run help

# Variables
SERVICE_NAME = scanimage-server
INSTALL_DIR = /usr/local/bin
SYSTEMD_DIR = /etc/systemd/system

# Default target
all: build

# Install dependencies
deps:
	@echo "Installing Go dependencies..."
	go mod download
	@echo "Installing frontend dependencies..."
	cd frontend && npm install

# Development mode - run both frontend and backend
dev: dev-backend

# Run backend in dev mode
dev-backend:
	@echo "Starting backend server..."
	go run ./cmd/server/main.go

# Run frontend dev server
dev-frontend:
	@echo "Starting frontend dev server..."
	cd frontend && npm run dev

# Build frontend
frontend:
	@echo "Building frontend..."
	cd frontend && npm run build

# Build Go binary
backend: frontend
	@echo "Building Go binary..."
	go build -o bin/scanimage-server ./cmd/server

# Full build
build: frontend backend
	@echo "Build complete!"
	@echo "Binary is at: bin/scanimage-server"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -rf cmd/server/dist/
	rm -rf frontend/node_modules/
	rm -rf frontend/dist/

# Run the built binary
run: build
	./bin/scanimage-server

# Production build with optimizations
release: frontend
	@echo "Building release binary..."
	CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/scanimage-server ./cmd/server
	@echo "Release build complete!"

# Docker build
docker:
	docker build -t scanimage-server .

# Install to system with systemd
install: release
	@echo "Installing $(SERVICE_NAME) to system..."
	@sudo install -m 755 bin/scanimage-server $(INSTALL_DIR)/$(SERVICE_NAME)
	@echo "Creating systemd service file..."
	@sudo tee $(SYSTEMD_DIR)/$(SERVICE_NAME).service > /dev/null <<EOF
[Unit]
Description=Scanimage Web Server
After=network.target

[Service]
Type=simple
ExecStart=$(INSTALL_DIR)/$(SERVICE_NAME)
Restart=on-failure
RestartSec=5
User=root
WorkingDirectory=/var/lib/$(SERVICE_NAME)

[Install]
WantedBy=multi-user.target
EOF
	@sudo mkdir -p /var/lib/$(SERVICE_NAME)
	@sudo systemctl daemon-reload
	@sudo systemctl enable $(SERVICE_NAME)
	@sudo systemctl start $(SERVICE_NAME)
	@echo "Service installed and started."

# Uninstall from system
uninstall:
	@echo "Stopping and disabling $(SERVICE_NAME) service..."
	-@sudo systemctl stop $(SERVICE_NAME) 2>/dev/null || true
	-@sudo systemctl disable $(SERVICE_NAME) 2>/dev/null || true
	@echo "Removing files..."
	@sudo rm -f $(INSTALL_DIR)/$(SERVICE_NAME)
	@sudo rm -f $(SYSTEMD_DIR)/$(SERVICE_NAME).service
	@sudo systemctl daemon-reload
	@echo "$(SERVICE_NAME) uninstalled."

# Help
help:
	@echo "Available targets:"
	@echo "  make deps         - Install all dependencies"
	@echo "  make dev          - Run backend in development mode"
	@echo "  make dev-frontend - Run frontend dev server (with hot reload)"
	@echo "  make frontend     - Build frontend only"
	@echo "  make backend      - Build backend only"
	@echo "  make build        - Build complete application"
	@echo "  make run          - Build and run the application"
	@echo "  make release      - Build optimized release binary"
	@echo "  make install      - Install to system as systemd service"
	@echo "  make uninstall    - Remove from system"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make docker       - Build Docker image"
	@echo "  make help         - Show this help"
