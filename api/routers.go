package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // swagger handler
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
	groupV1.Delete("/url/*", s.controllerUrl.RemoveUrl)
	groupV1.Put("/url/*", s.controllerUrl.UpdateUrl)
	groupV1.Post("/url", s.controllerUrl.CreateShortUrl)
	groupRedirect := s.app.Group("r/")
	groupRedirect.Get("/*", s.controllerUrl.RedirectUrl)

	groupSwagger := s.app.Group("/swagger")
	groupSwagger.Get("/*", swagger.HandlerDefault)
}

func (s Server) StartListening() {
	_ = s.app.Listen(":3000")
}
