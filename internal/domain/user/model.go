// File: internal/domain/user/model.go
// Defines the user domain model

package user

type Role string

const (
    RoleEmployee Role = "employee"
    RoleAdmin    Role = "admin"
    RoleTrainer  Role = "trainer"
)

type User struct {
    ID       string // UUID
    Name     string
    Email    string
    Password string // hashed password
    Role     Role
}
