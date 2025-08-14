-- File: migrations/003_create_modules.sql
-- SQL migration to create modules table

CREATE TABLE modules (
    id UUID PRIMARY KEY,
    course_id UUID REFERENCES courses(id),
    title VARCHAR(255),
    content_type VARCHAR(50),
    content_url TEXT,
    order_index INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
