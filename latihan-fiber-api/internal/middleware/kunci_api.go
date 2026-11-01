package middleware

import (
	"os"

	"github.com/gofiber/fiber/v3"
)

func KunciApi() fiber.Handler {
	apiKeyBenar := os.Getenv("API_KEY")

	return func(c fiber.Ctx) error {
		apiKey := c.Get("X-Api-Key")

		if apiKeyBenar == "" || apiKey != apiKeyBenar {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"sukses": false,
				"pesan":  "API Key tidak valid",
			})
		}

		return c.Next()
	}
}
