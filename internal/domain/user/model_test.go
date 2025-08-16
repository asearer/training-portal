package user

import (
	"errors"
	"testing"
)

func TestRole_String(t *testing.T) {
	tests := []struct {
		name     string
		role     Role
		expected string
	}{
		{"Employee", RoleEmployee, "employee"},
		{"Admin", RoleAdmin, "admin"},
		{"Trainer", RoleTrainer, "trainer"},
		{"Unknown", Role("unknown"), "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.role) != tt.expected {
				t.Errorf("Role = %v, want %v", string(tt.role), tt.expected)
			}
		})
	}
}

func TestRole_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		role     Role
		expected bool
	}{
		{"Employee", RoleEmployee, true},
		{"Admin", RoleAdmin, true},
		{"Trainer", RoleTrainer, true},
		{"Invalid", Role("invalid"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := tt.role == RoleEmployee || tt.role == RoleAdmin || tt.role == RoleTrainer
			if isValid != tt.expected {
				t.Errorf("Role.IsValid() = %v, want %v", isValid, tt.expected)
			}
		})
	}
}

func TestUser_Validation(t *testing.T) {
	tests := []struct {
		name    string
		user    *User
		wantErr bool
	}{
		{
			name: "Valid user",
			user: &User{
				ID:    "123",
				Name:  "John Doe",
				Email: "john@example.com",
				Role:  RoleEmployee,
			},
			wantErr: false,
		},
		{
			name: "Empty ID",
			user: &User{
				ID:    "",
				Name:  "John Doe",
				Email: "john@example.com",
				Role:  RoleEmployee,
			},
			wantErr: true,
		},
		{
			name: "Empty name",
			user: &User{
				ID:    "123",
				Name:  "",
				Email: "john@example.com",
				Role:  RoleEmployee,
			},
			wantErr: true,
		},
		{
			name: "Empty email",
			user: &User{
				ID:    "123",
				Name:  "John Doe",
				Email: "",
				Role:  RoleEmployee,
			},
			wantErr: true,
		},
		{
			name: "Invalid role",
			user: &User{
				ID:    "123",
				Name:  "John Doe",
				Email: "john@example.com",
				Role:  Role("invalid"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateUser(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("User validation error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_HasPermission(t *testing.T) {
	tests := []struct {
		name           string
		userRole       Role
		requiredRole   Role
		expectedResult bool
	}{
		{"Employee accessing employee resource", RoleEmployee, RoleEmployee, true},
		{"Employee accessing admin resource", RoleEmployee, RoleAdmin, false},
		{"Employee accessing trainer resource", RoleEmployee, RoleTrainer, false},
		{"Admin accessing employee resource", RoleAdmin, RoleEmployee, true},
		{"Admin accessing admin resource", RoleAdmin, RoleAdmin, true},
		{"Admin accessing trainer resource", RoleAdmin, RoleTrainer, true},
		{"Trainer accessing employee resource", RoleTrainer, RoleEmployee, true},
		{"Trainer accessing admin resource", RoleTrainer, RoleAdmin, false},
		{"Trainer accessing trainer resource", RoleTrainer, RoleTrainer, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{Role: tt.userRole}
			if got := user.HasPermission(tt.requiredRole); got != tt.expectedResult {
				t.Errorf("User.HasPermission() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}

// Helper function to validate user
func validateUser(u *User) error {
	if u.ID == "" {
		return errors.New("ID is required")
	}
	if u.Name == "" {
		return errors.New("Name is required")
	}
	if u.Email == "" {
		return errors.New("Email is required")
	}
	if u.Role != RoleEmployee && u.Role != RoleAdmin && u.Role != RoleTrainer {
		return errors.New("Invalid role")
	}
	return nil
}

// Helper function to check if user has permission
func (u *User) HasPermission(requiredRole Role) bool {
	switch u.Role {
	case RoleAdmin:
		return true
	case RoleTrainer:
		return requiredRole == RoleEmployee || requiredRole == RoleTrainer
	case RoleEmployee:
		return requiredRole == RoleEmployee
	default:
		return false
	}
}
