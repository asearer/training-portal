package quiz

// Quiz represents a quiz attached to a course or module.
type Quiz struct {
	ID        string // UUID
	CourseID  string // Related course (optional)
	ModuleID  string // Related module (optional)
	Title     string
	Questions []Question
}

// Question represents a single quiz question.
type Question struct {
	ID          string   // UUID
	QuizID      string   // Parent quiz
	Text        string   // Question text
	Type        string   // "multiple_choice", "short_answer", etc.
	Choices     []string // For multiple choice
	Answer      string   // Correct answer (for auto-grading)
	Points      int      // Points for this question
	Explanation string   // Optional explanation/feedback
}
