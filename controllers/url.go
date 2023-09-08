package controllers

import (
	"Ushort/shortener"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func CreateShortUrl(c *fiber.Ctx) error {
	payload := PostShortUrl{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseBadRequest{
			ErrorCode:    "E001",
			ErrorMessage: err.Error(),
		})
	}
	shortedUrl, err := shortener.CreateShortUrl(payload.Url)
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

func RedirectUrl(c *fiber.Ctx) error {
	originalUrl := shortener.GetOriginalUrl(strings.Split(c.Path(), "/")[2])
	return c.Redirect(originalUrl)
}
