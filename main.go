package main

import (
	"github.com/gofiber/fiber/v2"
	h "github.com/mirusky-dev/url-shortener/handlers"
	"github.com/mirusky-dev/url-shortener/middlewares"
	"github.com/mirusky-dev/url-shortener/storage"
)

func main() {
	app := fiber.New()

	app.Use(middlewares.New(middlewares.Config{
		DB: storage.NewInMem(),
	}))

	app.Get("/ping", h.Ping)

	app.Post("/short", h.Short)
	app.Get("/:id", h.Shorted)

	app.Listen(":3000")
}
