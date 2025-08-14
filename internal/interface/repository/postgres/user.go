package postgres

import (
	"database/sql"
	"errors"
	"training-portal/internal/domain/user"
)

// UserRepository implements user data access using PostgreSQL.
type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByID(id string) (*user.User, error) {
	var u user.User
	err := r.DB.QueryRow(
		`SELECT id, name, email, password, role FROM users WHERE id = $1`,
		id,
	).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) FindByEmail(email string) (*user.User, error) {
	var u user.User
	err := r.DB.QueryRow(
		`SELECT id, name, email, password, role FROM users WHERE email = $1`,
		email,
	).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) Create(u *user.User) error {
	_, err := r.DB.Exec(
		`INSERT INTO users (id, name, email, password, role) VALUES ($1, $2, $3, $4, $5)`,
		u.ID, u.Name, u.Email, u.Password, u.Role,
	)
	return err
}

func (r *UserRepository) Update(u *user.User) error {
	res, err := r.DB.Exec(
		`UPDATE users SET name = $1, email = $2, password = $3, role = $4 WHERE id = $5`,
		u.Name, u.Email, u.Password, u.Role, u.ID,
	)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *UserRepository) Delete(id string) error {
	res, err := r.DB.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *UserRepository) List() ([]*user.User, error) {
	rows, err := r.DB.Query(`SELECT id, name, email, password, role FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*user.User
	for rows.Next() {
		var u user.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}
