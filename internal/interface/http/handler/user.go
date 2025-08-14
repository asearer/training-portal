package handler

import (
	"os"
	"time"

	"training-portal/internal/domain/user"
	userusecase "training-portal/internal/usecase/user"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
	Service *userusecase.UserService
}

var _ = UserHandler{} // Exported for router.go

// Register handles POST /register
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	u, err := h.Service.Register(req.Name, req.Email, req.Password, user.RoleEmployee)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	u.Password = "" // never return password
	return c.Status(fiber.StatusCreated).JSON(u)
}

// Login handles POST /login
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	u, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}
	claims := jwt.MapClaims{
		"user_id": u.ID,
		"role":    u.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to sign token"})
	}
	return c.JSON(fiber.Map{"token": tokenStr})
}

// GetUser handles GET /user/:id
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	u, err := h.Service.GetUser(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	u.Password = ""
	return c.JSON(u)
}

// ListUsers handles GET /users
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.Service.ListUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	for _, u := range users {
		u.Password = ""
	}
	return c.JSON(users)
}

// UpdateUser handles PUT /user/:id
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var req struct {
		Name  string    `json:"name"`
		Email string    `json:"email"`
		Role  user.Role `json:"role"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	u := &user.User{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}
	if err := h.Service.UpdateUser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "User updated"})
}

// UpdatePassword handles PUT /user/:id/password
func (h *UserHandler) UpdatePassword(c *fiber.Ctx) error {
	id := c.Params("id")
	var req struct {
		NewPassword string `json:"new_password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.Service.UpdatePassword(id, req.NewPassword); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Password updated"})
}

// DeleteUser handles DELETE /user/:id
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.Service.DeleteUser(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "User deleted"})
}
