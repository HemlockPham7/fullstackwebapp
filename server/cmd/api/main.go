package main

import (
	"context"
	"fmt"

	"github.com/HemlockPham7/server/config"
	"github.com/HemlockPham7/server/db"
	"github.com/HemlockPham7/server/handlers"
	"github.com/HemlockPham7/server/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	db, client := db.Init(envConfig)
	defer client.Disconnect(context.Background())

	app := fiber.New()

	// Repositories
	todoRepository := repositories.NewTodoRepository(db)

	// Routing
	server := app.Group("/api")
	handlers.NewTodoHandler(server.Group("/todos"), todoRepository)

	// start
	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
