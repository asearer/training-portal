package handler

import (
	"training-portal/internal/domain/course"
	moduleusecase "training-portal/internal/usecase/course"

	"github.com/gofiber/fiber/v2"
)

type ModuleHandler struct {
	Service *moduleusecase.ModuleService
}

var _ = ModuleHandler{} // Exported for router.go

// CreateModule handles POST /module
func (h *ModuleHandler) CreateModule(c *fiber.Ctx) error {
	var req course.Module
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.Service.CreateModule(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(req)
}

// GetModule handles GET /module/:id
func (h *ModuleHandler) GetModule(c *fiber.Ctx) error {
	id := c.Params("id")
	module, err := h.Service.GetModule(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(module)
}

// ListModulesByCourse handles GET /course/:course_id/modules
func (h *ModuleHandler) ListModulesByCourse(c *fiber.Ctx) error {
	courseID := c.Params("course_id")
	modules, err := h.Service.ListModulesByCourse(courseID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(modules)
}

// UpdateModule handles PUT /module/:id
func (h *ModuleHandler) UpdateModule(c *fiber.Ctx) error {
	id := c.Params("id")
	var req course.Module
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	req.ID = id
	if err := h.Service.UpdateModule(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Module updated"})
}

// DeleteModule handles DELETE /module/:id
func (h *ModuleHandler) DeleteModule(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.Service.DeleteModule(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Module deleted"})
}
