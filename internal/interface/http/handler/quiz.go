package handler

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Quiz represents a basic quiz structure.
// In a real implementation, this would correspond to a database model.
type Quiz struct {
	ID          string              `json:"id"`          // unique quiz ID
	Title       string              `json:"title"`       // quiz title
	Questions   []Question          `json:"questions"`   // list of questions
	Submissions map[string][]Answer `json:"submissions"` // map[userID]answers
}

// Question represents a single quiz question.
type Question struct {
	ID     string `json:"id"`     // question ID
	Text   string `json:"text"`   // question text
	Answer string `json:"answer"` // correct answer
}

// Answer represents a user-submitted answer.
type Answer struct {
	QuestionID string `json:"question_id"` // ID of the question
	Response   string `json:"response"`    // user's answer
}

// QuizHandler provides HTTP handlers for quizzes.
// Currently uses in-memory storage; replace with database in production.
type QuizHandler struct {
	mu      sync.Mutex
	quizzes map[string]*Quiz // in-memory storage
}

// NewQuizHandler creates a new QuizHandler.
// In the future, inject a database repository/service here.
func NewQuizHandler() *QuizHandler {
	return &QuizHandler{
		quizzes: make(map[string]*Quiz),
	}
}

// CreateQuiz handles POST /api/quiz
// Stores the quiz in memory; replace with DB insert in the future.
func (h *QuizHandler) CreateQuiz(c *fiber.Ctx) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	var quiz Quiz
	if err := c.BodyParser(&quiz); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	quiz.ID = uuid.New().String()
	if quiz.Submissions == nil {
		quiz.Submissions = make(map[string][]Answer)
	}

	h.quizzes[quiz.ID] = &quiz // Replace with DB insert
	return c.Status(fiber.StatusCreated).JSON(quiz)
}

// GetQuiz handles GET /api/quiz/:id
// Fetches quiz from in-memory storage; replace with DB query in the future.
func (h *QuizHandler) GetQuiz(c *fiber.Ctx) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	id := c.Params("id")
	quiz, exists := h.quizzes[id] // Replace with DB query
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Quiz not found"})
	}
	return c.JSON(quiz)
}

// ListQuizzes handles GET /api/quizzes
// Returns all quizzes in memory; replace with DB query in the future.
func (h *QuizHandler) ListQuizzes(c *fiber.Ctx) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	list := make([]*Quiz, 0, len(h.quizzes))
	for _, q := range h.quizzes { // Replace with DB query
		list = append(list, q)
	}
	return c.JSON(list)
}

// SubmitQuiz handles POST /api/quiz/:id/submit
// Stores user submission in memory; replace with DB insert/update in the future.
func (h *QuizHandler) SubmitQuiz(c *fiber.Ctx) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	id := c.Params("id")
	quiz, exists := h.quizzes[id] // Replace with DB query
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Quiz not found"})
	}

	var submission struct {
		UserID  string   `json:"user_id"`
		Answers []Answer `json:"answers"`
	}
	if err := c.BodyParser(&submission); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid submission"})
	}

	quiz.Submissions[submission.UserID] = submission.Answers // Replace with DB insert/update
	return c.JSON(fiber.Map{"message": "Quiz submitted successfully"})
}

// GradeQuiz handles GET /api/quiz/:id/grade?user_id=xxx
// Computes score from in-memory submissions; replace with DB query in the future.
func (h *QuizHandler) GradeQuiz(c *fiber.Ctx) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	id := c.Params("id")
	userID := c.Query("user_id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing user_id query parameter"})
	}

	quiz, exists := h.quizzes[id] // Replace with DB query
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Quiz not found"})
	}

	answers, submitted := quiz.Submissions[userID] // Replace with DB query
	if !submitted {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No submission found for user"})
	}

	// Calculate score
	score := 0
	for _, ans := range answers {
		for _, q := range quiz.Questions {
			if q.ID == ans.QuestionID && q.Answer == ans.Response {
				score++
			}
		}
	}

	return c.JSON(fiber.Map{
		"user_id": userID,
		"quiz_id": id,
		"score":   score,
		"total":   len(quiz.Questions),
	})
}
