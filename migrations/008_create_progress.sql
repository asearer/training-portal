CREATE TABLE progress (
                          id UUID PRIMARY KEY,
                          user_id UUID REFERENCES users(id) ON DELETE CASCADE,
                          course_id UUID REFERENCES courses(id) ON DELETE CASCADE,
                          completed_modules UUID[],
                          completed_quizzes JSONB,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
