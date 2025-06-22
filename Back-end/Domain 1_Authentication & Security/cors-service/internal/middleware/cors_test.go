package middleware

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCORSMiddleware_AllowedOrigin(t *testing.T) {
	// Setup
	allowedOrigins := []string{"http://localhost:3000", "https://veterinary-front.com"}

	// Mock request with valid Origin
	req := httptest.NewRequest("GET", "/validate-cors", nil)
	req.Header.Set("Origin", "http://localhost:3000")

	w := httptest.NewRecorder()

	r := gin.New()
	r.Use(CORSMiddleware(allowedOrigins))
	r.GET("/validate-cors", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "CORS origin allowed"})
	})

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "http://localhost:3000", w.Header().Get("Access-Control-Allow-Origin"))
	assert.True(t, strings.Contains(w.Body.String(), "CORS origin allowed"))
}

func TestCORSMiddleware_DisallowedOrigin(t *testing.T) {
	allowedOrigins := []string{"http://localhost:3000", "https://veterinary-front.com"}

	req := httptest.NewRequest("GET", "/validate-cors", nil)
	req.Header.Set("Origin", "http://malicious-site.com")

	w := httptest.NewRecorder()

	r := gin.New()
	r.Use(CORSMiddleware(allowedOrigins))
	r.GET("/validate-cors", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "CORS origin allowed"})
	})

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Origin"))
	assert.True(t, strings.Contains(w.Body.String(), "CORS origin allowed"))
}
