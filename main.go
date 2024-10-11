package main

import (
	"github.com/gofiber/fiber"
	"github.com/naythukhant/gofiber-todos/bootstrap"
	"github.com/naythukhant/gofiber-todos/configs"
	"github.com/naythukhant/gofiber-todos/routes"
)

func main() {
	bootstrap.Bootstrap()

	app := fiber.New()
	routes.Init(app)

	app.Listen(configs.App.PORT)
}
