package enrollment

// Enrollment represents a user's enrollment in a course.
type Enrollment struct {
	ID        string // UUID
	UserID    string // UUID of the enrolled user
	CourseID  string // UUID of the course
	Status    string // e.g., "active", "completed", "dropped"
	CreatedAt int64  // Unix timestamp
	UpdatedAt int64  // Unix timestamp
}
