CREATE TABLE messages (
                          id UUID PRIMARY KEY,
                          sender_id UUID REFERENCES users(id) ON DELETE CASCADE,
                          receiver_id UUID REFERENCES users(id) ON DELETE CASCADE,
                          content TEXT NOT NULL,
                          read BOOLEAN DEFAULT FALSE,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
