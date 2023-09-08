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
	groupV1 := app.Group("api/v1/")
	groupRedirect := app.Group("r/")
	groupRedirect.Get("/*", controllers.RedirectUrl)
	groupV1.Post("/url", controllers.CreateShortUrl)
}

func StartListening() {
	app.Listen(":3000")
}
