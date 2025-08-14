package course

import (
	"errors"

	"training-portal/internal/domain/course"
	"training-portal/internal/interface/repository/postgres"

	"github.com/google/uuid"
)

type ModuleService struct {
	Repo *postgres.ModuleRepository
}

// ValidateTitle checks if the module title is valid (non-empty, reasonable length).
func ValidateTitle(title string) bool {
	return len(title) > 0 && len(title) <= 255
}

// CreateModule creates a new module with validation.
func (s *ModuleService) CreateModule(m *course.Module) error {
	if m == nil {
		return errors.New("module is required")
	}
	if !ValidateTitle(m.Title) {
		return errors.New("invalid module title")
	}
	if m.CourseID == "" {
		return errors.New("course_id is required")
	}
	m.ID = uuid.New().String()
	return s.Repo.Create(m)
}

// GetModule returns a module by ID.
func (s *ModuleService) GetModule(id string) (*course.Module, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	m, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, errors.New("module not found")
	}
	return m, nil
}

// UpdateModule updates module fields.
func (s *ModuleService) UpdateModule(m *course.Module) error {
	if m == nil || m.ID == "" {
		return errors.New("id is required")
	}
	if m.Title != "" && !ValidateTitle(m.Title) {
		return errors.New("invalid module title")
	}
	return s.Repo.Update(m)
}

// DeleteModule deletes a module by ID.
func (s *ModuleService) DeleteModule(id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return s.Repo.Delete(id)
}

// ListModulesByCourse returns all modules for a course.
func (s *ModuleService) ListModulesByCourse(courseID string) ([]*course.Module, error) {
	if courseID == "" {
		return nil, errors.New("course_id is required")
	}
	return s.Repo.ListByCourse(courseID)
}
