package main

import (
	"fmt"

	"github.com/alextavella/multithreading/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Listening: http://localhost:8080")
	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("I'm good!")
	})
	app.Get("/address/zipcode/:zipcode", handler.NewAddressHandler().Handle)
	app.Listen(":8080")
}
