package routers

import (
	"github.com/gofiber/fiber/v2"
	"ifconfig/handlers"
)

func Hello(router fiber.Router) {
	router.Get("/", handlers.Hello)
}

func IP(router fiber.Router) {
	router.Get("/", handlers.IP)
}
