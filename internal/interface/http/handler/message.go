package handler

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Message represents a chat message structure.
// Replace with proper domain model and DB integration in the future.
type Message struct {
	ID         string    `json:"id"`
	SenderID   string    `json:"senderId"`
	ReceiverID string    `json:"receiverId"`
	Content    string    `json:"content"`
	Read       bool      `json:"read"`
	CreatedAt  time.Time `json:"createdAt"`
}

// MessageHandler provides HTTP handlers for messaging features.
type MessageHandler struct{}

// In-memory message storage (stub for demonstration purposes)
var messages = []Message{}

// SendMessage handles sending a direct message between users.
// POST /api/message
func (h *MessageHandler) SendMessage(c *fiber.Ctx) error {
	var req struct {
		SenderID   string `json:"senderId"`
		ReceiverID string `json:"receiverId"`
		Content    string `json:"content"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	newMessage := Message{
		ID:         strconv.Itoa(len(messages) + 1),
		SenderID:   req.SenderID,
		ReceiverID: req.ReceiverID,
		Content:    req.Content,
		Read:       false,
		CreatedAt:  time.Now(),
	}

	messages = append(messages, newMessage)

	return c.Status(fiber.StatusCreated).JSON(newMessage)
}

// GetMessages retrieves messages for a user or conversation.
// GET /api/messages?userId=<id> or GET /api/messages/:userId
func (h *MessageHandler) GetMessages(c *fiber.Ctx) error {
	userID := c.Query("userId")
	if userID == "" {
		userID = c.Params("userId")
	}
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User ID is required",
		})
	}

	var userMessages []Message
	for _, msg := range messages {
		if msg.SenderID == userID || msg.ReceiverID == userID {
			userMessages = append(userMessages, msg)
		}
	}

	return c.JSON(userMessages)
}

// MarkMessageRead marks a message as read.
// POST /api/message/:id/read
func (h *MessageHandler) MarkMessageRead(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Message ID is required",
		})
	}

	for i, msg := range messages {
		if msg.ID == id {
			messages[i].Read = true
			return c.JSON(fiber.Map{
				"message": "Message marked as read",
				"data":    messages[i],
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Message not found",
	})
}
