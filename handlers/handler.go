package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	// return c.SendString("Hello, World!")
	return c.Render("hello", fiber.Map{
		"Title":       "Go Fiber Template Example",
		"Description": "An example template",
		"Greeting":    "Hello, world!",
	})
}

func IP(c *fiber.Ctx) error {
	ip := c.Context().RemoteIP()
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	// }
	return c.Render("ip", fiber.Map{
		"IP": ip,
	})
}
