package main

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterProductRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/gettasks", GetTasks) // Получить все продукты
	api.Post("/tasks", CreateTask) // Создать новый продукт
	// api.Get("/products/:id", handlers.GetProduct)   // Получить продукт по ID
	// api.Put("/products/:id", handlers.UpdateProduct) // Обновить продукт
	// api.Delete("/products/:id", handlers.DeleteProduct) // Удалить продукт
}
