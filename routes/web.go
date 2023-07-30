package routes

import (
	controllers "miniproject/controllers"

	"github.com/gofiber/fiber/v2"
)

func WebRoutes(app *fiber.App) {
	app.Get("/", controllers.Index)
	app.Get("/login", controllers.Login)
	app.Get("/register", controllers.Register)
	app.Get("/account", controllers.Account)
}