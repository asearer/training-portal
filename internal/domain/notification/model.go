package notification

// NotificationType represents the type of notification (e.g., email, in-app, etc.)
type NotificationType string

const (
	TypeEmail NotificationType = "email"
	TypeInApp NotificationType = "in_app"
	TypeSMS   NotificationType = "sms"
)

// NotificationStatus represents the delivery/read status of a notification
type NotificationStatus string

const (
	StatusUnread NotificationStatus = "unread"
	StatusRead   NotificationStatus = "read"
	StatusSent   NotificationStatus = "sent"
	StatusFailed NotificationStatus = "failed"
)

// Notification represents a user notification in the system
type Notification struct {
	ID        string             // UUID
	UserID    string             // User to notify
	Type      NotificationType   // Notification channel/type
	Title     string             // Short title
	Message   string             // Notification message
	Status    NotificationStatus // Read/sent status
	CreatedAt int64              // Unix timestamp
	ReadAt    *int64             // Unix timestamp (nullable)
}
