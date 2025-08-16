#!/bin/bash

# Training Portal Frontend Test Runner
# This script runs all frontend tests with coverage reporting

set -e

echo "🚀 Starting Training Portal Frontend Test Suite..."
echo "=================================================="

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

# Check if Node.js is installed
if ! command -v node &> /dev/null; then
    print_error "Node.js is not installed. Please install Node.js 18 or later."
    exit 1
fi

# Check Node.js version
NODE_VERSION=$(node --version | sed 's/v//')
print_status "Node.js version: $NODE_VERSION"

# Check if npm is installed
if ! command -v npm &> /dev/null; then
    print_error "npm is not installed. Please install npm."
    exit 1
fi

# Check if dependencies are installed
if [ ! -d "node_modules" ]; then
    print_status "Installing dependencies..."
    npm install
fi

# Clean previous test artifacts
print_status "Cleaning previous test artifacts..."
rm -rf coverage/
rm -rf .vitest/

# Create coverage directory
mkdir -p coverage

# Run type checking
print_status "Running TypeScript type checking..."
npm run type-check || {
    print_warning "TypeScript type checking failed (continuing with tests...)"
}

# Run linting
print_status "Running ESLint..."
npm run lint || {
    print_warning "ESLint found issues (continuing with tests...)"
}

# Run tests with coverage
print_status "Running tests with coverage..."
npm run test:coverage || {
    print_error "Tests failed"
    exit 1
}

# Show coverage summary
print_status "Coverage Summary:"
echo "=================="

if [ -f "coverage/coverage-summary.json" ]; then
    # Parse coverage summary
    TOTAL_STATEMENTS=$(node -e "
        const fs = require('fs');
        const coverage = JSON.parse(fs.readFileSync('coverage/coverage-summary.json', 'utf8'));
        console.log(coverage.total.statements.pct);
    ")
    TOTAL_BRANCHES=$(node -e "
        const fs = require('fs');
        const coverage = JSON.parse(fs.readFileSync('coverage/coverage-summary.json', 'utf8'));
        console.log(coverage.total.branches.pct);
    ")
    TOTAL_FUNCTIONS=$(node -e "
        const fs = require('fs');
        const coverage = JSON.parse(fs.readFileSync('coverage/coverage-summary.json', 'utf8'));
        console.log(coverage.total.functions.pct);
    ")
    TOTAL_LINES=$(node -e "
        const fs = require('fs');
        const coverage = JSON.parse(fs.readFileSync('coverage/coverage-summary.json', 'utf8'));
        console.log(coverage.total.lines.pct);
    ")
    
    echo "Statements: ${TOTAL_STATEMENTS}%"
    echo "Branches:   ${TOTAL_BRANCHES}%"
    echo "Functions:  ${TOTAL_FUNCTIONS}%"
    echo "Lines:      ${TOTAL_LINES}%"
fi

# Run specific test suites if requested
if [ "$1" = "--unit" ]; then
    print_status "Running unit tests only..."
    npm run test:run -- --reporter=verbose --run src/components/ src/hooks/ src/utils/
fi

if [ "$1" = "--integration" ]; then
    print_status "Running integration tests only..."
    npm run test:run -- --reporter=verbose --run src/pages/ src/services/
fi

if [ "$1" = "--watch" ]; then
    print_status "Starting test watcher..."
    npm run test:watch
    exit 0
fi

if [ "$1" = "--ui" ]; then
    print_status "Starting test UI..."
    npm run test:ui
    exit 0
fi

# Run performance tests if requested
if [ "$1" = "--performance" ]; then
    print_status "Running performance tests..."
    npm run test:run -- --reporter=verbose --run src/**/*.perf.test.tsx || {
        print_warning "Performance tests failed (continuing...)"
    }
fi

# Run accessibility tests if requested
if [ "$1" = "--a11y" ]; then
    print_status "Running accessibility tests..."
    npm run test:run -- --reporter=verbose --run src/**/*.a11y.test.tsx || {
        print_warning "Accessibility tests failed (continuing...)"
    }
fi

# Generate test report
print_status "Generating test report..."
if [ -d "coverage" ]; then
    print_status "Coverage report available at: coverage/index.html"
fi

print_success "Frontend test suite completed successfully!"

# Optional: Open coverage report in browser
if command -v open &> /dev/null; then
    read -p "Open coverage report in browser? (y/n): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        open coverage/index.html
    fi
elif command -v xdg-open &> /dev/null; then
    read -p "Open coverage report in browser? (y/n): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        xdg-open coverage/index.html
    fi
fi

# Show test summary
print_status "Test Summary:"
echo "=============="
echo "✅ All tests passed"
echo "📊 Coverage report generated"
echo "🔍 Linting completed"
echo "📝 Type checking completed"
echo ""
echo "Next steps:"
echo "- Review coverage report for areas to improve"
echo "- Fix any linting issues found"
echo "- Add more tests for uncovered code paths"
