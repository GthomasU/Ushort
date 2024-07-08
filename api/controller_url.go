package api

import (
	"Ushort/services"
	"Ushort/storage"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ControllerUrl struct {
	serviceShortener services.ServiceShortener
}

func NewControllerUrl() ControllerUrl {
	urlGenerator := services.NewUrlGenerator()
	redisClient := storage.NewRedisClient()
	serviceShortener := services.NewServiceShortener(urlGenerator, true, redisClient)
	return ControllerUrl{serviceShortener: serviceShortener}

}
func (cu ControllerUrl) CreateShortUrl(c *fiber.Ctx) error {
	payload := PostShortUrl{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseBadRequest{
			ErrorCode:    "E001",
			ErrorMessage: err.Error(),
		})
	}
	shortedUrl, err := cu.serviceShortener.CreateShortUrl(payload.Url)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseBadRequest{
			ErrorCode:    "E002",
			ErrorMessage: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(ResponseShortUrl{
		ShortedUrl: *shortedUrl,
	})
}

func (cu ControllerUrl) RedirectUrl(c *fiber.Ctx) error {
	originalUrl, err := cu.serviceShortener.GetOriginalUrl(strings.Split(c.Path(), "/")[2])
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.Redirect(originalUrl)
}

func (cu ControllerUrl) RemoveUrl(c *fiber.Ctx) error {
	urlId := strings.Split(c.Path(), "/")[2]
	result, err := cu.serviceShortener.RemoveOriginalUrl(urlId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if result {
		return c.SendStatus(fiber.StatusOK)
	} else {
		return c.SendStatus(fiber.StatusBadRequest)
	}
}

func (cu ControllerUrl) UpdateUrl(c *fiber.Ctx) error {
	urlId := strings.Split(c.Path(), "/")[2]
	payload := PostShortUrl{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseBadRequest{
			ErrorCode:    "E001",
			ErrorMessage: err.Error(),
		})
	}
	result, err := cu.serviceShortener.UpdateOriginalUrl(urlId, payload.Url)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if result {
		return c.SendStatus(fiber.StatusOK)
	}
	return c.SendStatus(fiber.StatusBadRequest)
}
