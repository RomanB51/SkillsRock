package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool
var ctx = context.Background()

func main() {
	var err error
	if err = Conn_to_DB(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Get("/tasks", GetTasks)

	app.Post("/tasks", CreateTask)

	app.Put("/tasks/:id", UpdateTask)

	app.Delete("/tasks/:id", DeleteTask)

	app.Listen(":3000")

}
