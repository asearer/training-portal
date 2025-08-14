package role

// Role represents a user role in the system.
type Role string

const (
	RoleEmployee Role = "employee"
	RoleAdmin    Role = "admin"
	RoleTrainer  Role = "trainer"
	RoleManager  Role = "manager"
	RoleGuest    Role = "guest"
)

// Permission represents a named permission for fine-grained access control.
type Permission string

// RolePermission maps roles to permissions (stub for future expansion).
type RolePermission struct {
	Role       Role
	Permission Permission
}

// UserRoleAssignment represents a user's assigned role(s).
type UserRoleAssignment struct {
	UserID string
	Roles  []Role
}
