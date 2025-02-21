package handler

import (
	"bytes"

	"github.com/gofiber/fiber/v2"
	"github.com/sebsvt/ihateapi/model"
	"github.com/sebsvt/ihateapi/service"
)

type workflowHandler struct {
	workflowService service.WorkflowService
}

func NewWorkflowHandler(workflowService service.WorkflowService) workflowHandler {
	return workflowHandler{
		workflowService: workflowService,
	}
}

func (h *workflowHandler) Start(c *fiber.Ctx) error {
	tool := c.Params("tool")
	res, err := h.workflowService.Start(tool)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(res)
}

func (h *workflowHandler) Upload(c *fiber.Ctx) error {
	res, err := h.workflowService.Upload()
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(res)
}

func (h *workflowHandler) Process(c *fiber.Ctx) error {
	var req model.ProcessWorkFlowRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, err := h.workflowService.Process(req)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(res)
}

func (h *workflowHandler) Download(c *fiber.Ctx) error {
	task := c.Params("task")
	// return as byte of file
	res, err := h.workflowService.Download(task)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.SendStream(bytes.NewReader(res))
}
