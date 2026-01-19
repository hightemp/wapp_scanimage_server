.PHONY: all build dev frontend backend clean install run help

# Default target
all: build

# Install dependencies
install:
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

# Help
help:
	@echo "Available targets:"
	@echo "  make install      - Install all dependencies"
	@echo "  make dev          - Run backend in development mode"
	@echo "  make dev-frontend - Run frontend dev server (with hot reload)"
	@echo "  make frontend     - Build frontend only"
	@echo "  make backend      - Build backend only"
	@echo "  make build        - Build complete application"
	@echo "  make run          - Build and run the application"
	@echo "  make release      - Build optimized release binary"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make docker       - Build Docker image"
	@echo "  make help         - Show this help"
