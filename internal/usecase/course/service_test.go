package course

import (
	"errors"
	"testing"

	"training-portal/internal/domain/course"
	"training-portal/internal/interface/repository"
)

// MockCourseRepository is a mock implementation of the course repository
type MockCourseRepository struct {
	courses   map[string]*course.Course
	nextID    int
	shouldFail bool
}

func NewMockCourseRepository() repository.CourseRepository {
	return &MockCourseRepository{
		courses: make(map[string]*course.Course),
		nextID:  1,
	}
}

func (m *MockCourseRepository) Create(c *course.Course) error {
	if m.shouldFail {
		return errors.New("database error")
	}
	if c.ID == "" {
		c.ID = string(rune(m.nextID))
		m.nextID++
	}
	m.courses[c.ID] = c
	return nil
}

func (m *MockCourseRepository) FindByID(id string) (*course.Course, error) {
	if m.shouldFail {
		return nil, errors.New("database error")
	}
	if c, exists := m.courses[id]; exists {
		return c, nil
	}
	return nil, nil
}

func (m *MockCourseRepository) Update(c *course.Course) error {
	if m.shouldFail {
		return errors.New("database error")
	}
	if _, exists := m.courses[c.ID]; !exists {
		return errors.New("course not found")
	}
	m.courses[c.ID] = c
	return nil
}

func (m *MockCourseRepository) Delete(id string) error {
	if m.shouldFail {
		return errors.New("database error")
	}
	if _, exists := m.courses[id]; !exists {
		return errors.New("course not found")
	}
	delete(m.courses, id)
	return nil
}

func (m *MockCourseRepository) List() ([]*course.Course, error) {
	if m.shouldFail {
		return nil, errors.New("database error")
	}
	courses := make([]*course.Course, 0, len(m.courses))
	for _, c := range m.courses {
		courses = append(courses, c)
	}
	return courses, nil
}

func TestCourseService_CreateCourse(t *testing.T) {
	tests := []struct {
		name        string
		course      *course.Course
		shouldFail  bool
		expectError bool
	}{
		{
			name: "Valid course creation",
			course: &course.Course{
				Title:       "Introduction to Go",
				Description: "Learn the basics of Go programming",
				Category:    "Programming",
				CreatedBy:   "user123",
				Published:   false,
			},
			shouldFail:  false,
			expectError: false,
		},
		{
			name: "Nil course",
			course: nil,
			shouldFail:  false,
			expectError: true,
		},
		{
			name: "Empty title",
			course: &course.Course{
				Title:       "",
				Description: "Learn the basics of Go programming",
				Category:    "Programming",
				CreatedBy:   "user123",
				Published:   false,
			},
			shouldFail:  false,
			expectError: true,
		},
		{
			name: "Empty created_by",
			course: &course.Course{
				Title:       "Introduction to Go",
				Description: "Learn the basics of Go programming",
				Category:    "Programming",
				CreatedBy:   "",
				Published:   false,
			},
			shouldFail:  false,
			expectError: true,
		},
		{
			name: "Database error",
			course: &course.Course{
				Title:       "Introduction to Go",
				Description: "Learn the basics of Go programming",
				Category:    "Programming",
				CreatedBy:   "user123",
				Published:   false,
			},
			shouldFail:  true,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockCourseRepository().(*MockCourseRepository)
			
			if tt.shouldFail {
				mockRepo.shouldFail = true
			}

			service := &CourseService{Repo: mockRepo}
			err := service.CreateCourse(tt.course)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if tt.course.ID == "" {
					t.Errorf("Expected course ID to be generated but got empty")
				}
			}
		})
	}
}

func TestCourseService_GetCourse(t *testing.T) {
	tests := []struct {
		name        string
		courseID    string
		setupCourse bool
		shouldFail  bool
		expectError bool
	}{
		{
			name:        "Valid course ID",
			courseID:    "1",
			setupCourse: true,
			shouldFail:  false,
			expectError: false,
		},
		{
			name:        "Empty course ID",
			courseID:    "",
			setupCourse: false,
			shouldFail:  false,
			expectError: true,
		},
		{
			name:        "Course not found",
			courseID:    "999",
			setupCourse: false,
			shouldFail:  false,
			expectError: true,
		},
		{
			name:        "Database error",
			courseID:    "1",
			setupCourse: true,
			shouldFail:  true,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockCourseRepository().(*MockCourseRepository)
			
			if tt.setupCourse {
				testCourse := &course.Course{
					ID:          "1",
					Title:       "Introduction to Go",
					Description: "Learn the basics of Go programming",
					Category:    "Programming",
					CreatedBy:   "user123",
					Published:   false,
				}
				mockRepo.Create(testCourse)
			}
			
			if tt.shouldFail {
				mockRepo.shouldFail = true
			}

			service := &CourseService{Repo: mockRepo}
			course, err := service.GetCourse(tt.courseID)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if course == nil {
					t.Errorf("Expected course but got nil")
				}
				if course.ID != tt.courseID {
					t.Errorf("Expected ID %s, got %s", tt.courseID, course.ID)
				}
			}
		})
	}
}

