package routes

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"github.com/naythukhant/gofiber-todos/models"
	"github.com/naythukhant/gofiber-todos/services"
	"gorm.io/gorm"
)

func InitTodoRoutes(app *fiber.App) {
	todos := app.Group("/todos")

	todos.Get("/", index)
	todos.Post("/", store)
	todos.Get("/:id", show)
	todos.Put("/:id", update)
	todos.Delete("/:id", delete)
}

func index(ctx *fiber.Ctx) {
	var todos []models.Todo
	result := services.DATABASE.Find(&todos)

	if result.Error != nil {
		ctx.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		return
	}

	ctx.JSON(todos)
}

func store(ctx *fiber.Ctx) {
	todo := new(models.Todo)

	if err := ctx.BodyParser(todo); err != nil {
		ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(todo); err != nil {
		ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
		return
	}

	result := services.DATABASE.Create(todo)

	if result.Error != nil {
		ctx.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		return
	}

	ctx.Status(201).JSON(todo)
}

func show(ctx *fiber.Ctx) {
	var todo models.Todo
	result := services.DATABASE.First(&todo, ctx.Params("id"))

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.Status(404).JSON(fiber.Map{"error": "Record not found!"})
			return
		}

		ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
		return
	}

	ctx.JSON(todo)
}

func update(ctx *fiber.Ctx) {
	var todo models.Todo
	result := services.DATABASE.First(&todo, ctx.Params("id"))

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.Status(404).JSON(fiber.Map{"error": "Record not found!"})
			return
		}

		ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
		return
	}

	var updateRequest models.UpdateTodoRequest
	if err := ctx.BodyParser(&updateRequest); err != nil {
		ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(updateRequest); err != nil {
		ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
		return
	}

	todo.Title = updateRequest.Title
	todo.Description = updateRequest.Description

	if updateRequest.CompletedAt != nil {
		now := time.Now()
		todo.CompletedAt = &now
	} else {
		todo.CompletedAt = nil
	}

	if err := services.DATABASE.Save(&todo).Error; err != nil {
		ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
		return
	}

	ctx.JSON(todo)
}

func delete(ctx *fiber.Ctx) {
	var todo models.Todo
	result := services.DATABASE.First(&todo, ctx.Params("id"))

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.Status(404).JSON(fiber.Map{"error": "Record not found!"})
			return
		}

		ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
		return
	}

	if err := services.DATABASE.Delete(&todo).Error; err != nil {
		ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
		return
	}

	ctx.Status(204).Send(nil)
}
