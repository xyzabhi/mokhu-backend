# Mokhu Backend

A Go-based backend API service built with clean architecture principles.

## Project Structure

```
mokhu-backend/
├── cmd/
│   └── main.go                  # Application entry point
├── internal/
│   ├── api/                     # HTTP handlers and route definitions
│   ├── config/                  # Application configuration (env, setup)
│   ├── db/                      # Database connection and migrations
│   ├── models/                  # Data structures and validation rules
│   ├── repository/              # Data access layer
│   └── service/                 # Business logic layer
├── pkg/                         # Shared utilities (logger, auth, etc.)
├── go.mod                       # Go module dependencies
├── Makefile                     # Build and development commands
└── .env                         # Environment variables
```

## Getting Started

### Prerequisites

- Go 1.24.3 or higher
- Make (optional, for using Makefile commands)

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd mokhu-backend
```

2. Install dependencies:
```bash
make deps
```

### Development

#### Using Makefile (Recommended)

```bash
# Run the application
make run

# Build the application
make build

# Run with live reload (requires air)
make dev

# Run tests
make test

# Format code
make fmt

# Clean build artifacts
make clean
```

#### Using Go commands directly

```bash
# Run the application
go run ./cmd

# Build the application
go build -o bin/mokhu-backend ./cmd

# Run tests
go test ./...
```

## Available Commands

| Command | Description |
|---------|-------------|
| `make run` | Run the application |
| `make build` | Build binary in `bin/` directory |
| `make dev` | Run with live reload using Air |
| `make test` | Run all tests |
| `make test-coverage` | Run tests with coverage report |
| `make fmt` | Format Go code |
| `make lint` | Lint code (requires golangci-lint) |
| `make deps` | Download and tidy dependencies |
| `make clean` | Remove build artifacts |
| `make install-tools` | Install development tools |

## Environment Setup

1. Copy `.env.example` to `.env` (if available)
2. Configure your environment variables
3. Run the application with `make run`

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests: `make test`
5. Format code: `make fmt`
6. Submit a pull request
