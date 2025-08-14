-- File: migrations/002_create_courses.sql
-- SQL migration to create courses table

CREATE TABLE courses (
    id UUID PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    category VARCHAR(100),
    created_by UUID REFERENCES users(id),
    is_published BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
