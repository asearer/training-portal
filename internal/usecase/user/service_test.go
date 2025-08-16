package user

import (
	"errors"
	"testing"

	"training-portal/internal/domain/user"
	"training-portal/internal/interface/repository"
	"golang.org/x/crypto/bcrypt"
)

// MockUserRepository is a mock implementation of the user repository
type MockUserRepository struct {
	users    map[string]*user.User
	nextID   int
	shouldFail bool
}

func NewMockUserRepository() repository.UserRepository {
	return &MockUserRepository{
		users:  make(map[string]*user.User),
		nextID: 1,
	}
}

func (m *MockUserRepository) Create(u *user.User) error {
	if m.shouldFail {
		return errors.New("database error")
	}
	if u.ID == "" {
		u.ID = string(rune(m.nextID))
		m.nextID++
	}
	m.users[u.ID] = u
	return nil
}

func (m *MockUserRepository) FindByID(id string) (*user.User, error) {
	if m.shouldFail {
		return nil, errors.New("database error")
	}
	if u, exists := m.users[id]; exists {
		return u, nil
	}
	return nil, nil
}

func (m *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	if m.shouldFail {
		return nil, errors.New("database error")
	}
	for _, u := range m.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, nil
}

func (m *MockUserRepository) Update(u *user.User) error {
	if m.shouldFail {
		return errors.New("database error")
	}
	if _, exists := m.users[u.ID]; !exists {
		return errors.New("user not found")
	}
	m.users[u.ID] = u
	return nil
}

func (m *MockUserRepository) Delete(id string) error {
	if m.shouldFail {
		return errors.New("database error")
	}
	if _, exists := m.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(m.users, id)
	return nil
}

func (m *MockUserRepository) List() ([]*user.User, error) {
	if m.shouldFail {
		return nil, errors.New("database error")
	}
	users := make([]*user.User, 0, len(m.users))
	for _, u := range m.users {
		users = append(users, u)
	}
	return users, nil
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{"Valid email", "test@example.com", true},
		{"Valid email with subdomain", "test@sub.example.com", true},
		{"Valid email with plus", "test+tag@example.com", true},
		{"Valid email with underscore", "test_user@example.com", true},
		{"Invalid email - no @", "testexample.com", false},
		{"Invalid email - no domain", "test@", false},
		{"Invalid email - no local part", "@example.com", false},
		{"Invalid email - spaces", "test @example.com", false},
		{"Empty email", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateEmail(tt.email); got != tt.expected {
				t.Errorf("ValidateEmail(%q) = %v, want %v", tt.email, got, tt.expected)
			}
		})
	}
}

func TestUserService_Register(t *testing.T) {
	tests := []struct {
		name        string
		nameInput   string
		emailInput  string
		password    string
		role        user.Role
		shouldFail  bool
		expectError bool
	}{
		{
			name:        "Valid registration",
			nameInput:   "John Doe",
			emailInput:  "john@example.com",
			password:    "password123",
			role:        user.RoleEmployee,
			shouldFail:  false,
			expectError: false,
		},
		{
			name:        "Empty name",
			nameInput:   "",
			emailInput:  "john@example.com",
			password:    "password123",
			role:        user.RoleEmployee,
			shouldFail:  false,
			expectError: true,
		},
		{
			name:        "Empty email",
			nameInput:   "John Doe",
			emailInput:  "",
			password:    "password123",
			role:        user.RoleEmployee,
			shouldFail:  false,
			expectError: true,
		},
		{
			name:        "Empty password",
			nameInput:   "John Doe",
			emailInput:  "john@example.com",
			password:    "",
			role:        user.RoleEmployee,
			shouldFail:  false,
			expectError: true,
		},
		{
			name:        "Invalid email format",
			nameInput:   "John Doe",
			emailInput:  "invalid-email",
			password:    "password123",
			role:        user.RoleEmployee,
			shouldFail:  false,
			expectError: true,
		},
		{
			name:        "Email already exists",
			nameInput:   "John Doe",
			emailInput:  "existing@example.com",
			password:    "password123",
			role:        user.RoleEmployee,
			shouldFail:  false,
			expectError: true,
		},
		{
			name:        "Database error",
			nameInput:   "John Doe",
			emailInput:  "john@example.com",
			password:    "password123",
			role:        user.RoleEmployee,
			shouldFail:  true,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockUserRepository().(*MockUserRepository)
			
			// Add existing user for duplicate email test
			if tt.name == "Email already exists" {
				existingUser := &user.User{
					ID:    "1",
					Name:  "Existing User",
					Email: "existing@example.com",
					Role:  user.RoleEmployee,
				}
				mockRepo.Create(existingUser)
			}
			
			if tt.shouldFail {
				mockRepo.shouldFail = true
			}

			service := &UserService{Repo: mockRepo}
			user, err := service.Register(tt.nameInput, tt.emailInput, tt.password, tt.role)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if user == nil {
					t.Errorf("Expected user but got nil")
				}
				if user.Name != tt.nameInput {
					t.Errorf("Expected name %s, got %s", tt.nameInput, user.Name)
				}
				if user.Email != tt.emailInput {
					t.Errorf("Expected email %s, got %s", tt.emailInput, user.Email)
				}
				if user.Role != tt.role {
					t.Errorf("Expected role %s, got %s", tt.role, user.Role)
				}
				if user.Password == tt.password {
					t.Errorf("Password should be hashed, not plain text")
				}
			}
		})
	}
}

