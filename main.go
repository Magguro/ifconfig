package main

import (
	"log"

	"ifconfig/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	main := app.Group("/")
	main.Route("/hello", routers.Hello)
	main.Route("/ip", routers.IP)

	log.Fatal(app.Listen(":8000"))
}
