package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebsvt/ihateapi/handler"
	"github.com/sebsvt/ihateapi/service"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")
	apiv1 := api.Group("/v1")

	// services
	workflowService := service.NewWorkflowService()

	// handlers
	workflowHandler := handler.NewWorkflowHandler(workflowService)

	apiv1.Get("/start/:tool", workflowHandler.Start)
	apiv1.Post("/upload", workflowHandler.Upload)
	apiv1.Post("/process", workflowHandler.Process)
	apiv1.Get("/download/:task", workflowHandler.Download)

	// routes
	app.Listen(":8000")
}