func TestUserService_Login(t *testing.T) {
	tests := []struct {
		name        string
		email       string
		password    string
		setupUser   bool
		shouldFail  bool
		expectError bool
	}{
		{
			name:        "Valid login",
			email:       "john@example.com",
			password:    "password123",
			setupUser:   true,
			shouldFail:  false,
			expectError: false,
		},
		{
			name:        "User not found",
			email:       "nonexistent@example.com",
			password:    "password123",
			setupUser:   false,
			shouldFail:  false,
			expectError: true,
		},
		{
			name:        "Wrong password",
			email:       "john@example.com",
			password:    "wrongpassword",
			setupUser:   true,
			shouldFail:  false,
			expectError: true,
		},
		{
			name:        "Database error",
			email:       "john@example.com",
			password:    "password123",
			setupUser:   true,
			shouldFail:  true,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockUserRepository().(*MockUserRepository)
			
			if tt.setupUser {
				// Create a user with hashed password
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
				testUser := &user.User{
					ID:       "1",
					Name:     "John Doe",
					Email:    "john@example.com",
					Password: string(hashedPassword),
					Role:     user.RoleEmployee,
				}
				mockRepo.Create(testUser)
			}
			
			if tt.shouldFail {
				mockRepo.shouldFail = true
			}

			service := &UserService{Repo: mockRepo}
			user, err := service.Login(tt.email, tt.password)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if user == nil {
					t.Errorf("Expected user but got nil")
				}
				if user.Email != tt.email {
					t.Errorf("Expected email %s, got %s", tt.email, user.Email)
				}
			}
		})
	}
}

