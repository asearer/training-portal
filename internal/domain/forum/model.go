package forum

// Forum represents a discussion forum for a course or module.
type Forum struct {
	ID        string // UUID
	Title     string
	CourseID  string // optional: link to course
	ModuleID  string // optional: link to module
	CreatedBy string // user ID
	CreatedAt int64  // Unix timestamp
}

// Post represents a post in a forum.
type Post struct {
	ID        string // UUID
	ForumID   string
	UserID    string
	Content   string
	CreatedAt int64 // Unix timestamp
}

// Reply represents a reply to a post.
type Reply struct {
	ID        string // UUID
	PostID    string
	UserID    string
	Content   string
	CreatedAt int64 // Unix timestamp
}
