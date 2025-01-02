package provider

import "github.com/gofiber/fiber/v2"

type IHandler interface {
	Handle(c *fiber.Ctx) error
}
