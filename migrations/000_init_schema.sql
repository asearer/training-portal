-- ============================================
-- Training Portal Database Schema
-- ============================================

-- Enable UUID generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ============================================
-- Users Table
-- ============================================
CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       username VARCHAR(100) UNIQUE NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password_hash VARCHAR(255) NOT NULL,
                       role VARCHAR(50) DEFAULT 'student',
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- Courses Table
-- ============================================
CREATE TABLE courses (
                         id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                         title VARCHAR(255) NOT NULL,
                         description TEXT,
                         created_by UUID REFERENCES users(id) ON DELETE SET NULL,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- Modules Table
-- ============================================
CREATE TABLE modules (
                         id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                         course_id UUID REFERENCES courses(id) ON DELETE CASCADE,
                         title VARCHAR(255) NOT NULL,
                         content TEXT,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- Quizzes Table
-- ============================================
CREATE TABLE quizzes (
                         id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                         module_id UUID REFERENCES modules(id) ON DELETE CASCADE,
                         title VARCHAR(255) NOT NULL,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- Questions Table
-- ============================================
CREATE TABLE questions (
                           id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                           quiz_id UUID REFERENCES quizzes(id) ON DELETE CASCADE,
                           question_text TEXT NOT NULL,
                           answer TEXT NOT NULL,
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- Quiz Submissions Table
-- ============================================
CREATE TABLE quiz_submissions (
                                  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
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

-- ============================================
-- Enrollments Table
-- ============================================
CREATE TABLE enrollments (
                             id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                             user_id UUID REFERENCES users(id) ON DELETE CASCADE,
                             course_id UUID REFERENCES courses(id) ON DELETE CASCADE,
                             enrolled_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- Progress Table
-- ============================================
CREATE TABLE progress (
                          id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                          user_id UUID REFERENCES users(id) ON DELETE CASCADE,
                          course_id UUID REFERENCES courses(id) ON DELETE CASCADE,
                          completed_modules UUID[],
                          completed_quizzes JSONB,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- Certificates Table
-- ============================================
CREATE TABLE certificates (
                              id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                              user_id UUID REFERENCES users(id) ON DELETE CASCADE,
                              course_id UUID REFERENCES courses(id) ON DELETE CASCADE,
                              issued_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              certificate_url TEXT
);

-- ============================================
-- Notifications Table
-- ============================================
CREATE TABLE notifications (
                               id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                               user_id UUID REFERENCES users(id) ON DELETE CASCADE,
                               message TEXT NOT NULL,
                               read BOOLEAN DEFAULT FALSE,
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- Messages Table
-- ============================================
CREATE TABLE messages (
                          id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                          sender_id UUID REFERENCES users(id) ON DELETE CASCADE,
                          receiver_id UUID REFERENCES users(id) ON DELETE CASCADE,
                          content TEXT NOT NULL,
                          read BOOLEAN DEFAULT FALSE,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- Forums Table
-- ============================================
CREATE TABLE forums (
                        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                        title VARCHAR(255) NOT NULL,
                        created_by UUID REFERENCES users(id) ON DELETE SET NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- Posts Table
-- ============================================
CREATE TABLE posts (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       forum_id UUID REFERENCES forums(id) ON DELETE CASCADE,
                       user_id UUID REFERENCES users(id) ON DELETE SET NULL,
                       content TEXT NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- Replies Table
-- ============================================
CREATE TABLE replies (
                         id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                         post_id UUID REFERENCES posts(id) ON DELETE CASCADE,
                         user_id UUID REFERENCES users(id) ON DELETE SET NULL,
                         content TEXT NOT NULL,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- Indexes for Performance
-- ============================================
CREATE INDEX idx_enrollments_user ON enrollments(user_id);
CREATE INDEX idx_enrollments_course ON enrollments(course_id);

CREATE INDEX idx_progress_user ON progress(user_id);
CREATE INDEX idx_progress_course ON progress(course_id);

CREATE INDEX idx_modules_course ON modules(course_id);
CREATE INDEX idx_quizzes_module ON quizzes(module_id);
CREATE INDEX idx_questions_quiz ON questions(quiz_id);

CREATE INDEX idx_quiz_submissions_quiz ON quiz_submissions(quiz_id);
CREATE INDEX idx_quiz_submissions_user ON quiz_submissions(user_id);

CREATE INDEX idx_messages_sender ON messages(sender_id);
CREATE INDEX idx_messages_receiver ON messages(receiver_id);

CREATE INDEX idx_posts_forum ON posts(forum_id);
CREATE INDEX idx_replies_post ON replies(post_id);
