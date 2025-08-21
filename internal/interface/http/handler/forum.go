package handler

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Forum represents a forum structure.
type Forum struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// Post represents a post in a forum.
type Post struct {
	ID       string    `json:"id"`
	ForumID  string    `json:"forumId"`
	AuthorID string    `json:"authorId"`
	Content  string    `json:"content"`
	Created  time.Time `json:"created"`
}

// Reply represents a reply to a post.
type Reply struct {
	ID       string    `json:"id"`
	PostID   string    `json:"postId"`
	AuthorID string    `json:"authorId"`
	Content  string    `json:"content"`
	Created  time.Time `json:"created"`
}

// ForumHandler provides HTTP handlers for forum-related endpoints.
type ForumHandler struct{}

// In-memory storage (replace with DB in the future)
var forums = []Forum{}
var posts = []Post{}
var replies = []Reply{}

// CreateForum handles POST /forum
func (h *ForumHandler) CreateForum(c *fiber.Ctx) error {
	var req struct {
		Title string `json:"title"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	forum := Forum{
		ID:    strconv.Itoa(len(forums) + 1),
		Title: req.Title,
	}
	forums = append(forums, forum)
	return c.Status(fiber.StatusCreated).JSON(forum)
}

// GetForum handles GET /forum/:id
func (h *ForumHandler) GetForum(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, f := range forums {
		if f.ID == id {
			return c.JSON(f)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Forum not found"})
}

// ListForums handles GET /forums
func (h *ForumHandler) ListForums(c *fiber.Ctx) error {
	return c.JSON(forums)
}

// CreatePost handles POST /forum/:forum_id/post
func (h *ForumHandler) CreatePost(c *fiber.Ctx) error {
	forumID := c.Params("forum_id")
	var req struct {
		AuthorID string `json:"authorId"`
		Content  string `json:"content"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	post := Post{
		ID:       strconv.Itoa(len(posts) + 1),
		ForumID:  forumID,
		AuthorID: req.AuthorID,
		Content:  req.Content,
		Created:  time.Now(),
	}
	posts = append(posts, post)
	return c.Status(fiber.StatusCreated).JSON(post)
}

// ListPosts handles GET /forum/:forum_id/posts
func (h *ForumHandler) ListPosts(c *fiber.Ctx) error {
	forumID := c.Params("forum_id")
	var forumPosts []Post
	for _, p := range posts {
		if p.ForumID == forumID {
			forumPosts = append(forumPosts, p)
		}
	}
	return c.JSON(forumPosts)
}

// CreateReply handles POST /post/:post_id/reply
func (h *ForumHandler) CreateReply(c *fiber.Ctx) error {
	postID := c.Params("post_id")
	var req struct {
		AuthorID string `json:"authorId"`
		Content  string `json:"content"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	reply := Reply{
		ID:       strconv.Itoa(len(replies) + 1),
		PostID:   postID,
		AuthorID: req.AuthorID,
		Content:  req.Content,
		Created:  time.Now(),
	}
	replies = append(replies, reply)
	return c.Status(fiber.StatusCreated).JSON(reply)
}

// ListReplies handles GET /post/:post_id/replies
func (h *ForumHandler) ListReplies(c *fiber.Ctx) error {
	postID := c.Params("post_id")
	var postReplies []Reply
	for _, r := range replies {
		if r.PostID == postID {
			postReplies = append(postReplies, r)
		}
	}
	return c.JSON(postReplies)
}
