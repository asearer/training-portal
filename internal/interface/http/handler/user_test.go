package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"training-portal/internal/domain/user"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock implementation of the user service
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Register(name, email, password string, role user.Role) (*user.User, error) {
	args := m.Called(name, email, password, role)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockUserService) Login(email, password string) (*user.User, error) {
	args := m.Called(email, password)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockUserService) GetUser(id string) (*user.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockUserService) UpdateUser(u *user.User) error {
	args := m.Called(u)
	return args.Error(0)
}

func (m *MockUserService) UpdatePassword(id, newPassword string) error {
	args := m.Called(id, newPassword)
	return args.Error(0)
}

func (m *MockUserService) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserService) ListUsers() ([]*user.User, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*user.User), args.Error(1)
}

// Helper function to create a test Fiber app
func createTestApp() *fiber.App {
	app := fiber.New()
	return app
}

func TestUserHandler_Register(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    map[string]interface{}
		mockSetup      func(*MockUserService)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "Successful registration",
			requestBody: map[string]interface{}{
				"name":     "John Doe",
				"email":    "john@example.com",
				"password": "password123",
			},
			mockSetup: func(mockService *MockUserService) {
				expectedUser := &user.User{
					ID:    "123",
					Name:  "John Doe",
					Email: "john@example.com",
					Role:  user.RoleEmployee,
				}
				mockService.On("Register", "John Doe", "john@example.com", "password123", user.RoleEmployee).
					Return(expectedUser, nil)
			},
			expectedStatus: fiber.StatusCreated,
			expectedBody: map[string]interface{}{
				"id":    "123",
				"name":  "John Doe",
				"email": "john@example.com",
				"role":  "employee",
			},
		},
		{
			name: "Invalid request body",
			requestBody: map[string]interface{}{
				"name": "John Doe",
				// Missing email and password
			},
			mockSetup:      func(mockService *MockUserService) {},
			expectedStatus: fiber.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "Invalid request",
			},
		},
		{
			name: "Service error",
			requestBody: map[string]interface{}{
				"name":     "John Doe",
				"email":    "john@example.com",
				"password": "password123",
			},
			mockSetup: func(mockService *MockUserService) {
				mockService.On("Register", "John Doe", "john@example.com", "password123", user.RoleEmployee).
					Return(nil, errors.New("email already exists"))
			},
			expectedStatus: fiber.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "email already exists",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock service
			mockService := new(MockUserService)
			tt.mockSetup(mockService)

			// Create handler
			handler := &UserHandler{Service: mockService}

			// Create test app
			app := createTestApp()
			app.Post("/register", handler.Register)

			// Create request
			requestBody, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("POST", "/register", bytes.NewReader(requestBody))
			req.Header.Set("Content-Type", "application/json")

			// Make request
			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			// Assert status code
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// Parse response body
			var responseBody map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			assert.NoError(t, err)

			// Assert response body
			for key, expectedValue := range tt.expectedBody {
				if key == "password" {
					// Password should never be returned
					assert.NotContains(t, responseBody, "password")
				} else {
					assert.Equal(t, expectedValue, responseBody[key])
				}
			}

			// Verify all mocks were called
			mockService.AssertExpectations(t)
		})
	}
}

func TestUserHandler_Login(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    map[string]interface{}
		mockSetup      func(*MockUserService)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "Successful login",
			requestBody: map[string]interface{}{
				"email":    "john@example.com",
				"password": "password123",
			},
			mockSetup: func(mockService *MockUserService) {
				expectedUser := &user.User{
					ID:    "123",
					Name:  "John Doe",
					Email: "john@example.com",
					Role:  user.RoleEmployee,
				}
				mockService.On("Login", "john@example.com", "password123").
					Return(expectedUser, nil)
			},
			expectedStatus: fiber.StatusOK,
			expectedBody: map[string]interface{}{
				"token": mock.AnythingOfType("string"),
			},
		},
		{
			name: "Invalid request body",
			requestBody: map[string]interface{}{
				"email": "john@example.com",
				// Missing password
			},
			mockSetup:      func(mockService *MockUserService) {},
			expectedStatus: fiber.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "Invalid request",
			},
		},
		{
			name: "Invalid credentials",
			requestBody: map[string]interface{}{
				"email":    "john@example.com",
				"password": "wrongpassword",
			},
			mockSetup: func(mockService *MockUserService) {
				mockService.On("Login", "john@example.com", "wrongpassword").
					Return(nil, errors.New("invalid credentials"))
			},
			expectedStatus: fiber.StatusUnauthorized,
			expectedBody: map[string]interface{}{
				"error": "Invalid credentials",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock service
			mockService := new(MockUserService)
			tt.mockSetup(mockService)

			// Create handler
			handler := &UserHandler{Service: mockService}

			// Create test app
			app := createTestApp()
			app.Post("/login", handler.Login)

			// Create request
			requestBody, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("POST", "/login", bytes.NewReader(requestBody))
			req.Header.Set("Content-Type", "application/json")

			// Make request
			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			// Assert status code
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// Parse response body
			var responseBody map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			assert.NoError(t, err)

			// Assert response body
			for key, expectedValue := range tt.expectedBody {
				if key == "token" {
					// Token should be a non-empty string
					assert.NotEmpty(t, responseBody[key])
				} else {
					assert.Equal(t, expectedValue, responseBody[key])
				}
			}

			// Verify all mocks were called
			mockService.AssertExpectations(t)
		})
	}
}

