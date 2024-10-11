package routes

import "github.com/gofiber/fiber"

func Init(app *fiber.App) {
	InitTodoRoutes(app)
}
