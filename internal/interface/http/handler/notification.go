package handler

import (
	"github.com/gofiber/fiber/v2"
)

// NotificationHandler provides HTTP handlers for notification-related endpoints.
type NotificationHandler struct{}

// ListNotifications handles GET /notifications for the current user.
// TODO: Implement fetching notifications for the authenticated user.
func (h *NotificationHandler) ListNotifications(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "List notifications endpoint (stub)"})
}

// MarkAsRead handles POST /notifications/:id/read to mark a notification as read.
// TODO: Implement marking a notification as read.
func (h *NotificationHandler) MarkAsRead(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Mark notification as read endpoint (stub)"})
}

// CreateNotification handles POST /notifications (admin/system use).
// TODO: Implement creating/sending a notification.
func (h *NotificationHandler) CreateNotification(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Create notification endpoint (stub)"})
}
