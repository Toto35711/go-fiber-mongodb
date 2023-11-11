package routes

import (
	"github.com/Toto35711/go-fiber-mongodb/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)      //add this
	app.Get("/user/:userId", controllers.GetAUser) //add this
}
