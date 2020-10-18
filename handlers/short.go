package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mirusky-dev/url-shortener/storage"
)

// Short create a shorted link
func Short(c *fiber.Ctx) error {
	db := c.Context().UserValue("db").(storage.Storage)
	v := c.FormValue("url")
	id, _ := db.Save(v)
	return c.SendString("http://localhost:3000/" + id)
}
