CREATE TABLE quizzes (
                         id UUID PRIMARY KEY,
                         module_id UUID REFERENCES modules(id) ON DELETE CASCADE,
                         title VARCHAR(255) NOT NULL,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
