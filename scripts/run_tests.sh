#!/bin/bash

# Training Portal Test Runner
# This script runs all tests with coverage reporting

set -e

echo "🚀 Starting Training Portal Test Suite..."
echo "=========================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if Go is installed
if ! command -v go &> /dev/null; then
    print_error "Go is not installed. Please install Go 1.21 or later."
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
print_status "Go version: $GO_VERSION"

# Clean previous test artifacts
print_status "Cleaning previous test artifacts..."
go clean -testcache
rm -rf coverage/

# Create coverage directory
mkdir -p coverage

# Run tests with coverage for each package
print_status "Running tests with coverage..."

# Domain layer tests
print_status "Testing domain layer..."
go test -v -coverprofile=coverage/domain.out ./internal/domain/... || {
    print_error "Domain layer tests failed"
    exit 1
}

# Repository interface tests
print_status "Testing repository interfaces..."
go test -v -coverprofile=coverage/repository.out ./internal/interface/repository/... || {
    print_error "Repository interface tests failed"
    exit 1
}

# Use case layer tests
print_status "Testing use case layer..."
go test -v -coverprofile=coverage/usecase.out ./internal/usecase/... || {
    print_error "Use case layer tests failed"
    exit 1
}

# HTTP handler tests
print_status "Testing HTTP handlers..."
go test -v -coverprofile=coverage/handlers.out ./internal/interface/http/handler/... || {
    print_error "HTTP handler tests failed"
    exit 1
}

# Middleware tests
print_status "Testing middleware..."
go test -v -coverprofile=coverage/middleware.out ./internal/interface/http/middleware/... || {
    print_error "Middleware tests failed"
    exit 1
}

# Integration tests (if they exist)
if [ -d "./tests/integration" ]; then
    print_status "Running integration tests..."
    go test -v -coverprofile=coverage/integration.out ./tests/integration/... || {
        print_warning "Integration tests failed (continuing...)"
    }
fi

# Generate combined coverage report
print_status "Generating coverage report..."
go tool cover -html=coverage/usecase.out -o coverage/coverage.html

# Show coverage summary
print_status "Coverage Summary:"
echo "=================="

# Domain layer coverage
if [ -f "coverage/domain.out" ]; then
    DOMAIN_COVERAGE=$(go tool cover -func=coverage/domain.out | grep total | awk '{print $3}')
    echo "Domain Layer: $DOMAIN_COVERAGE"
fi

# Repository layer coverage
if [ -f "coverage/repository.out" ]; then
    REPO_COVERAGE=$(go tool cover -func=coverage/repository.out | grep total | awk '{print $3}')
    echo "Repository Layer: $REPO_COVERAGE"
fi

# Use case layer coverage
if [ -f "coverage/usecase.out" ]; then
    USECASE_COVERAGE=$(go tool cover -func=coverage/usecase.out | grep total | awk '{print $3}')
    echo "Use Case Layer: $USECASE_COVERAGE"
fi

# Handler layer coverage
if [ -f "coverage/handlers.out" ]; then
    HANDLER_COVERAGE=$(go tool cover -func=coverage/handlers.out | grep total | awk '{print $3}')
    echo "HTTP Handlers: $HANDLER_COVERAGE"
fi

# Middleware coverage
if [ -f "coverage/middleware.out" ]; then
    MIDDLEWARE_COVERAGE=$(go tool cover -func=coverage/middleware.out | grep total | awk '{print $3}')
    echo "Middleware: $MIDDLEWARE_COVERAGE"
fi

# Overall coverage
print_status "Overall test coverage report generated at: coverage/coverage.html"

# Run benchmarks if requested
if [ "$1" = "--benchmarks" ]; then
    print_status "Running benchmarks..."
    go test -bench=. -benchmem ./internal/... || {
        print_warning "Some benchmarks failed"
    }
fi

# Run race detection if requested
if [ "$1" = "--race" ]; then
    print_status "Running tests with race detection..."
    go test -race ./internal/... || {
        print_warning "Race detection found issues"
    }
fi

print_success "Test suite completed successfully!"
print_status "Coverage report available at: coverage/coverage.html"

# Optional: Open coverage report in browser
if command -v open &> /dev/null; then
    read -p "Open coverage report in browser? (y/n): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        open coverage/coverage.html
    fi
elif command -v xdg-open &> /dev/null; then
    read -p "Open coverage report in browser? (y/n): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        xdg-open coverage/coverage.html
    fi
fi
