CREATE TABLE enrollments (
                             id UUID PRIMARY KEY,
                             user_id UUID REFERENCES users(id) ON DELETE CASCADE,
                             course_id UUID REFERENCES courses(id) ON DELETE CASCADE,
                             enrolled_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
