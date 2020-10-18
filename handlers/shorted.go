package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mirusky-dev/url-shortener/storage"
)

// Shorted handles a shorted link
func Shorted(c *fiber.Ctx) error {
	s := c.Params("id")
	db := c.Context().UserValue("db").(storage.Storage)
	v, err := db.Get(s)
	if err != nil {
		return c.SendString("Not Found")
	}
	return c.SendString(v)
}