func TestUserHandler_GetUser(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		mockSetup      func(*MockUserService)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:   "Successful get user",
			userID: "123",
			mockSetup: func(mockService *MockUserService) {
				expectedUser := &user.User{
					ID:    "123",
					Name:  "John Doe",
					Email: "john@example.com",
					Role:  user.RoleEmployee,
				}
				mockService.On("GetUser", "123").Return(expectedUser, nil)
			},
			expectedStatus: fiber.StatusOK,
			expectedBody: map[string]interface{}{
				"id":    "123",
				"name":  "John Doe",
				"email": "john@example.com",
				"role":  "employee",
			},
		},
		{
			name:   "User not found",
			userID: "999",
			mockSetup: func(mockService *MockUserService) {
				mockService.On("GetUser", "999").Return(nil, errors.New("user not found"))
			},
			expectedStatus: fiber.StatusNotFound,
			expectedBody: map[string]interface{}{
				"error": "user not found",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock service
			mockService := new(MockUserService)
			tt.mockSetup(mockService)

			// Create handler
			handler := &UserHandler{Service: mockService}

			// Create test app
			app := createTestApp()
			app.Get("/user/:id", handler.GetUser)

			// Create request
			req := httptest.NewRequest("GET", "/user/"+tt.userID, nil)

			// Make request
			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			// Assert status code
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// Parse response body
			var responseBody map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			assert.NoError(t, err)

			// Assert response body
			for key, expectedValue := range tt.expectedBody {
				if key == "password" {
					// Password should never be returned
					assert.NotContains(t, responseBody, "password")
				} else {
					assert.Equal(t, expectedValue, responseBody[key])
				}
			}

			// Verify all mocks were called
			mockService.AssertExpectations(t)
		})
	}
}

func TestUserHandler_ListUsers(t *testing.T) {
	tests := []struct {
		name           string
		mockSetup      func(*MockUserService)
		expectedStatus int
		expectedBody   []map[string]interface{}
	}{
		{
			name: "Successful list users",
			mockSetup: func(mockService *MockUserService) {
				expectedUsers := []*user.User{
					{
						ID:    "1",
						Name:  "John Doe",
						Email: "john@example.com",
						Role:  user.RoleEmployee,
					},
					{
						ID:    "2",
						Name:  "Jane Smith",
						Email: "jane@example.com",
						Role:  user.RoleAdmin,
					},
				}
				mockService.On("ListUsers").Return(expectedUsers, nil)
			},
			expectedStatus: fiber.StatusOK,
			expectedBody: []map[string]interface{}{
				{
					"id":    "1",
					"name":  "John Doe",
					"email": "john@example.com",
					"role":  "employee",
				},
				{
					"id":    "2",
					"name":  "Jane Smith",
					"email": "jane@example.com",
					"role":  "admin",
				},
			},
		},
		{
			name: "Service error",
			mockSetup: func(mockService *MockUserService) {
				mockService.On("ListUsers").Return(nil, errors.New("database error"))
			},
			expectedStatus: fiber.StatusInternalServerError,
			expectedBody: []map[string]interface{}{
				{
					"error": "database error",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock service
			mockService := new(MockUserService)
			tt.mockSetup(mockService)

			// Create handler
			handler := &UserHandler{Service: mockService}

			// Create test app
			app := createTestApp()
			app.Get("/users", handler.ListUsers)

			// Create request
			req := httptest.NewRequest("GET", "/users", nil)

			// Make request
			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			// Assert status code
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// Parse response body
			if tt.expectedStatus == fiber.StatusOK {
				var responseBody []map[string]interface{}
				err = json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.NoError(t, err)

				// Assert response body
				assert.Equal(t, len(tt.expectedBody), len(responseBody))
				for i, expectedUser := range tt.expectedBody {
					for key, expectedValue := range expectedUser {
						if key == "password" {
							// Password should never be returned
							assert.NotContains(t, responseBody[i], "password")
						} else {
							assert.Equal(t, expectedValue, responseBody[i][key])
						}
					}
				}
			} else {
				var responseBody map[string]interface{}
				err = json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody[0]["error"], responseBody["error"])
			}

			// Verify all mocks were called
			mockService.AssertExpectations(t)
		})
	}
}

