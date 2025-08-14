package analytics

// Stub for analytics domain model

type UserEngagement struct {
	UserID           string
	CourseID         string
	ModulesCompleted int
	TotalModules     int
	LastActive       int64 // Unix timestamp
}

type CourseAnalytics struct {
	CourseID        string
	Enrollments     int
	Completions     int
	AverageProgress float64
}

type AnalyticsEvent struct {
	ID        string
	UserID    string
	EventType string
	Timestamp int64
	Metadata  map[string]interface{}
}
