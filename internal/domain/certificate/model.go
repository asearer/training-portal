package certificate

// Certificate represents a course completion certificate.
type Certificate struct {
	ID          string // UUID
	UserID      string // UUID of the user who earned the certificate
	CourseID    string // UUID of the course
	IssuedAt    int64  // Unix timestamp of issuance
	DownloadURL string // URL to download the certificate PDF or badge
}
