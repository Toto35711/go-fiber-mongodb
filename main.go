package main

import (
	"github.com/Toto35711/go-fiber-mongodb/configs"
	"github.com/Toto35711/go-fiber-mongodb/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()
	port := configs.EnvPort()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})

	routes.UserRoute(app)

	app.Listen(":" + port)
}
