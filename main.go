package main

import (
	"github.com/Toto35711/go-fiber-mongodb/configs"
	"github.com/Toto35711/go-fiber-mongodb/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})

	//routes
	routes.UserRoute(app) //add this

	app.Listen(":6000")
}
