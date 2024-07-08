package api

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app           *fiber.App
	ControllerUrl ControllerUrl
}

func NewServer() Server {
	app := fiber.New()
	server := Server{
		app: app,
	}
	server.configRoutes()
	return server

}
func (s Server) configRoutes() {
	groupV1 := s.app.Group("api/v1/")
	groupRedirect := s.app.Group("r/")
	groupRedirect.Get("/*", s.ControllerUrl.RedirectUrl)
	groupRedirect.Delete("/*", s.ControllerUrl.RemoveUrl)
	groupRedirect.Put("/*", s.ControllerUrl.UpdateUrl)
	groupV1.Post("/url", s.ControllerUrl.CreateShortUrl)
}

func (s Server) StartListening() {
	_ = s.app.Listen(":3000")
}