func TestUserService_GetUser(t *testing.T) {
	tests := []struct {
		name        string
		userID      string
		setupUser   bool
		shouldFail  bool
		expectError bool
	}{
		{
			name:        "Valid user ID",
			userID:      "1",
			setupUser:   true,
			shouldFail:  false,
			expectError: false,
		},
		{
			name:        "Empty user ID",
			userID:      "",
			setupUser:   false,
			shouldFail:  false,
			expectError: true,
		},
		{
			name:        "User not found",
			userID:      "999",
			setupUser:   false,
			shouldFail:  false,
			expectError: true,
		},
		{
			name:        "Database error",
			userID:      "1",
			setupUser:   true,
			shouldFail:  true,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockUserRepository().(*MockUserRepository)
			
			if tt.setupUser {
				testUser := &user.User{
					ID:    "1",
					Name:  "John Doe",
					Email: "john@example.com",
					Role:  user.RoleEmployee,
				}
				mockRepo.Create(testUser)
			}
			
			if tt.shouldFail {
				mockRepo.shouldFail = true
			}

			service := &UserService{Repo: mockRepo}
			user, err := service.GetUser(tt.userID)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if user == nil {
					t.Errorf("Expected user but got nil")
				}
				if user.ID != tt.userID {
					t.Errorf("Expected ID %s, got %s", tt.userID, user.ID)
				}
			}
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	tests := []struct {
		name        string
		user        *user.User
		setupUser   bool
		shouldFail  bool
		expectError bool
	}{
		{
			name: "Valid update",
			user: &user.User{
				ID:    "1",
				Name:  "John Updated",
				Email: "john.updated@example.com",
				Role:  user.RoleAdmin,
			},
			setupUser:   true,
			shouldFail:  false,
			expectError: false,
		},
		{
			name: "Empty ID",
			user: &user.User{
				ID:    "",
				Name:  "John Updated",
				Email: "john.updated@example.com",
				Role:  user.RoleAdmin,
			},
			setupUser:   false,
			shouldFail:  false,
			expectError: true,
		},
		{
			name: "Invalid email format",
			user: &user.User{
				ID:    "1",
				Name:  "John Updated",
				Email: "invalid-email",
				Role:  user.RoleAdmin,
			},
			setupUser:   true,
			shouldFail:  false,
			expectError: true,
		},
		{
			name: "Database error",
			user: &user.User{
				ID:    "1",
				Name:  "John Updated",
				Email: "john.updated@example.com",
				Role:  user.RoleAdmin,
			},
			setupUser:   true,
			shouldFail:  true,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockUserRepository().(*MockUserRepository)
			
			if tt.setupUser {
				testUser := &user.User{
					ID:    "1",
					Name:  "John Doe",
					Email: "john@example.com",
					Role:  user.RoleEmployee,
				}
				mockRepo.Create(testUser)
			}
			
			if tt.shouldFail {
				mockRepo.shouldFail = true
			}

			service := &UserService{Repo: mockRepo}
			err := service.UpdateUser(tt.user)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestUserService_UpdatePassword(t *testing.T) {
	tests := []struct {
		name         string
		userID       string
		newPassword  string
		setupUser    bool
		shouldFail   bool
		expectError  bool
	}{
		{
			name:        "Valid password update",
			userID:      "1",
			newPassword: "newpassword123",
			setupUser:   true,
			shouldFail:  false,
			expectError: false,
		},
		{
			name:        "Empty user ID",
			userID:      "",
			newPassword: "newpassword123",
			setupUser:   false,
			shouldFail:  false,
			expectError: true,
		},
		{
			name:        "Empty new password",
			userID:      "1",
			newPassword: "",
			setupUser:   true,
			shouldFail:  false,
			expectError: true,
		},
		{
			name:        "User not found",
			userID:      "999",
			newPassword: "newpassword123",
			setupUser:   false,
			shouldFail:  false,
			expectError: true,
		},
		{
			name:        "Database error",
			userID:      "1",
			newPassword: "newpassword123",
			setupUser:   true,
			shouldFail:  true,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockUserRepository().(*MockUserRepository)
			
			if tt.setupUser {
				testUser := &user.User{
					ID:    "1",
					Name:  "John Doe",
					Email: "john@example.com",
					Role:  user.RoleEmployee,
				}
				mockRepo.Create(testUser)
			}
			
			if tt.shouldFail {
				mockRepo.shouldFail = true
			}

			service := &UserService{Repo: mockRepo}
			err := service.UpdatePassword(tt.userID, tt.newPassword)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	tests := []struct {
		name        string
		userID      string
		setupUser   bool
		shouldFail  bool
		expectError bool
	}{
		{
			name:        "Valid delete",
			userID:      "1",
			setupUser:   true,
			shouldFail:  false,
			expectError: false,
		},
		{
			name:        "Empty user ID",
			userID:      "",
			setupUser:   false,
			shouldFail:  false,
			expectError: true,
		},
		{
			name:        "User not found",
			userID:      "999",
			setupUser:   false,
			shouldFail:  false,
			expectError: true,
		},
		{
			name:        "Database error",
			userID:      "1",
			setupUser:   true,
			shouldFail:  true,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockUserRepository().(*MockUserRepository)
			
			if tt.setupUser {
				testUser := &user.User{
					ID:    "1",
					Name:  "John Doe",
					Email: "john@example.com",
					Role:  user.RoleEmployee,
				}
				mockRepo.Create(testUser)
			}
			
			if tt.shouldFail {
				mockRepo.shouldFail = true
			}

			service := &UserService{Repo: mockRepo}
			err := service.DeleteUser(tt.userID)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestUserService_ListUsers(t *testing.T) {
	tests := []struct {
		name        string
		setupUsers  bool
		shouldFail  bool
		expectError bool
		expectedCount int
	}{
		{
			name:         "List users successfully",
			setupUsers:   true,
			shouldFail:   false,
			expectError:  false,
			expectedCount: 2,
		},
		{
			name:         "Empty user list",
			setupUsers:   false,
			shouldFail:   false,
			expectError:  false,
			expectedCount: 0,
		},
		{
			name:         "Database error",
			setupUsers:   true,
			shouldFail:   true,
			expectError:  true,
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockUserRepository().(*MockUserRepository)
			
			if tt.setupUsers {
				user1 := &user.User{
					ID:    "1",
					Name:  "John Doe",
					Email: "john@example.com",
					Role:  user.RoleEmployee,
				}
				user2 := &user.User{
					ID:    "2",
					Name:  "Jane Smith",
					Email: "jane@example.com",
					Role:  user.RoleAdmin,
				}
				mockRepo.Create(user1)
				mockRepo.Create(user2)
			}
			
			if tt.shouldFail {
				mockRepo.shouldFail = true
			}

			service := &UserService{Repo: mockRepo}
			users, err := service.ListUsers()

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if len(users) != tt.expectedCount {
					t.Errorf("Expected %d users, got %d", tt.expectedCount, len(users))
				}
			}
		})
	}
}
