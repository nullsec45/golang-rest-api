package main

import (
	"github.com/gofiber/fiber/v2"
	bookController "github.com/nullsec45/golang-rest-api/contollers"
	"github.com/nullsec45/golang-rest-api/models"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	api := app.Group("/api")
	book := api.Group("/books")

	book.Get("/", bookController.Index)
	book.Get("/:id", bookController.Show)
	book.Post("/", bookController.Create)
	book.Put("/:id", bookController.Update)
	book.Delete("/:id", bookController.Delete)
	app.Listen(":8000")

}
