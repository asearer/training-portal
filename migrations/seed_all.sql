-- =====================================================
-- Training Portal: Comprehensive Seed Data
-- =====================================================

-- USERS
INSERT INTO users (id, username, email, password_hash, role)
VALUES
    ('11111111-1111-1111-1111-111111111111', 'admin', 'admin@example.com', 'hashedpassword', 'admin'),
    ('22222222-2222-2222-2222-222222222222', 'teacher1', 'teacher1@example.com', 'hashedpassword', 'teacher'),
    ('33333333-3333-3333-3333-333333333333', 'student1', 'student1@example.com', 'hashedpassword', 'student'),
    ('44444444-4444-4444-4444-444444444444', 'student2', 'student2@example.com', 'hashedpassword', 'student');

-- COURSES
INSERT INTO courses (id, title, description, created_by)
VALUES
    ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Intro to Go', 'Learn Go programming basics', '22222222-2222-2222-2222-222222222222'),
    ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'Advanced Go', 'Deep dive into Go', '22222222-2222-2222-2222-222222222222');

-- MODULES
INSERT INTO modules (id, course_id, title, content)
VALUES
    ('11111111-aaaa-1111-aaaa-111111111111', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Go Basics', 'Variables, Loops, Functions'),
    ('22222222-bbbb-2222-bbbb-222222222222', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'Go Concurrency', 'Goroutines and Channels');

-- QUIZZES
INSERT INTO quizzes (id, course_id, title)
VALUES
    ('q1111111-aaaa-1111-aaaa-111111111111', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Go Basics Quiz'),
    ('q2222222-bbbb-2222-bbbb-222222222222', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'Go Concurrency Quiz');

-- QUESTIONS
INSERT INTO questions (id, quiz_id, text, answer)
VALUES
    ('qq111111-aaaa-1111-aaaa-111111111111', 'q1111111-aaaa-1111-aaaa-111111111111', 'What is a Go variable?', 'A container for storing data'),
    ('qq111112-aaaa-1111-aaaa-111111111111', 'q1111111-aaaa-1111-aaaa-111111111111', 'How do you write a for loop?', 'for init; condition; post {}'),
    ('qq222221-bbbb-2222-bbbb-222222222222', 'q2222222-bbbb-2222-bbbb-222222222222', 'What is a goroutine?', 'Lightweight thread'),
    ('qq222222-bbbb-2222-bbbb-222222222222', 'q2222222-bbbb-2222-bbbb-222222222222', 'How do you create a channel?', 'make(chan Type)');

-- ENROLLMENTS
INSERT INTO enrollments (id, user_id, course_id)
VALUES
    ('eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', '33333333-3333-3333-3333-333333333333', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa'),
    ('ffffffff-ffff-ffff-ffff-ffffffffffff', '44444444-4444-4444-4444-444444444444', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb');

-- MESSAGES
INSERT INTO messages (id, sender_id, receiver_id, content)
VALUES
    ('m1111111-aaaa-1111-aaaa-111111111111', '33333333-3333-3333-3333-333333333333', '22222222-2222-2222-2222-222222222222', 'Hello teacher!'),
    ('m2222222-bbbb-2222-bbbb-222222222222', '22222222-2222-2222-2222-222222222222', '33333333-3333-3333-3333-333333333333', 'Hello student!');

-- NOTIFICATIONS
INSERT INTO notifications (id, user_id, message, read)
VALUES
    ('n1111111-aaaa-1111-aaaa-111111111111', '33333333-3333-3333-3333-333333333333', 'Welcome to the platform!', FALSE),
    ('n2222222-bbbb-2222-bbbb-222222222222', '44444444-4444-4444-4444-444444444444', 'Your course starts tomorrow', FALSE);

-- FORUMS
INSERT INTO forums (id, title, description, created_by)
VALUES
    ('f1111111-aaaa-1111-aaaa-111111111111', 'Go Basics Discussion', 'Discuss Go basics here', '22222222-2222-2222-2222-222222222222');

-- POSTS
INSERT INTO posts (id, forum_id, user_id, content)
VALUES
    ('p1111111-aaaa-1111-aaaa-111111111111', 'f1111111-aaaa-1111-aaaa-111111111111', '33333333-3333-3333-3333-333333333333', 'I have a question about loops'),
    ('p2222222-bbbb-2222-bbbb-222222222222', 'f1111111-aaaa-1111-aaaa-111111111111', '44444444-4444-4444-4444-444444444444', 'What about arrays?');

-- REPLIES
INSERT INTO replies (id, post_id, user_id, content)
VALUES
    ('r1111111-aaaa-1111-aaaa-111111111111', 'p1111111-aaaa-1111-aaaa-111111111111', '22222222-2222-2222-2222-222222222222', 'Loops are explained in chapter 2'),
    ('r2222222-bbbb-2222-bbbb-222222222222', 'p2222222-bbbb-2222-bbbb-222222222222', '22222222-2222-2222-2222-222222222222', 'Arrays are in chapter 3');

-- PROGRESS
INSERT INTO progress (user_id, course_id, completed_modules, completed_quizzes)
VALUES
    ('33333333-3333-3333-3333-333333333333', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', ARRAY['11111111-aaaa-1111-aaaa-111111111111'], '{"qq111111-aaaa-1111-aaaa-111111111111": 1}'),
    ('44444444-4444-4444-4444-444444444444', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', ARRAY['22222222-bbbb-2222-bbbb-222222222222'], '{"qq222221-bbbb-2222-bbbb-222222222222": 1}');

-- CERTIFICATES
INSERT INTO certificates (id, user_id, course_id)
VALUES
    ('c1111111-aaaa-1111-aaaa-111111111111', '33333333-3333-3333-3333-333333333333', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa');

-- ANALYTICS
INSERT INTO analytics (id, event_type, user_id, course_id, metadata)
VALUES
    ('a1111111-aaaa-1111-aaaa-111111111111', 'course_view', '33333333-3333-3333-3333-333333333333', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '{"page":"overview"}');
