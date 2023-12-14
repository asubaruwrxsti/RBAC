package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRequestToken_Correct_Username_Password(t *testing.T) {
	app := fiber.New()
	app.Use(RequestToken())

	req := httptest.NewRequest(http.MethodPost, "/authreq", strings.NewReader("username=test&password=test"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestRequestToken_Incorrect_Username_Password(t *testing.T) {
	app := fiber.New()
	app.Use(RequestToken())

	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader("username=test&password=wrongpassword"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
