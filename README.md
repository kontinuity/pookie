# Pookie - Go Login Application

A simple Go web application with a login page built using the Bulma CSS framework.

## Features

- Simple login page with Bulma CSS styling
- Basic authentication (demo: admin/password)
- Dashboard page after successful login
- Comprehensive test suite
- Linting with golangci-lint
- GitHub Actions CI/CD pipeline

## Getting Started

### Prerequisites

- Go 1.19 or higher
- golangci-lint (for linting)

### Installation

1. Clone the repository:
```bash
git clone <your-repo-url>
cd pookie
```

2. Install dependencies:
```bash
go mod download
```

3. Install development tools:
```bash
make install-tools
```

### Running the Application

```bash
# Using Go directly
go run .

# Or using Make
make run
```

The application will start on `http://localhost:8080`

### Demo Credentials

- Username: `admin`
- Password: `password`

## Development

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage
```

### Linting

```bash
make lint
```

### Building

```bash
# Build the application
make build

# The binary will be created in bin/pookie
```

## Project Structure

```
.
├── .github/workflows/    # GitHub Actions workflows
├── templates/           # HTML templates
├── static/             # Static assets
├── main.go             # Main application file
├── handlers.go         # HTTP handlers
├── handlers_test.go    # Handler tests
├── .golangci.yml       # Linting configuration
├── Makefile           # Build commands
└── README.md          # This file
```

## CI/CD

The project includes GitHub Actions workflows that:

- Run tests on multiple Go versions
- Perform linting checks
- Build the application for multiple platforms
- Generate test coverage reports

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linting
5. Submit a pull request