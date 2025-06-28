package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gofiber/fiber/v2"
)

const (
	// Define the JWT secret key
	JWT_SECRET = "your_secret_key" // Change this to a secure key
)

// AuthMiddleware checks for a valid JWT token in the request
func AuthMiddleware(c *fiber.Ctx) error {
	// Get the Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is required",
		})
	}

	// Split the header to get the token
	tokenString := strings.Split(authHeader, "Bearer ")[1]
	if tokenString == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bearer token is required",
		})
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(http.StatusUnauthorized, "Invalid token")
		}
		return []byte(JWT_SECRET), nil
	})

	if err != nil || !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	// Token is valid, proceed to the next handler
	return c.Next()
}