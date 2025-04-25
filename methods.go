package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
	rows, err := pool.Query(ctx, "SELECT id, title, description, status, created_at, updated_at FROM tasks")

	if err != nil {
		return c.Status(500).SendString("Ошибка выполнения запроса к БД")
	}
	defer rows.Close()

	var tasks []Table
	for rows.Next() {
		var task Table
		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.Created_at, &task.Updated_at)
		if err != nil {
			return c.Status(500).SendString("Ошибка сканирования данных")
		}
		tasks = append(tasks, task)
	}
	return c.JSON(tasks)
}

func CreateTask(c *fiber.Ctx) error {
	task := new(Table)
	if err := c.BodyParser(task); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	_, err := pool.Exec(ctx, "INSERT INTO tasks (title, description) VALUES ($1, $2)",
		task.Title, task.Description)
	if err != nil {
		return c.Status(500).SendString("Ошибка вставки задания в БД")
	}

	return c.Status(201).SendString("Задание успешно создано")
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	num_id := strings.Split(id, "=")
	task := new(Table)

	task.Updated_at = time.Now()
	task.Id, _ = strconv.Atoi(num_id[1])

	if err := c.BodyParser(task); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	_, err := pool.Exec(ctx, "UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = $4 WHERE id = $5",
		task.Title, task.Description, task.Status, task.Updated_at, task.Id)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString("Ошибка обновления задания")
	}

	return c.SendString("Задание успешно обновлено")
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	num_id := strings.Split(id, "=")
	int_id, _ := strconv.Atoi(num_id[1])
	_, err := pool.Exec(ctx, "DELETE FROM tasks WHERE id = $1", int_id)
	if err != nil {
		return c.Status(500).SendString("Ошибка удаления задания")
	}

	return c.SendString("Задание успешно удалёно")
}
