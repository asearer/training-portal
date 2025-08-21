package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Enrollment represents a user's enrollment in a course.
type Enrollment struct {
	ID       string `json:"id"`
	UserID   string `json:"userId"`
	CourseID string `json:"courseId"`
}

// EnrollmentHandler provides HTTP handlers for enrollment-related endpoints.
type EnrollmentHandler struct{}

// In-memory storage for enrollments (replace with DB in the future)
var enrollments = []Enrollment{}

// EnrollUser handles POST /enroll
func (h *EnrollmentHandler) EnrollUser(c *fiber.Ctx) error {
	var req struct {
		UserID   string `json:"userId"`
		CourseID string `json:"courseId"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Check if user is already enrolled
	for _, e := range enrollments {
		if e.UserID == req.UserID && e.CourseID == req.CourseID {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User already enrolled in this course"})
		}
	}

	enrollment := Enrollment{
		ID:       strconv.Itoa(len(enrollments) + 1),
		UserID:   req.UserID,
		CourseID: req.CourseID,
	}
	enrollments = append(enrollments, enrollment)

	return c.Status(fiber.StatusCreated).JSON(enrollment)
}

// UnenrollUser handles POST /unenroll
func (h *EnrollmentHandler) UnenrollUser(c *fiber.Ctx) error {
	var req struct {
		UserID   string `json:"userId"`
		CourseID string `json:"courseId"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	for i, e := range enrollments {
		if e.UserID == req.UserID && e.CourseID == req.CourseID {
			// Remove enrollment
			enrollments = append(enrollments[:i], enrollments[i+1:]...)
			return c.JSON(fiber.Map{"message": "User unenrolled successfully"})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Enrollment not found"})
}

// ListEnrollments handles GET /enrollments
func (h *EnrollmentHandler) ListEnrollments(c *fiber.Ctx) error {
	userID := c.Query("userId")
	courseID := c.Query("courseId")

	var filtered []Enrollment
	for _, e := range enrollments {
		if (userID == "" || e.UserID == userID) && (courseID == "" || e.CourseID == courseID) {
			filtered = append(filtered, e)
		}
	}

	return c.JSON(filtered)
}
