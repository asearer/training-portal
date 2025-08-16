# Training Portal Testing Guide

This document provides comprehensive guidance for testing the Training Portal application, covering both backend (Go) and frontend (React) testing strategies.

## Table of Contents

1. [Overview](#overview)
2. [Backend Testing](#backend-testing)
3. [Frontend Testing](#frontend-testing)
4. [Test Coverage](#test-coverage)
5. [Running Tests](#running-tests)
6. [Writing Tests](#writing-tests)
7. [Best Practices](#best-practices)
8. [Troubleshooting](#troubleshooting)

## Overview

The Training Portal uses a comprehensive testing strategy with:
- **Backend**: Go testing with `testing` package and mocks
- **Frontend**: React testing with Vitest, React Testing Library, and Jest DOM
- **Coverage**: Minimum 80% coverage requirement for all layers
- **Automation**: CI/CD integration with automated test runs

## Backend Testing

### Architecture

The backend follows a layered architecture with corresponding test layers:

```
┌─────────────────┐
│   HTTP Layer    │ ← Handler tests
├─────────────────┤
│  Business Logic │ ← Service tests
├─────────────────┤
│   Data Access   │ ← Repository tests
└─────────────────┘
```

### Test Structure

```
internal/
├── domain/
│   └── user/
│       ├── model.go
│       └── model_test.go          # Domain model tests
├── usecase/
│   └── user/
│       ├── service.go
│       └── service_test.go        # Business logic tests
├── interface/
│   ├── http/
│   │   ├── handler/
│   │   │   ├── user.go
│   │   │   └── user_test.go       # HTTP handler tests
│   │   └── middleware/
│   │       ├── auth.go
│   │       └── auth_test.go       # Middleware tests
│   └── repository/
│       ├── user_repository.go     # Interface definitions
│       └── postgres/
│           └── user.go            # Concrete implementations
```

### Running Backend Tests

```bash
# Run all tests
./scripts/run_tests.sh

# Run specific test packages
go test ./internal/domain/...
go test ./internal/usecase/...
go test ./internal/interface/http/handler/...

# Run with coverage
go test -coverprofile=coverage.out ./internal/...
go tool cover -html=coverage.out

# Run with race detection
go test -race ./internal/...

# Run benchmarks
go test -bench=. -benchmem ./internal/...
```

### Test Examples

#### Domain Model Test

```go
func TestUser_Validation(t *testing.T) {
    tests := []struct {
        name    string
        user    *User
        wantErr bool
    }{
        {
            name: "Valid user",
            user: &User{
                ID:    "123",
                Name:  "John Doe",
                Email: "john@example.com",
                Role:  RoleEmployee,
            },
            wantErr: false,
        },
        // ... more test cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := validateUser(tt.user)
            if (err != nil) != tt.wantErr {
                t.Errorf("User validation error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

#### Service Test with Mock

```go
func TestUserService_Register(t *testing.T) {
    mockRepo := NewMockUserRepository()
    service := &UserService{Repo: mockRepo}
    
    user, err := service.Register("John Doe", "john@example.com", "password123", RoleEmployee)
    
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    
    if user == nil {
        t.Error("Expected user, got nil")
    }
}
```

#### HTTP Handler Test

```go
func TestUserHandler_Register(t *testing.T) {
    mockService := new(MockUserService)
    handler := &UserHandler{Service: mockService}
    
    app := createTestApp()
    app.Post("/register", handler.Register)
    
    req := httptest.NewRequest("POST", "/register", bytes.NewReader(requestBody))
    resp, err := app.Test(req)
    
    assert.NoError(t, err)
    assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}
```

## Frontend Testing

### Architecture

The frontend uses a modern testing stack:

```
┌─────────────────┐
│   Components    │ ← Component tests
├─────────────────┤
│     Hooks       │ ← Hook tests
├─────────────────┤
│    Services     │ ← Service tests
├─────────────────┤
│     Utils       │ ← Utility tests
└─────────────────┘
```

### Test Structure

```
web/src/
├── components/
│   ├── layout/
│   │   ├── Header.tsx
│   │   └── Header.test.tsx        # Component tests
│   └── ui/
│       ├── Button.tsx
│       └── Button.test.tsx
├── pages/
│   ├── auth/
│   │   ├── Login.tsx
│   │   └── Login.test.tsx         # Page tests
│   └── dashboard/
│       ├── Dashboard.tsx
│       └── Dashboard.test.tsx
├── hooks/
│   ├── useAuth.ts
│   └── useAuth.test.ts            # Hook tests
├── services/
│   ├── authService.ts
│   └── authService.test.ts        # Service tests
└── test/
    └── setup.ts                   # Test configuration
```

### Running Frontend Tests

```bash
# Navigate to web directory
cd web

# Run all tests
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

### Test Examples

#### Component Test

```tsx
import { render, screen, fireEvent } from '@testing-library/react'
import { describe, it, expect, vi } from 'vitest'
import Login from './Login'

describe('Login Component', () => {
    it('renders login form correctly', () => {
        render(<Login />)
        
        expect(screen.getByText(/sign in/i)).toBeInTheDocument()
        expect(screen.getByLabelText(/email/i)).toBeInTheDocument()
        expect(screen.getByLabelText(/password/i)).toBeInTheDocument()
    })
    
    it('handles form submission', async () => {
        const mockLogin = vi.fn()
        render(<Login onLogin={mockLogin} />)
        
        fireEvent.change(screen.getByLabelText(/email/i), {
            target: { value: 'test@example.com' }
        })
        fireEvent.click(screen.getByRole('button', { name: /sign in/i }))
        
        expect(mockLogin).toHaveBeenCalledWith('test@example.com')
    })
})
```

#### Hook Test

```tsx
import { renderHook, act } from '@testing-library/react'
import { describe, it, expect } from 'vitest'
import { useCounter } from './useCounter'

describe('useCounter Hook', () => {
    it('initializes with default value', () => {
        const { result } = renderHook(() => useCounter())
        
        expect(result.current.count).toBe(0)
    })
    
    it('increments counter', () => {
        const { result } = renderHook(() => useCounter())
        
        act(() => {
            result.current.increment()
        })
        
        expect(result.current.count).toBe(1)
    })
})
```

#### Service Test

```tsx
import { describe, it, expect, vi, beforeEach } from 'vitest'
import { authService } from './authService'

// Mock axios
const mockAxios = {
    post: vi.fn(),
}

vi.mock('axios', () => ({
    default: mockAxios,
}))

describe('Auth Service', () => {
    beforeEach(() => {
        vi.clearAllMocks()
    })
    
    it('handles successful login', async () => {
        mockAxios.post.mockResolvedValue({
            data: { token: 'fake-token' }
        })
        
        const result = await authService.login('test@example.com', 'password')
        
        expect(result.token).toBe('fake-token')
        expect(mockAxios.post).toHaveBeenCalledWith('/login', {
            email: 'test@example.com',
            password: 'password'
        })
    })
})
```

## Test Coverage

### Coverage Requirements

- **Minimum Coverage**: 80% for all layers
- **Critical Paths**: 100% coverage required
- **Business Logic**: 95% coverage required
- **Error Handling**: 90% coverage required

### Coverage Reports

#### Backend Coverage

```bash
# Generate coverage report
go test -coverprofile=coverage.out ./internal/...

# View in browser
go tool cover -html=coverage.out -o coverage.html
open coverage.html
```

#### Frontend Coverage

```bash
# Generate coverage report
npm run test:coverage

# View in browser
open coverage/index.html
```

### Coverage Types

1. **Statement Coverage**: Each statement executed at least once
2. **Branch Coverage**: Each branch of conditional statements executed
3. **Function Coverage**: Each function called at least once
4. **Line Coverage**: Each line of source code executed

## Running Tests

### Backend Test Runner

```bash
# Full test suite with coverage
./scripts/run_tests.sh

# Specific options
./scripts/run_tests.sh --benchmarks    # Include benchmarks
./scripts/run_tests.sh --race          # Include race detection
```

### Frontend Test Runner

```bash
# Full test suite with coverage
cd web && ./scripts/run_tests.sh

# Specific options
./scripts/run_tests.sh --unit          # Unit tests only
./scripts/run_tests.sh --integration   # Integration tests only
./scripts/run_tests.sh --watch         # Watch mode
./scripts/run_tests.sh --ui            # Visual test runner
```

### CI/CD Integration

```yaml
# .github/workflows/test.yml
name: Tests
on: [push, pull_request]

jobs:
  backend-tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v3
        with:
          go-version: '1.21'
      - run: ./scripts/run_tests.sh
  
  frontend-tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-node@v3
        with:
          node-version: '18'
      - run: cd web && ./scripts/run_tests.sh
```

## Writing Tests

### Test Structure

Follow the AAA pattern (Arrange, Act, Assert):

```go
func TestExample(t *testing.T) {
    // Arrange - Set up test data and mocks
    mockRepo := NewMockRepository()
    service := NewService(mockRepo)
    
    // Act - Execute the function being tested
    result, err := service.DoSomething("test")
    
    // Assert - Verify the results
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if result != "expected" {
        t.Errorf("Expected 'expected', got %s", result)
    }
}
```

### Test Naming

Use descriptive test names that explain the scenario:

```go
// Good
func TestUserService_Register_WithValidData_ReturnsUser(t *testing.T)
func TestUserService_Register_WithInvalidEmail_ReturnsError(t *testing.T)

// Avoid
func TestRegister(t *testing.T)
func TestError(t *testing.T)
```

### Mocking Strategy

1. **Use interfaces** for dependency injection
2. **Create mock implementations** for testing
3. **Verify mock calls** to ensure correct behavior
4. **Use table-driven tests** for multiple scenarios

```go
type MockUserRepository struct {
    users map[string]*User
}

func (m *MockUserRepository) Create(u *User) error {
    m.users[u.ID] = u
    return nil
}

func (m *MockUserRepository) FindByID(id string) (*User, error) {
    if user, exists := m.users[id]; exists {
        return user, nil
    }
    return nil, nil
}
```

### Error Testing

Always test error conditions:

```go
func TestUserService_GetUser_WithInvalidID_ReturnsError(t *testing.T) {
    mockRepo := NewMockUserRepository()
    service := &UserService{Repo: mockRepo}
    
    user, err := service.GetUser("")
    
    if err == nil {
        t.Error("Expected error for empty ID")
    }
    if user != nil {
        t.Error("Expected nil user for error case")
    }
}
```

## Best Practices

### General Testing

1. **Test the behavior, not the implementation**
2. **Use descriptive test names**
3. **Test one thing per test**
4. **Keep tests independent**
5. **Use table-driven tests for multiple scenarios**
6. **Mock external dependencies**
7. **Test error conditions**
8. **Maintain high test coverage**

### Backend Testing

1. **Test each layer independently**
2. **Use interfaces for dependency injection**
3. **Mock database calls**
4. **Test HTTP handlers with httptest**
5. **Use testify for assertions**
6. **Test middleware separately**

### Frontend Testing

1. **Test user interactions, not implementation details**
2. **Use React Testing Library queries**
3. **Mock API calls and external services**
4. **Test accessibility features**
5. **Use data-testid sparingly**
6. **Test error boundaries**
7. **Test responsive behavior**

### Performance Testing

1. **Use benchmarks for performance-critical code**
2. **Test with realistic data sizes**
3. **Measure memory usage**
4. **Test concurrent scenarios**

```go
func BenchmarkUserService_GetUser(b *testing.B) {
    mockRepo := NewMockUserRepository()
    service := &UserService{Repo: mockRepo}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        service.GetUser("test-id")
    }
}
```

## Troubleshooting

### Common Issues

#### Backend Tests

1. **Import cycle errors**
   - Use interfaces to break dependencies
   - Move mocks to test files

2. **Database connection issues**
   - Use mocks instead of real database
   - Use test containers for integration tests

3. **Race conditions**
   - Run tests with `-race` flag
   - Use proper synchronization in tests

#### Frontend Tests

1. **Module resolution errors**
   - Check tsconfig.json paths
   - Verify import aliases

2. **Mock not working**
   - Ensure mocks are defined before imports
   - Check mock implementation

3. **Test environment issues**
   - Verify jsdom setup
   - Check global mocks in setup.ts

### Debugging Tests

```bash
# Backend - verbose output
go test -v ./internal/...

# Frontend - debug mode
npm run test:run -- --reporter=verbose

# Run single test
go test -run TestSpecificFunction ./internal/...
npm run test:run -- --run src/path/to/test.tsx
```

### Performance Issues

1. **Slow test execution**
   - Use parallel tests where possible
   - Reduce setup/teardown overhead
   - Use test suites for related tests

2. **Memory leaks**
   - Clean up resources in tests
   - Use `t.Cleanup()` for cleanup
   - Reset mocks between tests

## Conclusion

This testing guide provides a comprehensive approach to testing the Training Portal application. By following these practices, you can ensure:

- High code quality and reliability
- Easy maintenance and refactoring
- Confidence in deployments
- Better developer experience

Remember to:
- Write tests as you develop features
- Maintain high test coverage
- Review and update tests regularly
- Use tests as documentation

For additional help, refer to:
- [Go Testing Documentation](https://golang.org/pkg/testing/)
- [React Testing Library](https://testing-library.com/docs/react-testing-library/intro/)
- [Vitest Documentation](https://vitest.dev/)
- [Jest DOM Matchers](https://github.com/testing-library/jest-dom)
