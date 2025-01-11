package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sujitkumr/go_api/controllers"
)

// RegisterRoutes sets up the application routes
func RegisterRoutes(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
}
