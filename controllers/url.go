package controllers

import (
	"Ushort/shortener"

	"github.com/gofiber/fiber/v2"
)

func CreateShortUrl(c *fiber.Ctx) error {
	return c.SendString(shortener.CreateShortUrl(""))
}

func RedirectUrl(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
