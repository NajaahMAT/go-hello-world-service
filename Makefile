# Makefile for Go Hello World Service

.PHONY: help build run run-port test test-coverage clean lint format install-tools

# Default target
help:
	@echo "Available targets:"
	@echo "  make build          - Build the application"
	@echo "  make run            - Run the application"
	@echo "  make run-port PORT=9999 - Run on a custom port"
	@echo "  make test           - Run all tests"
	@echo "  make test-coverage  - Run tests with coverage report"
	@echo "  make clean          - Clean build artifacts"
	@echo "  make format         - Format code with gofmt"
	@echo "  make lint           - Run golint (requires golint)"
	@echo "  make install-tools  - Install development tools"

# Build the application
build:
	@echo "Building application..."
	@mkdir -p bin
	@go build -o bin/server ./cmd/server
	@echo "Build complete: bin/server"

# Run the application
run:
	@echo "Starting server..."
	@go run ./cmd/server

# Run all tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out
	@echo "\nTo view HTML coverage report, run: go tool cover -html=coverage.out"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@rm -f coverage.out
	@echo "Clean complete"

# Format code
format:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Format complete"

# Lint code (requires golangci-lint)
lint:
	@echo "Linting code..."
	@golangci-lint run ./... || echo "Install golangci-lint: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"

# Install development tools
install-tools:
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Tools installed"

# Run the application with custom port
run-port:
	@echo "Starting server on custom port (set PORT env var)..."
	@PORT=$(PORT) go run ./cmd/server