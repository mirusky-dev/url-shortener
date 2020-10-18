package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mirusky-dev/url-shortener/storage"
)

// Config ..
type Config struct {
	DB storage.Storage
}

// New ...
func New(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Context().SetUserValue("db", config.DB)
		return c.Next()
	}
}
