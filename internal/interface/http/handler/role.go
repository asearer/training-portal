package handler

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

// Role represents a system role with a name and optional description.
type Role struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// RoleHandler provides HTTP handlers for role and permission management.
type RoleHandler struct {
	mu        sync.Mutex
	roles     map[string]Role            // key = role name
	userRoles map[string]map[string]bool // key = userID, value = set of role names
}

// NewRoleHandler creates a new RoleHandler with in-memory storage.
// TODO: Replace in-memory storage with persistent DB integration.
func NewRoleHandler() *RoleHandler {
	return &RoleHandler{
		roles:     make(map[string]Role),
		userRoles: make(map[string]map[string]bool),
	}
}

// AssignRole assigns a role to a user.
func (h *RoleHandler) AssignRole(c *fiber.Ctx) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	var input struct {
		UserID string `json:"user_id"`
		Role   string `json:"role"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if _, exists := h.roles[input.Role]; !exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Role does not exist"})
	}

	if _, exists := h.userRoles[input.UserID]; !exists {
		h.userRoles[input.UserID] = make(map[string]bool)
	}
	h.userRoles[input.UserID][input.Role] = true // TODO: Replace with DB insert

	return c.JSON(fiber.Map{"message": "Role assigned successfully"})
}

// RevokeRole revokes a role from a user.
func (h *RoleHandler) RevokeRole(c *fiber.Ctx) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	var input struct {
		UserID string `json:"user_id"`
		Role   string `json:"role"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if roles, exists := h.userRoles[input.UserID]; exists {
		delete(roles, input.Role) // TODO: Replace with DB delete
	}

	return c.JSON(fiber.Map{"message": "Role revoked successfully"})
}

// ListRoles lists all roles in the system.
func (h *RoleHandler) ListRoles(c *fiber.Ctx) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	roleList := make([]Role, 0, len(h.roles))
	for _, r := range h.roles { // TODO: Replace with DB query
		roleList = append(roleList, r)
	}

	return c.JSON(roleList)
}

// ListUserRoles lists all roles assigned to a specific user.
func (h *RoleHandler) ListUserRoles(c *fiber.Ctx) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	userID := c.Params("user_id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing user_id"})
	}

	roles := []string{}
	if assigned, exists := h.userRoles[userID]; exists {
		for role := range assigned {
			roles = append(roles, role)
		}
	}

	return c.JSON(fiber.Map{"user_id": userID, "roles": roles})
}
