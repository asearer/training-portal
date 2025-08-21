package handler

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Certificate represents a user's certificate for a completed course.
// In a real implementation, replace with a DB model and store files as needed.
type Certificate struct {
	ID       string    `json:"id"`
	UserID   string    `json:"userId"`
	CourseID string    `json:"courseId"`
	IssuedAt time.Time `json:"issuedAt"`
	FileURL  string    `json:"fileUrl"` // Could point to a PDF or image
}

// CertificateHandler provides HTTP handlers for certificate-related endpoints.
type CertificateHandler struct{}

// In-memory storage for demonstration (replace with DB in the future)
var certificates = []Certificate{}

// GetCertificate handles GET /certificate/:id
func (h *CertificateHandler) GetCertificate(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, cert := range certificates {
		if cert.ID == id {
			return c.JSON(cert)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Certificate not found"})
}

// ListCertificates handles GET /user/:user_id/certificates
func (h *CertificateHandler) ListCertificates(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User ID required"})
	}

	var userCertificates []Certificate
	for _, cert := range certificates {
		if cert.UserID == userID {
			userCertificates = append(userCertificates, cert)
		}
	}

	return c.JSON(userCertificates)
}

// DownloadCertificate handles GET /certificate/:id/download
func (h *CertificateHandler) DownloadCertificate(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, cert := range certificates {
		if cert.ID == id {
			// Stub: In a real system, return file bytes or redirect to file storage
			return c.JSON(fiber.Map{
				"message": "Download endpoint (stub)",
				"fileUrl": cert.FileURL,
			})
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Certificate not found"})
}

// Optional helper: create a new certificate (could be called after course completion)
func (h *CertificateHandler) IssueCertificate(userID, courseID string) Certificate {
	cert := Certificate{
		ID:       strconv.Itoa(len(certificates) + 1),
		UserID:   userID,
		CourseID: courseID,
		IssuedAt: time.Now(),
		FileURL:  "/path/to/certificate.pdf", // Replace with real file generation
	}
	certificates = append(certificates, cert)
	return cert
}