func TestUserHandler_UpdateUser(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		requestBody    map[string]interface{}
		mockSetup      func(*MockUserService)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:   "Successful update",
			userID: "123",
			requestBody: map[string]interface{}{
				"name":  "John Updated",
				"email": "john.updated@example.com",
				"role":  "admin",
			},
			mockSetup: func(mockService *MockUserService) {
				expectedUser := &user.User{
					ID:    "123",
					Name:  "John Updated",
					Email: "john.updated@example.com",
					Role:  user.RoleAdmin,
				}
				mockService.On("UpdateUser", expectedUser).Return(nil)
			},
			expectedStatus: fiber.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "User updated",
			},
		},
		{
			name:   "Invalid request body",
			userID: "123",
			requestBody: map[string]interface{}{
				"name": "John Updated",
				// Missing email and role
			},
			mockSetup:      func(mockService *MockUserService) {},
			expectedStatus: fiber.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "Invalid request",
			},
		},
		{
			name:   "Service error",
			userID: "123",
			requestBody: map[string]interface{}{
				"name":  "John Updated",
				"email": "john.updated@example.com",
				"role":  "admin",
			},
			mockSetup: func(mockService *MockUserService) {
				expectedUser := &user.User{
					ID:    "123",
					Name:  "John Updated",
					Email: "john.updated@example.com",
					Role:  user.RoleAdmin,
				}
				mockService.On("UpdateUser", expectedUser).Return(errors.New("user not found"))
			},
			expectedStatus: fiber.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "user not found",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock service
			mockService := new(MockUserService)
			tt.mockSetup(mockService)

			// Create handler
			handler := &UserHandler{Service: mockService}

			// Create test app
			app := createTestApp()
			app.Put("/user/:id", handler.UpdateUser)

			// Create request
			requestBody, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("PUT", "/user/"+tt.userID, bytes.NewReader(requestBody))
			req.Header.Set("Content-Type", "application/json")

			// Make request
			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			// Assert status code
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// Parse response body
			var responseBody map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			assert.NoError(t, err)

			// Assert response body
			for key, expectedValue := range tt.expectedBody {
				assert.Equal(t, expectedValue, responseBody[key])
			}

			// Verify all mocks were called
			mockService.AssertExpectations(t)
		})
	}
}

func TestUserHandler_UpdatePassword(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		requestBody    map[string]interface{}
		mockSetup      func(*MockUserService)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:   "Successful password update",
			userID: "123",
			requestBody: map[string]interface{}{
				"new_password": "newpassword123",
			},
			mockSetup: func(mockService *MockUserService) {
				mockService.On("UpdatePassword", "123", "newpassword123").Return(nil)
			},
			expectedStatus: fiber.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "Password updated",
			},
		},
		{
			name:        "Invalid request body",
			userID:      "123",
			requestBody: map[string]interface{}{
				// Missing new_password
			},
			mockSetup:      func(mockService *MockUserService) {},
			expectedStatus: fiber.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "Invalid request",
			},
		},
		{
			name:   "Service error",
			userID: "123",
			requestBody: map[string]interface{}{
				"new_password": "newpassword123",
			},
			mockSetup: func(mockService *MockUserService) {
				mockService.On("UpdatePassword", "123", "newpassword123").Return(errors.New("user not found"))
			},
			expectedStatus: fiber.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "user not found",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock service
			mockService := new(MockUserService)
			tt.mockSetup(mockService)

			// Create handler
			handler := &UserHandler{Service: mockService}

			// Create test app
			app := createTestApp()
			app.Put("/user/:id/password", handler.UpdatePassword)

			// Create request
			requestBody, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("PUT", "/user/"+tt.userID+"/password", bytes.NewReader(requestBody))
			req.Header.Set("Content-Type", "application/json")

			// Make request
			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			// Assert status code
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// Parse response body
			var responseBody map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			assert.NoError(t, err)

			// Assert response body
			for key, expectedValue := range tt.expectedBody {
				assert.Equal(t, expectedValue, responseBody[key])
			}

			// Verify all mocks were called
			mockService.AssertExpectations(t)
		})
	}
}

func TestUserHandler_DeleteUser(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		mockSetup      func(*MockUserService)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:   "Successful delete",
			userID: "123",
			mockSetup: func(mockService *MockUserService) {
				mockService.On("DeleteUser", "123").Return(nil)
			},
			expectedStatus: fiber.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "User deleted",
			},
		},
		{
			name:   "Service error",
			userID: "123",
			mockSetup: func(mockService *MockUserService) {
				mockService.On("DeleteUser", "123").Return(errors.New("user not found"))
			},
			expectedStatus: fiber.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "user not found",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock service
			mockService := new(MockUserService)
			tt.mockSetup(mockService)

			// Create handler
			handler := &UserHandler{Service: mockService}

			// Create test app
			app := createTestApp()
			app.Delete("/user/:id", handler.DeleteUser)

			// Create request
			req := httptest.NewRequest("DELETE", "/user/"+tt.userID, nil)

			// Make request
			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			// Assert status code
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// Parse response body
			var responseBody map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			assert.NoError(t, err)

			// Assert response body
			for key, expectedValue := range tt.expectedBody {
				assert.Equal(t, expectedValue, responseBody[key])
			}

			// Verify all mocks were called
			mockService.AssertExpectations(t)
		})
	}
}
