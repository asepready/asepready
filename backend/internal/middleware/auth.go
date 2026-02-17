package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware validates JWT tokens
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: Implement JWT validation
		// 1. Extract token from Authorization header
		// 2. Validate token
		// 3. Extract user ID from token
		// 4. Store user ID in context
		// 5. Call next handler
		
		// For now, just pass through
		return c.Next()
	}
}
