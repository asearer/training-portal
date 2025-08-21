// File: internal/repository/postgres/user_repository.go
package postgres

import (
	"database/sql"
	"errors"
	"training-portal/internal/domain/user"
)

// UserRepository provides methods to interact with users table in Postgres.
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository instance.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create inserts a new user into the database.
func (r *UserRepository) Create(u *user.User) error {
	query := `INSERT INTO users (id, name, email, password, role) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, u.ID, u.Name, u.Email, u.Password, u.Role)
	return err
}

// GetByID retrieves a user by their ID.
func (r *UserRepository) GetByID(id string) (*user.User, error) {
	u := &user.User{}
	query := `SELECT id, name, email, password, role FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // user not found
		}
		return nil, err
	}
	return u, nil
}

// Update modifies an existing user.
func (r *UserRepository) Update(u *user.User) error {
	query := `UPDATE users SET name = $1, email = $2, password = $3, role = $4 WHERE id = $5`
	result, err := r.db.Exec(query, u.Name, u.Email, u.Password, u.Role, u.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no rows updated")
	}
	return nil
}

// Delete removes a user by ID.
func (r *UserRepository) Delete(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

// List returns all users.
func (r *UserRepository) List() ([]*user.User, error) {
	query := `SELECT id, name, email, password, role FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*user.User
	for rows.Next() {
		u := &user.User{}
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
