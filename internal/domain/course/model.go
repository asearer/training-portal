// File: internal/domain/course/model.go
// Defines the course and module domain models

package course

type Course struct {
	ID          string // UUID
	Title       string
	Description string
	Category    string
	CreatedBy   string // user ID
	Published   bool
}

type Module struct {
	ID          string
	CourseID    string
	Title       string
	ContentType string // "video", "pdf", etc.
	ContentURL  string // URL to the file/resource
	OrderIndex  int    // for sequencing
}
