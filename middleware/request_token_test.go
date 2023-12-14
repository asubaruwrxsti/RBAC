package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRequestToken(t *testing.T) {
	// Create a Fiber app with the middleware
	app := fiber.New()
	app.Use(RequestToken())

	// Test case 1: Correct username and password
	req1 := httptest.NewRequest(http.MethodPost, "/authreq", strings.NewReader("username=test&password=test"))
	req1.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp1, err := app.Test(req1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp1.StatusCode)

	// Test case 2: Incorrect username and password
	req2 := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader("username=test&password=wrongpassword"))
	req2.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp2, err := app.Test(req2)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp2.StatusCode)
}
