package handlers

import (
	"context"
	"time"

	"github.com/HemlockPham7/server/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoHandler struct {
	repository models.TodoRepository
}

func NewTodoHandler(router fiber.Router, repository models.TodoRepository) {
	handler := &TodoHandler{
		repository: repository,
	}

	router.Get("/", handler.GetTodos)
	router.Get("/:id", handler.GetTodo)
	router.Post("/", handler.CreateTodo)
	router.Patch("/:id", handler.UpdateTodo)
	router.Delete("/:id", handler.DeleteTodo)
}

func (h *TodoHandler) GetTodos(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	todos, err := h.repository.GetTodos(c)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Failed to fetch todos",
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Todos retrieved successfully",
		"data":    todos,
	})
}

func (h *TodoHandler) GetTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Invalid ID format",
			"data":    nil,
		})
	}

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	todo, err := h.repository.GetTodo(c, objectID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Todo not found",
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Todo retrieved successfully",
		"data":    todo,
	})
}

func (h *TodoHandler) CreateTodo(ctx *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := ctx.BodyParser(todo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Invalid request body",
			"data":    nil,
		})
	}

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createdTodo, err := h.repository.CreateTodo(c, todo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Failed to create todo",
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Todo created successfully",
		"data":    createdTodo,
	})
}

func (h *TodoHandler) UpdateTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Invalid ID format",
			"data":    nil,
		})
	}

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updatedTodo, err := h.repository.UpdateTodo(c, objectID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Failed to update todo",
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Todo updated successfully",
		"data":    updatedTodo,
	})
}

func (h *TodoHandler) DeleteTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Invalid ID format",
			"data":    nil,
		})
	}

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = h.repository.DeleteTodo(c, objectID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Failed to delete todo",
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Todo deleted successfully",
		"data":    nil,
	})
}
