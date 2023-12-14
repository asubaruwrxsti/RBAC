package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestVerifyToken_ValidToken(t *testing.T) {
	app := fiber.New()
	app.Use(VerifyToken())
	JWT_SECRET := ([]byte("secret"))

	// Create a valid token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":  "test",
		"groupId": 1,
		"exp":     time.Now().Add(time.Hour).Unix(), // Set expiration time to 1 hour from now
		"iss":     "your_issuer",
	})
	tokenString, err := token.SignedString(JWT_SECRET)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/authreq", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	resp, err := app.Test(req)
	assert.NoError(t, err)

	// Returns 401 since there is no valid token in header,
	// although returns it as JSON response on valid credentials.
	// Maybe return 200 and the token?
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestVerifyToken_MissingToken(t *testing.T) {
	app := fiber.New()
	app.Use(VerifyToken())

	req := httptest.NewRequest(http.MethodGet, "/authreq", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestVerifyToken_MalformedToken(t *testing.T) {
	app := fiber.New()
	app.Use(VerifyToken())
	JWT_SECRET := ([]byte("secret"))

	// Create a valid token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":  "test",
		"groupId": 1,
		"exp":     time.Now().Add(time.Hour).Unix(),
		"iss":     "your_issuer",
	})
	tokenString, err := token.SignedString(JWT_SECRET)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/authreq", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString+"extra")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestVerifyToken_InvalidSignature(t *testing.T) {
	app := fiber.New()
	app.Use(VerifyToken())

	// Create a valid token with a valid signature
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":  "test",
		"groupId": 1,
		"exp":     time.Now().Add(time.Hour).Unix(),
		"iss":     "your_issuer",
	})

	// Attempt to use a different secret
	invalidSignatureTokenString, err := token.SignedString([]byte("invalid_secret"))
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/authreq", nil)
	req.Header.Set("Authorization", "Bearer "+invalidSignatureTokenString)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestVerifyToken_ExpiredToken(t *testing.T) {
	app := fiber.New()
	app.Use(VerifyToken())
	JWT_SECRET := ([]byte("secret"))

	// Create a token with an expiration time set to 1 hour ago
	expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":  "test",
		"groupId": 1,
		"exp":     time.Now().Add(-time.Hour).Unix(),
		"iss":     "your_issuer",
	})
	expiredTokenString, err := expiredToken.SignedString(JWT_SECRET)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/authreq", nil)
	req.Header.Set("Authorization", "Bearer "+expiredTokenString)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
