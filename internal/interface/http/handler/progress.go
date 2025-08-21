package handler

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

// Progress represents a user's progress on a course or quiz.
// In a real implementation, these would map to database tables.
type Progress struct {
	UserID           string         `json:"user_id"`
	CourseID         string         `json:"course_id"`
	CompletedModules []string       `json:"completed_modules"`
	CompletedQuizzes map[string]int `json:"completed_quizzes"` // quizID -> score
}

// ProgressHandler provides HTTP handlers for tracking progress.
// Currently uses in-memory storage for demo purposes.
type ProgressHandler struct {
	mu       sync.Mutex
	progress map[string]*Progress // key = userID:courseID
}

// NewProgressHandler creates a new ProgressHandler.
// In the future, inject a database repository/service for persistent storage.
func NewProgressHandler() *ProgressHandler {
	return &ProgressHandler{
		progress: make(map[string]*Progress),
	}
}

// GetUserProgress handles GET /user/:id/progress
// Returns the progress of a specific user across all courses.
func (h *ProgressHandler) GetUserProgress(c *fiber.Ctx) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	userID := c.Params("id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing user ID"})
	}

	var result []Progress
	for _, p := range h.progress {
		if p.UserID == userID {
			result = append(result, *p) // Replace with DB query in the future
		}
	}

	return c.JSON(result)
}

// UpdateUserProgress handles PUT /user/:id/progress
// Updates user progress for a specific course.
func (h *ProgressHandler) UpdateUserProgress(c *fiber.Ctx) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	userID := c.Params("id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing user ID"})
	}

	var input Progress
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Composite key userID:courseID for in-memory storage
	key := userID + ":" + input.CourseID

	h.progress[key] = &input // Replace with DB INSERT/UPDATE in the future

	return c.JSON(fiber.Map{
		"message":  "User progress updated successfully",
		"progress": input,
	})
}

// GetCourseProgress handles GET /course/:id/progress
// Returns progress analytics for a course across all users.
func (h *ProgressHandler) GetCourseProgress(c *fiber.Ctx) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	courseID := c.Params("id")
	if courseID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing course ID"})
	}

	var courseProgress []Progress
	for _, p := range h.progress {
		if p.CourseID == courseID {
			courseProgress = append(courseProgress, *p) // Replace with DB query in the future
		}
	}

	// Example analytics: count of users who have completed at least one module
	completedCount := 0
	for _, p := range courseProgress {
		if len(p.CompletedModules) > 0 { // Simplistic check; replace with real logic
			completedCount++
		}
	}

	return c.JSON(fiber.Map{
		"course_id":       courseID,
		"user_count":      len(courseProgress),
		"completed_count": completedCount,
		"progress":        courseProgress,
	})
}
