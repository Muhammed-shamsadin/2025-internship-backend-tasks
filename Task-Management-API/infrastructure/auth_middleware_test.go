package infrastructure

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Helper function to create a test server with the middleware
	setupServer := func() *gin.Engine {
		r := gin.Default()
		r.Use(AuthMiddleware())
		r.GET("/test", func(c *gin.Context) {
			userID, exists := c.Get("userID")
			assert.True(t, exists)
			assert.NotEmpty(t, userID)
			c.JSON(http.StatusOK, gin.H{"message": "passed"})
		})
		return r
	}

	t.Run("success_valid_token", func(t *testing.T) {
		router := setupServer()
		w := httptest.NewRecorder()

		// Generate a valid token
		testUserID := "test-user-123"
		token, err := GenerateJWT(testUserID)
		assert.NoError(t, err)

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "passed")
	})

	t.Run("no_authorization_header", func(t *testing.T) {
		router := setupServer()
		w := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "Authorization header is required")
	})

	t.Run("invalid_header_format", func(t *testing.T) {
		router := setupServer()
		w := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "invalid-token-format")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "Authorization header format must be Bearer {token}")
	})

	t.Run("invalid_token", func(t *testing.T) {
		router := setupServer()
		w := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "Bearer an-invalid-jwt-token")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid token")
	})

	t.Run("expired_token", func(t *testing.T) {
		// Temporarily shorten the token lifespan for this test
		originalExpiry := 24 * time.Hour
		jwtExpiry = -1 * time.Hour                    // Set expiry to the past
		defer func() { jwtExpiry = originalExpiry }() // Reset it after the test

		router := setupServer()
		w := httptest.NewRecorder()

		expiredToken, err := GenerateJWT("test-user-expired")
		assert.NoError(t, err)

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "Bearer "+expiredToken)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid token")
	})
}
