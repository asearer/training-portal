// File: internal/usecase/course/module_service.go
package course

import (
	"errors"

	"training-portal/internal/domain/course"
	"training-portal/internal/interface/repository"

	"github.com/google/uuid"
)

// ModuleService provides business logic for course modules.
type ModuleService struct {
	// Use interface type for easier testing and future DB swaps
	Repo repository.ModuleRepository
}

// ValidateTitle checks if the module title is non-empty and within a reasonable length.
func ValidateTitle(title string) bool {
	return len(title) > 0 && len(title) <= 255
}

// CreateModule creates a new module after validation.
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

	// Assign a new UUID for the module
	m.ID = uuid.New().String()

	// Persist module in repository
	// Implement DB transaction if needed in future for consistency
	if err := s.Repo.Create(m); err != nil {
		return err
	}

	return nil
}

// GetModule retrieves a module by ID.
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

// UpdateModule updates fields of an existing module.
func (s *ModuleService) UpdateModule(m *course.Module) error {
	if m == nil || m.ID == "" {
		return errors.New("module ID is required")
	}
	if m.Title != "" && !ValidateTitle(m.Title) {
		return errors.New("invalid module title")
	}

	// Placeholder for permission checks or other business rules
	// Example: check if current user is course creator before updating
	// TODO: Implement permission checks based on user context

	if err := s.Repo.Update(m); err != nil {
		return err
	}

	return nil
}

// DeleteModule deletes a module by its ID.
func (s *ModuleService) DeleteModule(id string) error {
	if id == "" {
		return errors.New("id is required")
	}

	// Handle cascading deletes if module has associated content (e.g., lessons, quizzes)
	// TODO: Implement cascading deletes in future DB integration

	if err := s.Repo.Delete(id); err != nil {
		return err
	}

	return nil
}

// ListModulesByCourse returns all modules for a specific course.
func (s *ModuleService) ListModulesByCourse(courseID string) ([]*course.Module, error) {
	if courseID == "" {
		return nil, errors.New("course_id is required")
	}

	// Future enhancement: add pagination/filtering
	modules, err := s.Repo.ListByCourse(courseID)
	if err != nil {
		return nil, err
	}

	return modules, nil
}
