.PHONY: build run test lint clean install-tools

# Build the application
build:
	go build -o bin/pookie .

# Run the application
run:
	go run .

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run linting
lint:
	golangci-lint run

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Install development tools
install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run the application in development mode with auto-reload
dev:
	go run .

# Format code
fmt:
	go fmt ./...