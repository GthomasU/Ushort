package server

import (
	"Ushort/controllers"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func init() {
	app = fiber.New()
	configRoutes()
}
func configRoutes() {
	v1 := app.Group("api/v1/")
	v1.Post("/url", controllers.CreateShortUrl)
	v1.Get("/url/*", controllers.RedirectUrl)
}

func StartListening() {
	app.Listen(":3000")
}
