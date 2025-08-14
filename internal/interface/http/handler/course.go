package handler

import (
	"training-portal/internal/domain/course"
	courseusecase "training-portal/internal/usecase/course"

	"github.com/gofiber/fiber/v2"
)

type CourseHandler struct {
	Service *courseusecase.CourseService
}

var _ = CourseHandler{} // Exported for router.go

// CreateCourse handles POST /course
func (h *CourseHandler) CreateCourse(c *fiber.Ctx) error {
	var req course.Course
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.Service.CreateCourse(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(req)
}

// GetCourse handles GET /course/:id
func (h *CourseHandler) GetCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	course, err := h.Service.GetCourse(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(course)
}

// ListCourses handles GET /courses
func (h *CourseHandler) ListCourses(c *fiber.Ctx) error {
	courses, err := h.Service.ListCourses()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(courses)
}

// UpdateCourse handles PUT /course/:id
func (h *CourseHandler) UpdateCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	var req course.Course
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	req.ID = id
	if err := h.Service.UpdateCourse(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Course updated"})
}

// DeleteCourse handles DELETE /course/:id
func (h *CourseHandler) DeleteCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.Service.DeleteCourse(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Course deleted"})
}
