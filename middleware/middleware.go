package middleware

import "github.com/gofiber/fiber/v2"

type MiddlewareFunc func(*fiber.Ctx) error

type Auth struct {
	middlewareChain []MiddlewareFunc
}

func NewAuth() *Auth {
	return &Auth{}
}

// Use appends a middleware function to the middleware chain
func (a *Auth) Use(mw MiddlewareFunc) *Auth {
	a.middlewareChain = append(a.middlewareChain, mw)
	return a
}

// Apply applies the middleware chain to the given context
func (a *Auth) Apply(c *fiber.Ctx) error {
	for _, mw := range a.middlewareChain {
		if err := mw(c); err != nil {
			return err
		}
	}
	return nil
}
