package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// UserEngagement represents a stub for user engagement metrics
type UserEngagement struct {
	UserID       string `json:"userId"`
	LastLogin    string `json:"lastLogin"`
	CoursesTaken int    `json:"coursesTaken"`
}

// CourseAnalytics represents a stub for course-level analytics
type CourseAnalytics struct {
	CourseID       string `json:"courseId"`
	EnrolledUsers  int    `json:"enrolledUsers"`
	CompletedUsers int    `json:"completedUsers"`
	AverageScore   int    `json:"averageScore"`
}

// AnalyticsEvent represents a stub for analytics events
type AnalyticsEvent struct {
	EventID   string    `json:"eventId"`
	UserID    string    `json:"userId"`
	EventType string    `json:"eventType"`
	Timestamp time.Time `json:"timestamp"`
}

// AnalyticsHandler provides HTTP handlers for analytics endpoints.
type AnalyticsHandler struct{}

// In-memory stub data (replace with real DB/analytics service in the future)
var (
	userEngagements = []UserEngagement{
		{UserID: "1", LastLogin: time.Now().AddDate(0, 0, -1).Format(time.RFC3339), CoursesTaken: 3},
		{UserID: "2", LastLogin: time.Now().AddDate(0, 0, -2).Format(time.RFC3339), CoursesTaken: 5},
	}

	courseAnalytics = []CourseAnalytics{
		{CourseID: "101", EnrolledUsers: 50, CompletedUsers: 30, AverageScore: 85},
		{CourseID: "102", EnrolledUsers: 40, CompletedUsers: 20, AverageScore: 78},
	}

	analyticsEvents = []AnalyticsEvent{
		{EventID: "1", UserID: "1", EventType: "login", Timestamp: time.Now().Add(-48 * time.Hour)},
		{EventID: "2", UserID: "2", EventType: "course_completed", Timestamp: time.Now().Add(-24 * time.Hour)},
	}
)

// GetUserEngagement handles GET /api/analytics/user-engagement
func (h *AnalyticsHandler) GetUserEngagement(c *fiber.Ctx) error {
	// Return all user engagement data (stub)
	return c.JSON(userEngagements)
}

// GetCourseAnalytics handles GET /api/analytics/course/:courseID
func (h *AnalyticsHandler) GetCourseAnalytics(c *fiber.Ctx) error {
	courseID := c.Params("courseID")
	for _, ca := range courseAnalytics {
		if ca.CourseID == courseID {
			return c.JSON(ca)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Course analytics not found"})
}

// GetEvents handles GET /api/analytics/events
func (h *AnalyticsHandler) GetEvents(c *fiber.Ctx) error {
	// Return all analytics events (stub)
	return c.JSON(analyticsEvents)
}
