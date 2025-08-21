// File: internal/repository/postgres/course_repository.go
package postgres

import (
	"database/sql"
	"errors"
	"training-portal/internal/domain/course"
)

// CourseRepository provides methods to interact with courses table in Postgres.
type CourseRepository struct {
	db *sql.DB
}

// NewCourseRepository creates a new CourseRepository instance.
func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

// Create inserts a new course into the database.
func (r *CourseRepository) Create(c *course.Course) error {
	query := `INSERT INTO courses (id, title, description, category, published) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, c.ID, c.Title, c.Description, c.Category, c.Published)
	return err
}

// GetByID retrieves a course by its ID.
func (r *CourseRepository) GetByID(id string) (*course.Course, error) {
	c := &course.Course{}
	query := `SELECT id, title, description, category, published FROM courses WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&c.ID, &c.Title, &c.Description, &c.Category, &c.Published)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // course not found
		}
		return nil, err
	}
	return c, nil
}

// Update modifies an existing course.
func (r *CourseRepository) Update(c *course.Course) error {
	query := `UPDATE courses SET title = $1, description = $2, category = $3, published = $4 WHERE id = $5`
	result, err := r.db.Exec(query, c.Title, c.Description, c.Category, c.Published, c.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no rows updated")
	}
	return nil
}

// Delete removes a course by ID.
func (r *CourseRepository) Delete(id string) error {
	query := `DELETE FROM courses WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("course not found")
	}
	return nil
}

// List returns all courses.
func (r *CourseRepository) List() ([]*course.Course, error) {
	query := `SELECT id, title, description, category, published FROM courses`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*course.Course
	for rows.Next() {
		c := &course.Course{}
		if err := rows.Scan(&c.ID, &c.Title, &c.Description, &c.Category, &c.Published); err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}
	return courses, nil
}
