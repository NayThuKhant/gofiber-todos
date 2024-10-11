package main

import (
	"fmt"

	"github.com/naythukhant/gofiber-todos/bootstrap"
	"github.com/naythukhant/gofiber-todos/configs"
	"github.com/naythukhant/gofiber-todos/models"
	"github.com/naythukhant/gofiber-todos/services"
)

func main() {
	bootstrap.Bootstrap()

	if configs.IsProduction() {
		panic("You can not setup the database schema in production mode!")
	}

	setupTodoSchema()

	fmt.Println("Database schema has been setup")
}

func setupTodoSchema() {
	if err := services.DATABASE.Migrator().DropTable(&models.Todo{}); err != nil {
		panic("Failed to drop the existing Todo table: " + err.Error())
	}

	if err := services.DATABASE.AutoMigrate(&models.Todo{}); err != nil {
		panic("Failed to drop the existing Todo table: " + err.Error())
	}
}
