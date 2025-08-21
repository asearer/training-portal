CREATE TABLE quiz_submissions (
                                  id UUID PRIMARY KEY,
                                  quiz_id UUID REFERENCES quizzes(id) ON DELETE CASCADE,
                                  user_id UUID REFERENCES users(id) ON DELETE CASCADE,
                                  submitted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE quiz_answers (
                              submission_id UUID REFERENCES quiz_submissions(id) ON DELETE CASCADE,
                              question_id UUID REFERENCES questions(id) ON DELETE CASCADE,
                              response TEXT,
                              PRIMARY KEY(submission_id, question_id)
);
