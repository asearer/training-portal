CREATE TABLE notifications (
                               id UUID PRIMARY KEY,
                               user_id UUID REFERENCES users(id) ON DELETE CASCADE,
                               message TEXT NOT NULL,
                               read BOOLEAN DEFAULT FALSE,
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
