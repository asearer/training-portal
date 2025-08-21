CREATE TABLE questions (
                           id UUID PRIMARY KEY,
                           quiz_id UUID REFERENCES quizzes(id) ON DELETE CASCADE,
                           question_text TEXT NOT NULL,
                           answer TEXT NOT NULL,
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
