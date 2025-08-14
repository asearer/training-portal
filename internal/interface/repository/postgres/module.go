package postgres

import (
	"database/sql"
	"errors"
	"training-portal/internal/domain/course"
)

type ModuleRepository struct {
	DB *sql.DB
}

func NewModuleRepository(db *sql.DB) *ModuleRepository {
	return &ModuleRepository{DB: db}
}

func (r *ModuleRepository) FindByID(id string) (*course.Module, error) {
	var m course.Module
	err := r.DB.QueryRow(
		`SELECT id, course_id, title, content_type, content_url, order_index FROM modules WHERE id = $1`,
		id,
	).Scan(&m.ID, &m.CourseID, &m.Title, &m.ContentType, &m.ContentURL, &m.OrderIndex)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (r *ModuleRepository) Create(m *course.Module) error {
	_, err := r.DB.Exec(
		`INSERT INTO modules (id, course_id, title, content_type, content_url, order_index) VALUES ($1, $2, $3, $4, $5, $6)`,
		m.ID, m.CourseID, m.Title, m.ContentType, m.ContentURL, m.OrderIndex,
	)
	return err
}

func (r *ModuleRepository) Update(m *course.Module) error {
	res, err := r.DB.Exec(
		`UPDATE modules SET course_id = $1, title = $2, content_type = $3, content_url = $4, order_index = $5 WHERE id = $6`,
		m.CourseID, m.Title, m.ContentType, m.ContentURL, m.OrderIndex, m.ID,
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

func (r *ModuleRepository) Delete(id string) error {
	res, err := r.DB.Exec(`DELETE FROM modules WHERE id = $1`, id)
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

func (r *ModuleRepository) ListByCourse(courseID string) ([]*course.Module, error) {
	rows, err := r.DB.Query(
		`SELECT id, course_id, title, content_type, content_url, order_index FROM modules WHERE course_id = $1 ORDER BY order_index ASC`,
		courseID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modules []*course.Module
	for rows.Next() {
		var m course.Module
		if err := rows.Scan(&m.ID, &m.CourseID, &m.Title, &m.ContentType, &m.ContentURL, &m.OrderIndex); err != nil {
			return nil, err
		}
		modules = append(modules, &m)
	}
	return modules, nil
}
