// File: internal/usecase/course/service.go
package course

import (
	"errors"

	"training-portal/internal/domain/course"
	"training-portal/internal/interface/repository"

	"github.com/google/uuid"
)

// CourseService provides business logic for courses.
type CourseService struct {
	// Use interface type for easier testing and future DB swaps
	Repo repository.CourseRepository
}

// ValidateTitle checks if the course title is non-empty and within a reasonable length.
func ValidateTitle(title string) bool {
	return len(title) > 0 && len(title) <= 255
}

// CreateCourse creates a new course after validating input.
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

	// Assign a new UUID for the course
	c.ID = uuid.New().String()

	// Persist course in repository
	// TODO: Replace with DB transaction if needed
	return s.Repo.Create(c)
}

// GetCourse retrieves a course by ID.
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

// UpdateCourse updates fields of an existing course.
func (s *CourseService) UpdateCourse(c *course.Course) error {
	if c == nil || c.ID == "" {
		return errors.New("course ID is required")
	}

	if c.Title != "" && !ValidateTitle(c.Title) {
		return errors.New("invalid course title")
	}

	// TODO: Add permission checks if needed
	return s.Repo.Update(c)
}

// DeleteCourse deletes a course by its ID.
func (s *CourseService) DeleteCourse(id string) error {
	if id == "" {
		return errors.New("id is required")
	}

	// TODO: Handle cascading deletes for modules/enrollments
	return s.Repo.Delete(id)
}

// ListCourses retrieves all courses.
func (s *CourseService) ListCourses() ([]*course.Course, error) {
	// TODO: Add pagination/filtering
	return s.Repo.List()
}
