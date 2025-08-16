package repository

import (
	"training-portal/internal/domain/course"
)

// CourseRepository defines the interface for course data access
type CourseRepository interface {
	Create(c *course.Course) error
	FindByID(id string) (*course.Course, error)
	Update(c *course.Course) error
	Delete(id string) error
	List() ([]*course.Course, error)
}
