package middleware

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestJWTMiddleware(t *testing.T) {
	// Set up test environment
	os.Setenv("JWT_SECRET", "test-secret-key")

	// Create test app
	app := fiber.New()
	app.Use(cors.New())

	// Protected route
	app.Get("/protected", JWTMiddleware(), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "protected"})
	})

	// Public route
	app.Get("/public", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "public"})
	})

	tests := []struct {
		name           string
		url            string
		headers        map[string]string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Public route without token",
			url:            "/public",
			headers:        map[string]string{},
			expectedStatus: fiber.StatusOK,
			expectedBody:   `{"message":"public"}`,
		},
		{
			name:           "Protected route without token",
			url:            "/protected",
			headers:        map[string]string{},
			expectedStatus: fiber.StatusUnauthorized,
			expectedBody:   `{"error":"Missing or invalid token"}`,
		},
		{
			name:           "Protected route with invalid token",
			url:            "/protected",
			headers:        map[string]string{"Authorization": "Bearer invalid-token"},
			expectedStatus: fiber.StatusUnauthorized,
			expectedBody:   `{"error":"Invalid token"}`,
		},
		{
			name:           "Protected route with malformed header",
			url:            "/protected",
			headers:        map[string]string{"Authorization": "invalid-format"},
			expectedStatus: fiber.StatusUnauthorized,
			expectedBody:   `{"error":"Missing or invalid token"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request
			req := httptest.NewRequest("GET", tt.url, nil)

			// Add headers
			for key, value := range tt.headers {
				req.Header.Set(key, value)
			}

			// Make request
			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			// Assert status code
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// Read response body
			body, err := io.ReadAll(resp.Body)
			assert.NoError(t, err)

			// Assert response body
			assert.Equal(t, tt.expectedBody, string(body))
		})
	}
}

func TestJWTMiddleware_ValidToken(t *testing.T) {
	// Set up test environment
	os.Setenv("JWT_SECRET", "test-secret-key")

	// Create test app
	app := fiber.New()
	app.Use(cors.New())

	// Protected route that returns user info
	app.Get("/protected", JWTMiddleware(), func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(string)
		role := c.Locals("role").(string)
		return c.JSON(fiber.Map{
			"message": "protected",
			"user_id": userID,
			"role":    role,
		})
	})

	// Generate a valid JWT token
	token := generateTestToken(t, "test-user-123", "admin")

	// Test with valid token
	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Assert status code
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Read response body
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	// Assert response body contains user info
	assert.Contains(t, string(body), "test-user-123")
	assert.Contains(t, string(body), "admin")
}

// Helper function to generate a test JWT token
func generateTestToken(t *testing.T, userID, role string) string {
	// This is a simplified token generation for testing
	// In a real application, you would use the actual JWT library
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("test-secret-key"))
	assert.NoError(t, err)

	return tokenString
}

// Helper function to create test requests
func createTestRequest(method, url string, headers map[string]string) *http.Request {
	req := httptest.NewRequest(method, url, nil)
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return req
}
