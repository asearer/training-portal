package repository

import (
	"training-portal/internal/domain/user"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	Create(u *user.User) error
	FindByID(id string) (*user.User, error)
	FindByEmail(email string) (*user.User, error)
	Update(u *user.User) error
	Delete(id string) error
	List() ([]*user.User, error)
}
