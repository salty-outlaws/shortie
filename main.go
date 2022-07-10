package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// CRUD user
	app.Get("/users", func(c *fiber.Ctx) error { return nil })
	app.Get("/users/:id", func(c *fiber.Ctx) error { return nil })
	app.Put("/users/:id", func(c *fiber.Ctx) error { return nil })
	app.Delete("/users/:id", func(c *fiber.Ctx) error { return nil })

	// CRUD token
	app.Get("/tokens", func(c *fiber.Ctx) error { return nil })
	app.Get("/tokens/:id", func(c *fiber.Ctx) error { return nil })
	app.Put("/tokens/:id", func(c *fiber.Ctx) error { return nil })
	app.Delete("/tokens/:id", func(c *fiber.Ctx) error { return nil })

	// CRUD endpoint
	app.Get("/e", func(c *fiber.Ctx) error { return nil })
	app.Get("/e/:id", func(c *fiber.Ctx) error { return nil })
	app.Put("/e/:id", func(c *fiber.Ctx) error { return nil })
	app.Delete("/e/:id", func(c *fiber.Ctx) error { return nil })

	// trigger endpoint
	app.Get("/u/:code", func(c *fiber.Ctx) error {
		return c.Redirect(string(""))
	})

	// status
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"status": "ok"})
	})

	app.Listen(":8000")
}
