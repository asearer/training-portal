package postgres

import (
	"database/sql"
	"errors"
	"training-portal/internal/domain/course"
)

type CourseRepository struct {
	DB *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{DB: db}
}

func (r *CourseRepository) FindByID(id string) (*course.Course, error) {
	var c course.Course
	err := r.DB.QueryRow(
		`SELECT id, title, description, category, created_by, is_published FROM courses WHERE id = $1`,
		id,
	).Scan(&c.ID, &c.Title, &c.Description, &c.Category, &c.CreatedBy, &c.Published)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (r *CourseRepository) Create(c *course.Course) error {
	_, err := r.DB.Exec(
		`INSERT INTO courses (id, title, description, category, created_by, is_published) VALUES ($1, $2, $3, $4, $5, $6)`,
		c.ID, c.Title, c.Description, c.Category, c.CreatedBy, c.Published,
	)
	return err
}

func (r *CourseRepository) Update(c *course.Course) error {
	res, err := r.DB.Exec(
		`UPDATE courses SET title = $1, description = $2, category = $3, created_by = $4, is_published = $5 WHERE id = $6`,
		c.Title, c.Description, c.Category, c.CreatedBy, c.Published, c.ID,
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

func (r *CourseRepository) Delete(id string) error {
	res, err := r.DB.Exec(`DELETE FROM courses WHERE id = $1`, id)
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

func (r *CourseRepository) List() ([]*course.Course, error) {
	rows, err := r.DB.Query(`SELECT id, title, description, category, created_by, is_published FROM courses`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*course.Course
	for rows.Next() {
		var c course.Course
		if err := rows.Scan(&c.ID, &c.Title, &c.Description, &c.Category, &c.CreatedBy, &c.Published); err != nil {
			return nil, err
		}
		courses = append(courses, &c)
	}
	return courses, nil
}
