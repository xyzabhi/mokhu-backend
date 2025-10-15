# Go application Makefile

# Build the application
build:
	go build -o bin/mokhu-backend ./cmd

# Run the application
run:
	go run ./cmd

# Run the application with live reload (requires air)
dev:
	air

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod download
	go mod tidy

# Run tests
test:
	go test ./...

# Run tests with coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Start MailHog for email testing
mailhog:
	docker-compose up mailhog -d

# Start all services (database, redis, mailhog)
services:
	docker-compose up -d

# Stop all services
stop-services:
	docker-compose down

# Install development tools
install-tools:
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: build run dev clean deps test test-coverage fmt lint install-tools mailhog services stop-services
