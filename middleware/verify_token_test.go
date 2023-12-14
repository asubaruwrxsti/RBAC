package middleware

import (
	"RBAC/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestVerifyToken(t *testing.T) {
	app := fiber.New()
	app.Use(VerifyToken())
	JWT_SECRET := ([]byte(config.Config("JWT_SECRET")))

	// Create a valid token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":  "test",
		"groupId": 1,
		"exp":     0,
		"iss":     "your_issuer",
	})
	tokenString, err := token.SignedString(JWT_SECRET)
	assert.NoError(t, err)

	// Test case 1: Valid token
	req1 := httptest.NewRequest(http.MethodGet, "/authreq", nil)
	req1.Header.Set("Authorization", "Bearer"+tokenString)
	resp1, err := app.Test(req1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp1.StatusCode)

	// Test case 2: Missing token
	req2 := httptest.NewRequest(http.MethodGet, "/authreq", nil)
	resp2, err := app.Test(req2)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp2.StatusCode)

	// Test case 3: Malformed token
	req3 := httptest.NewRequest(http.MethodGet, "/authreq", nil)
	req3.Header.Set("Authorization", "Bearer"+tokenString+"extra")
	resp3, err := app.Test(req3)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp3.StatusCode)

	// Test case 4: Invalid signature
	invalidSignatureToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":  "test",
		"groupId": 1,
		"exp":     0,
		"iss":     "your_issuer",
	})

	invalidSignatureTokenString, err := invalidSignatureToken.SignedString([]byte("invalid_secret"))
	assert.NoError(t, err)

	req4 := httptest.NewRequest(http.MethodGet, "/authreq", nil)
	req4.Header.Set("Authorization", "Bearer"+invalidSignatureTokenString)
	resp4, err := app.Test(req4)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp4.StatusCode)

	// Test case 5: Expired token or not active yet
	expired_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":  "test",
		"groupId": 1,
		"exp":     0,
		"iss":     "your_issuer",
	})

	expired_token_string, err := expired_token.SignedString(JWT_SECRET)
	assert.NoError(t, err)

	req5 := httptest.NewRequest(http.MethodGet, "/authreq", nil)
	req5.Header.Set("Authorization", "Bearer"+expired_token_string)
	resp5, err := app.Test(req5)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp5.StatusCode)
}
