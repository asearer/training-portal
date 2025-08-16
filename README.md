# Training Portal

A comprehensive learning management system built with Go (backend) and React (frontend).

## 🚀 Features

- **User Management**: Registration, authentication, and role-based access control
- **Course Management**: Create, update, and manage training courses
- **Module System**: Organize course content into structured modules
- **Progress Tracking**: Monitor user learning progress and completion
- **Analytics Dashboard**: Insights into learning patterns and course performance
- **Responsive Design**: Modern UI that works on all devices

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐
│   React App     │    │   Go Backend    │
│   (Frontend)    │◄──►│   (Backend)     │
└─────────────────┘    └─────────────────┘
         │                       │
         │                       │
         ▼                       ▼
┌─────────────────┐    ┌─────────────────┐
│   Vitest Tests  │    │   Go Tests      │
│   (Frontend)    │    │   (Backend)     │
└─────────────────┘    └─────────────────┘
```

## 🧪 Testing Strategy

This project implements comprehensive testing across all layers:

### Backend Testing (Go)
- **Domain Layer**: Unit tests for business models and validation
- **Use Case Layer**: Service logic tests with mocked dependencies
- **HTTP Layer**: Handler tests with mocked services
- **Middleware**: Authentication and CORS tests
- **Repository Layer**: Interface-based testing with mocks

### Frontend Testing (React)
- **Components**: User interaction and rendering tests
- **Hooks**: Custom hook logic and state management tests
- **Services**: API integration and error handling tests
- **Pages**: Full page functionality and routing tests

### Test Coverage Requirements
- **Minimum Coverage**: 80% for all layers
- **Critical Paths**: 100% coverage required
- **Business Logic**: 95% coverage required
- **Error Handling**: 90% coverage required

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- Node.js 18+
- PostgreSQL 13+

### Backend Setup

```bash
# Clone the repository
git clone <repository-url>
cd training-portal

# Install Go dependencies
go mod tidy

# Set up environment variables
cp .env.example .env
# Edit .env with your database credentials

# Run database migrations
go run migrations/migrate.go

# Start the backend server
go run cmd/server/main.go
```

### Frontend Setup

```bash
# Navigate to web directory
cd web

# Install dependencies
npm install

# Start development server
npm run dev
```

## 🧪 Running Tests

### Backend Tests

```bash
# Run all backend tests with coverage
./scripts/run_tests.sh

# Run specific test packages
go test ./internal/domain/...
go test ./internal/usecase/...
go test ./internal/interface/http/handler/...

# Run with race detection
go test -race ./internal/...

# Run benchmarks
go test -bench=. -benchmem ./internal/...
```

### Frontend Tests

```bash
# Navigate to web directory
cd web

# Run all frontend tests with coverage
./scripts/run_tests.sh

# Run specific test types
npm run test:run                    # Run tests once
npm run test:watch                  # Watch mode
npm run test:coverage               # With coverage
npm run test:ui                     # Visual test runner

# Run specific test suites
npm run test:run -- --run src/components/
npm run test:run -- --run src/pages/
npm run test:run -- --run src/hooks/
```

### Test Scripts

The project includes automated test runners:

- **`scripts/run_tests.sh`**: Backend test suite with coverage reporting
- **`web/scripts/run_tests.sh`**: Frontend test suite with coverage reporting

Both scripts provide:
- Comprehensive test execution
- Coverage reporting
- Performance metrics
- Error handling and reporting

## 📊 Test Coverage

### Backend Coverage
- **Domain Models**: 95%+
- **Business Logic**: 90%+
- **HTTP Handlers**: 85%+
- **Middleware**: 90%+
- **Overall**: 85%+

### Frontend Coverage
- **Components**: 85%+
- **Hooks**: 90%+
- **Services**: 85%+
- **Pages**: 80%+
- **Overall**: 85%+

## 🏗️ Project Structure

```
training-portal/
├── cmd/                           # Application entry points
│   └── server/
│       └── main.go               # Backend server
├── internal/                      # Backend application code
│   ├── domain/                   # Business models and logic
│   ├── usecase/                  # Application services
│   └── interface/                # External interfaces
│       ├── http/                 # HTTP handlers and middleware
│       └── repository/           # Data access layer
├── web/                          # Frontend React application
│   ├── src/
│   │   ├── components/           # Reusable UI components
│   │   ├── pages/                # Page components
│   │   ├── hooks/                # Custom React hooks
│   │   ├── services/             # API services
│   │   └── test/                 # Test configuration
│   ├── package.json              # Frontend dependencies
│   └── vitest.config.ts          # Test configuration
├── scripts/                       # Test and utility scripts
├── docs/                          # Documentation
│   └── testing-guide.md          # Comprehensive testing guide
├── go.mod                         # Go module definition
└── README.md                      # This file
```

## 🔧 Development

### Adding New Tests

#### Backend Tests
1. Create test file: `filename_test.go`
2. Follow naming convention: `TestFunctionName_Scenario_ExpectedResult`
3. Use table-driven tests for multiple scenarios
4. Mock external dependencies using interfaces

#### Frontend Tests
1. Create test file: `ComponentName.test.tsx`
2. Test user interactions, not implementation details
3. Mock API calls and external services
4. Use React Testing Library queries

### Test Best Practices

1. **Test the behavior, not the implementation**
2. **Use descriptive test names**
3. **Test one thing per test**
4. **Keep tests independent**
5. **Use table-driven tests for multiple scenarios**
6. **Mock external dependencies**
7. **Test error conditions**
8. **Maintain high test coverage**

## 📚 Documentation

- **[Testing Guide](docs/testing-guide.md)**: Comprehensive testing documentation
- **[API Documentation](docs/api.md)**: Backend API endpoints and usage
- **[Frontend Guide](docs/frontend.md)**: React component library and patterns

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Write tests for new functionality
4. Ensure all tests pass: `./scripts/run_tests.sh`
5. Commit your changes: `git commit -am 'Add feature'`
6. Push to the branch: `git push origin feature-name`
7. Submit a pull request

### Testing Requirements for Contributions

- All new code must include corresponding tests
- Test coverage must not decrease
- All tests must pass before merging
- Include integration tests for new features

## 🚀 Deployment

### Backend Deployment
```bash
# Build binary
go build -o bin/server cmd/server/main.go

# Run with environment variables
./bin/server
```

### Frontend Deployment
```bash
cd web

# Build for production
npm run build

# Serve static files
npm run preview
```

## 📊 Monitoring and Analytics

- **Test Coverage**: Automated coverage reporting
- **Performance**: Benchmark testing for critical paths
- **Quality**: Linting and type checking
- **Security**: Dependency vulnerability scanning

## 🆘 Troubleshooting

### Common Issues

1. **Test failures**: Check mock implementations and dependencies
2. **Coverage issues**: Ensure all code paths are tested
3. **Performance problems**: Run benchmarks to identify bottlenecks
4. **Import errors**: Verify Go module setup and import paths




