package message

// Message represents a direct message between users.
type Message struct {
	ID         string // UUID
	SenderID   string // User ID of sender
	ReceiverID string // User ID of receiver
	Content    string
	SentAt     int64 // Unix timestamp
	Read       bool
}