func TestCourseService_UpdateCourse(t *testing.T) {
	tests := []struct {
		name        string
		course      *course.Course
		setupCourse bool
		shouldFail  bool
		expectError bool
	}{
		{
			name: "Valid update",
			course: &course.Course{
				ID:          "1",
				Title:       "Introduction to Go Updated",
				Description: "Learn the basics of Go programming - Updated",
				Category:    "Programming",
				CreatedBy:   "user123",
				Published:   true,
			},
			setupCourse: true,
			shouldFail:  false,
			expectError: false,
		},
		{
			name: "Nil course",
			course: nil,
			setupCourse: false,
			shouldFail:  false,
			expectError: true,
		},
		{
			name: "Empty ID",
			course: &course.Course{
				ID:          "",
				Title:       "Introduction to Go Updated",
				Description: "Learn the basics of Go programming - Updated",
				Category:    "Programming",
				CreatedBy:   "user123",
				Published:   true,
			},
			setupCourse: false,
			shouldFail:  false,
			expectError: true,
		},
		{
			name: "Empty title",
			course: &course.Course{
				ID:          "1",
				Title:       "",
				Description: "Learn the basics of Go programming - Updated",
				Category:    "Programming",
				CreatedBy:   "user123",
				Published:   true,
			},
			setupCourse: true,
			shouldFail:  false,
			expectError: true,
		},
		{
			name: "Database error",
			course: &course.Course{
				ID:          "1",
				Title:       "Introduction to Go Updated",
				Description: "Learn the basics of Go programming - Updated",
				Category:    "Programming",
				CreatedBy:   "user123",
				Published:   true,
			},
			setupCourse: true,
			shouldFail:  true,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockCourseRepository().(*MockCourseRepository)
			
			if tt.setupCourse {
				testCourse := &course.Course{
					ID:          "1",
					Title:       "Introduction to Go",
					Description: "Learn the basics of Go programming",
					Category:    "Programming",
					CreatedBy:   "user123",
					Published:   false,
				}
				mockRepo.Create(testCourse)
			}
			
			if tt.shouldFail {
				mockRepo.shouldFail = true
			}

			service := &CourseService{Repo: mockRepo}
			err := service.UpdateCourse(tt.course)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestCourseService_DeleteCourse(t *testing.T) {
	tests := []struct {
		name        string
		courseID    string
		setupCourse bool
		shouldFail  bool
		expectError bool
	}{
		{
			name:        "Valid delete",
			courseID:    "1",
			setupCourse: true,
			shouldFail:  false,
			expectError: false,
		},
		{
			name:        "Empty course ID",
			courseID:    "",
			setupCourse: false,
			shouldFail:  false,
			expectError: true,
		},
		{
			name:        "Course not found",
			courseID:    "999",
			setupCourse: false,
			shouldFail:  false,
			expectError: true,
		},
		{
			name:        "Database error",
			courseID:    "1",
			setupCourse: true,
			shouldFail:  true,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockCourseRepository().(*MockCourseRepository)
			
			if tt.setupCourse {
				testCourse := &course.Course{
					ID:          "1",
					Title:       "Introduction to Go",
					Description: "Learn the basics of Go programming",
					Category:    "Programming",
					CreatedBy:   "user123",
					Published:   false,
				}
				mockRepo.Create(testCourse)
			}
			
			if tt.shouldFail {
				mockRepo.shouldFail = true
			}

			service := &CourseService{Repo: mockRepo}
			err := service.DeleteCourse(tt.courseID)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestCourseService_ListCourses(t *testing.T) {
	tests := []struct {
		name         string
		setupCourses bool
		shouldFail   bool
		expectError  bool
		expectedCount int
	}{
		{
			name:         "List courses successfully",
			setupCourses: true,
			shouldFail:   false,
			expectError:  false,
			expectedCount: 2,
		},
		{
			name:         "Empty course list",
			setupCourses: false,
			shouldFail:   false,
			expectError:  false,
			expectedCount: 0,
		},
		{
			name:         "Database error",
			setupCourses: true,
			shouldFail:   true,
			expectError:  true,
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockCourseRepository().(*MockCourseRepository)
			
			if tt.setupCourses {
				course1 := &course.Course{
					ID:          "1",
					Title:       "Introduction to Go",
					Description: "Learn the basics of Go programming",
					Category:    "Programming",
					CreatedBy:   "user123",
					Published:   false,
				}
				course2 := &course.Course{
					ID:          "2",
					Title:       "Advanced Go",
					Description: "Advanced Go programming concepts",
					Category:    "Programming",
					CreatedBy:   "user456",
					Published:   true,
				}
				mockRepo.Create(course1)
				mockRepo.Create(course2)
			}
			
			if tt.shouldFail {
				mockRepo.shouldFail = true
			}

			service := &CourseService{Repo: mockRepo}
			courses, err := service.ListCourses()

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if len(courses) != tt.expectedCount {
					t.Errorf("Expected %d courses, got %d", tt.expectedCount, len(courses))
				}
			}
		})
	}
}

func TestValidateTitle(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		expected bool
	}{
		{"Valid title", "Introduction to Go", true},
		{"Empty title", "", false},
		{"Very long title", string(make([]byte, 256)), false},
		{"Single character", "A", true},
		{"255 characters", string(make([]byte, 255)), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateTitle(tt.title); got != tt.expected {
				t.Errorf("ValidateTitle(%q) = %v, want %v", tt.title, got, tt.expected)
			}
		})
	}
}
