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

// CreateShortUrl doc
// @Summary Create Shorted Url
// @Description Endpoint for create shorted url
// @Tags Url
// @Accept json
// @Produce json
// @Param url body ResponseShortUrl true "Url to short"
// @Success 200 {object} ResponseShortUrl "OK"
// @Failure 400 {object} ResponseBadRequest "Bad request"
// @Router /api/v1/url [post]
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

// RedirectUrl doc
// @Summary Redirect an Url
// @Description Endpoint for redirect shorted url
// @Tags Redirect
// @Accept json
// @Produce json
// @Param urlId path string true "UrlId to redirect"
// @Success 302
// @Failure 404
// @Router /r/{urlId} [get]
func (cu ControllerUrl) RedirectUrl(c *fiber.Ctx) error {
	originalUrl, err := cu.serviceShortener.GetOriginalUrl(strings.Split(c.Path(), "/")[2])
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.Redirect(originalUrl)
}

// RemoveUrl doc
// @Summary Remove Shorted Url
// @Description Endpoint for remove shorted url
// @Tags Url
// @Accept json
// @Produce json
// @Param urlId path string true "UrlId to remove"
// @Success 200
// @Failure 400
// @Router /api/v1/url/{urlId} [delete]
func (cu ControllerUrl) RemoveUrl(c *fiber.Ctx) error {
	urlId := strings.Split(c.Path(), "/")[4]
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

// UpdateUrl doc
// @Summary Update original Url
// @Description Endpoint for update original url
// @Tags Url
// @Accept json
// @Produce json
// @Param url body PostShortUrl true "Url to update"
// @Param urlId path string true "UrlId to update"
// @Success 200
// @Failure 400
// @Router /api/v1/url/{urlId} [put]
func (cu ControllerUrl) UpdateUrl(c *fiber.Ctx) error {
	urlId := strings.Split(c.Path(), "/")[4]
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
