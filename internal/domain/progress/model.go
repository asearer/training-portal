package progress

// Progress represents a user's progress in a course or module.
type Progress struct {
	UserID    string // UUID of the user
	CourseID  string // UUID of the course
	ModuleID  string // UUID of the module (optional, for module-level tracking)
	Completed bool   // Whether the module/course is completed
	Score     int    // Score or percentage (for quizzes/assessments)
	UpdatedAt int64  // Unix timestamp of last update
}
