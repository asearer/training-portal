package handler

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Notification represents a user/system notification.
type Notification struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Message   string    `json:"message"`
	Read      bool      `json:"read"`
	CreatedAt time.Time `json:"created_at"`
}

// NotificationHandler provides HTTP handlers for notifications.
// Currently uses in-memory storage for demonstration.
type NotificationHandler struct {
	mu            sync.Mutex
	notifications map[string]*Notification // key = notification ID
}

// NewNotificationHandler creates a new NotificationHandler.
// In the future, inject a database repository/service for persistent storage.
func NewNotificationHandler() *NotificationHandler {
	return &NotificationHandler{
		notifications: make(map[string]*Notification),
	}
}

// ListNotifications handles GET /notifications
// Returns all notifications for the authenticated user.
func (h *NotificationHandler) ListNotifications(c *fiber.Ctx) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	userID := c.Locals("user_id") // Assuming JWT middleware sets this
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var userNotifications []Notification
	for _, n := range h.notifications {
		if n.UserID == userID.(string) {
			userNotifications = append(userNotifications, *n) // Replace with DB query in the future
		}
	}

	return c.JSON(userNotifications)
}

// MarkAsRead handles POST /notifications/:id/read
// Marks a notification as read.
func (h *NotificationHandler) MarkAsRead(c *fiber.Ctx) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing notification ID"})
	}

	notification, exists := h.notifications[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Notification not found"})
	}

	notification.Read = true
	// Replace with DB UPDATE query in the future

	return c.JSON(fiber.Map{
		"message":      "Notification marked as read",
		"notification": notification,
	})
}

// CreateNotification handles POST /notifications
// Creates a new notification (for admin/system use).
func (h *NotificationHandler) CreateNotification(c *fiber.Ctx) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	var input Notification
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Generate a simple unique ID using timestamp + user ID; replace with UUID or DB auto-ID
	input.ID = time.Now().Format("20060102150405") + "_" + input.UserID
	input.Read = false
	input.CreatedAt = time.Now()

	h.notifications[input.ID] = &input // Replace with DB INSERT in the future

	return c.Status(fiber.StatusCreated).JSON(input)
}
