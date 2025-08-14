package user

import (
	"errors"
	"regexp"

	"training-portal/internal/domain/user"
	"training-portal/internal/interface/repository/postgres"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *postgres.UserRepository
}

// ValidateEmail checks if the email is in a valid format.
func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// Register creates a new user with hashed password and basic validation.
func (s *UserService) Register(name, email, password string, role user.Role) (*user.User, error) {
	if name == "" || email == "" || password == "" {
		return nil, errors.New("name, email, and password are required")
	}
	if !ValidateEmail(email) {
		return nil, errors.New("invalid email format")
	}
	existing, err := s.Repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("email already registered")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u := &user.User{
		ID:       uuid.New().String(),
		Name:     name,
		Email:    email,
		Password: string(hashed),
		Role:     role,
	}
	if err := s.Repo.Create(u); err != nil {
		return nil, err
	}
	return u, nil
}

// Login authenticates a user by email and password.
func (s *UserService) Login(email, password string) (*user.User, error) {
	u, err := s.Repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("invalid credentials")
	}
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil {
		return nil, errors.New("invalid credentials")
	}
	return u, nil
}

// GetUser returns a user by ID.
func (s *UserService) GetUser(id string) (*user.User, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	u, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("user not found")
	}
	return u, nil
}

// UpdateUser updates user fields except password.
func (s *UserService) UpdateUser(u *user.User) error {
	if u.ID == "" {
		return errors.New("id is required")
	}
	if u.Email != "" && !ValidateEmail(u.Email) {
		return errors.New("invalid email format")
	}
	return s.Repo.Update(u)
}

// UpdatePassword updates a user's password.
func (s *UserService) UpdatePassword(id, newPassword string) error {
	if id == "" || newPassword == "" {
		return errors.New("id and new password are required")
	}
	u, err := s.Repo.FindByID(id)
	if err != nil {
		return err
	}
	if u == nil {
		return errors.New("user not found")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return s.Repo.Update(u)
}

// DeleteUser deletes a user by ID.
func (s *UserService) DeleteUser(id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return s.Repo.Delete(id)
}

// ListUsers returns all users.
func (s *UserService) ListUsers() ([]*user.User, error) {
	return s.Repo.List()
}
