package api

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app           *fiber.App
	controllerUrl ControllerUrl
}

func NewServer() Server {
	app := fiber.New()
	controllerUrl := NewControllerUrl()
	server := Server{
		controllerUrl: controllerUrl,
		app:           app,
	}
	server.configRoutes()
	return server

}
func (s Server) configRoutes() {
	groupV1 := s.app.Group("api/v1/")
	groupRedirect := s.app.Group("r/")
	groupRedirect.Get("/*", s.controllerUrl.RedirectUrl)
	groupRedirect.Delete("/*", s.controllerUrl.RemoveUrl)
	groupRedirect.Put("/*", s.controllerUrl.UpdateUrl)
	groupV1.Post("/url", s.controllerUrl.CreateShortUrl)
}

func (s Server) StartListening() {
	_ = s.app.Listen(":3000")
}
