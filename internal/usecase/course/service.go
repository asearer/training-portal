package course

import (
	"errors"

	"training-portal/internal/domain/course"
	"training-portal/internal/interface/repository"

	"github.com/google/uuid"
)

type CourseService struct {
	Repo repository.CourseRepository
}

// ValidateTitle checks if the course title is valid (non-empty, reasonable length).
func ValidateTitle(title string) bool {
	return len(title) > 0 && len(title) <= 255
}

// CreateCourse creates a new course with validation.
func (s *CourseService) CreateCourse(c *course.Course) error {
	if c == nil {
		return errors.New("course is required")
	}
	if !ValidateTitle(c.Title) {
		return errors.New("invalid course title")
	}
	if c.CreatedBy == "" {
		return errors.New("created_by is required")
	}
	c.ID = uuid.New().String()
	return s.Repo.Create(c)
}

// GetCourse returns a course by ID.
func (s *CourseService) GetCourse(id string) (*course.Course, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	c, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, errors.New("course not found")
	}
	return c, nil
}

// UpdateCourse updates course fields.
func (s *CourseService) UpdateCourse(c *course.Course) error {
	if c == nil || c.ID == "" {
		return errors.New("id is required")
	}
	if c.Title != "" && !ValidateTitle(c.Title) {
		return errors.New("invalid course title")
	}
	return s.Repo.Update(c)
}

// DeleteCourse deletes a course by ID.
func (s *CourseService) DeleteCourse(id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return s.Repo.Delete(id)
}

// ListCourses returns all courses.
func (s *CourseService) ListCourses() ([]*course.Course, error) {
	return s.Repo.List()
}
