package handler

import "github.com/gofiber/fiber/v2"

// MessageHandler provides HTTP handlers for messaging features.
type MessageHandler struct{}

// SendMessage handles sending a direct message between users.
// POST /api/message
func (h *MessageHandler) SendMessage(c *fiber.Ctx) error {
	// TODO: Parse request, call usecase, return result
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"error": "SendMessage not implemented",
	})
}

// GetMessages handles retrieving messages for a user or conversation.
// GET /api/messages or /api/messages/:userId
func (h *MessageHandler) GetMessages(c *fiber.Ctx) error {
	// TODO: Parse params, call usecase, return messages
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"error": "GetMessages not implemented",
	})
}

// MarkMessageRead handles marking a message as read.
// POST /api/message/:id/read
func (h *MessageHandler) MarkMessageRead(c *fiber.Ctx) error {
	// TODO: Parse message ID, call usecase, return result
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"error": "MarkMessageRead not implemented",
	})
}
