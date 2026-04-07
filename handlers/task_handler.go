package handlers

import (
	"strconv"

	"stability-test-task-api/models"
	"stability-test-task-api/store"

	"github.com/gofiber/fiber/v2"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func GetTasks(c *fiber.Ctx) error {
	tasks := store.GetAllTasks()
	return c.JSON(tasks)
}

func GetTask(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, _ := strconv.Atoi(idParam)

	task := store.GetTaskByID(id)

	if task == nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "task not found",
		})
	}

	return c.JSON(task)
}

func CreateTask(c *fiber.Ctx) error {
	var task models.Task

	// cek json
	if err := c.BodyParser(&task); err != nil {
		// return 400 bad request, karena format format input tidak sesuai dengan struct Task
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid JSON",
		})
	}

	// validasi inputan berdasarkan struct Task
	if err := validate.Struct(task); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		errorLists := []string{}
		// append setiap pesan error ke arr errorLists
		for _, e := range validationErrors {
			errorLists = append(errorLists, e.Field()+" is "+e.Tag())
		}

		// return pesan error pertama saja agar tidak mengubah format response (errors: msg), dengan status 400 bad request
		return c.Status(400).JSON(fiber.Map{
			"errors": errorLists[0],
		})
	}

	// store
	store.AddTask(task)

	return c.JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	var task models.Task

	// cek json
	if err := c.BodyParser(&task); err != nil {
		// return 400 bad request, karena format format input tidak sesuai dengan struct Task
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid JSON",
		})
	}

	// validasi inputan berdasarkan struct Task
	if err := validate.Struct(task); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		errorLists := []string{}
		// append setiap pesan error ke arr errorLists
		for _, e := range validationErrors {
			errorLists = append(errorLists, e.Field()+" is "+e.Tag())
		}

		// return pesan error pertama saja agar tidak mengubah format response (errors: msg), dengan status 400 bad request
		return c.Status(400).JSON(fiber.Map{
			"errors": errorLists[0],
		})
	}

	// store
	store.EditTask(task)

	return c.JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, _ := strconv.Atoi(idParam)

	store.DeleteTask(id)

	return c.JSON(fiber.Map{
		"message": "deleted",
	})
}
