package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuhari7/superbank_assessment/pkg"
)

func JWTMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if len(token) < 7 || token[:7] != "Bearer " {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid token"})
	}

	userID, err := pkg.ValidateJWT(token[7:])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	c.Locals("user_id", userID)
	return c.Next()
}
