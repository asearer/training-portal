CREATE TABLE forums (
                        id UUID PRIMARY KEY,
                        title VARCHAR(255) NOT NULL,
                        created_by UUID REFERENCES users(id) ON DELETE SET NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE posts (
                       id UUID PRIMARY KEY,
                       forum_id UUID REFERENCES forums(id) ON DELETE CASCADE,
                       user_id UUID REFERENCES users(id) ON DELETE SET NULL,
                       content TEXT NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE replies (
                         id UUID PRIMARY KEY,
                         post_id UUID REFERENCES posts(id) ON DELETE CASCADE,
                         user_id UUID REFERENCES users(id) ON DELETE SET NULL,
                         content TEXT NOT NULL,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
