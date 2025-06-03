package service

import (
	"TaskManager/internal/domain/model"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestValidateUser_ValidInput(t *testing.T) {
	testTable := []struct {
		name     string
		username string
		password string
		wantErr  bool
	}{
		{
			name:     "Valid User",
			username: "validUser12",
			password: "ValidPassword123!",
			wantErr:  false,
		},
		{
			name:     "valid User with special characters",
			username: "ValidUser6473",
			password: "Valid_Password!123",
			wantErr:  false,
		},
		{
			name:     "valid User without special characters",
			username: "ValidUser6473",
			password: "Valid!Password_123",
			wantErr:  false,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			user := model.User{
				Username: tt.username,
				Password: tt.password,
			}
			err := validateUser(user)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateUser_InvalidUsername(t *testing.T) {
	testTable := []struct {
		name          string
		username      string
		password      string
		expectedError string
	}{
		{
			name:          "Short Username",
			username:      "ab",
			password:      "ValidPassword123!",
			expectedError: "username must be between 3 and 30 characters",
		},
		{
			name:          "Long Username",
			username:      "thisUsernameIsWayTooLongForValidation",
			password:      "ValidPassword123!",
			expectedError: "username must be between 3 and 30 characters",
		},
		{
			name:          "Invalid Characters in Username",
			username:      "Invalid@User",
			password:      "ValidPassword123!",
			expectedError: "username can contain only English letters and digits",
		},
		{
			name:          "Empty Username",
			username:      "",
			password:      "ValidPassword123!",
			expectedError: "username must be between 3 and 30 characters",
		},
		{
			name:          "Non-English Characters in Username",
			username:      "Пользователь",
			password:      "ValidPassword123!",
			expectedError: "username can contain only English letters and digits",
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			user := model.User{
				Username: tt.username,
				Password: tt.password,
			}
			err := validateUser(user)
			if err == nil || err.Error() != tt.expectedError {
				t.Errorf("validateUser() error = %v, expectedError %v", err, tt.expectedError)
			}
		})
	}
}

func TestValidateUser_InvalidPassword(t *testing.T) {
	testTable := []struct {
		name          string
		username      string
		password      string
		expectedError string
	}{
		{
			name:          "Short Password",
			username:      "validUser",
			password:      "123",
			expectedError: "password must be at least 6 characters",
		},
		{
			name:          "Invalid Characters in Password",
			username:      "validUser",
			password:      "Valid@Password",
			expectedError: "password can contain only English letters,digits and symbols (_ , !)",
		},
		{
			name:          "Empty Password",
			username:      "validUser",
			password:      "",
			expectedError: "password must be at least 6 characters",
		},
		{
			name:          "Non-English Characters in Password",
			username:      "validUser",
			password:      "Пароль123",
			expectedError: "password can contain only English letters,digits and symbols (_ , !)",
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			user := model.User{
				Username: tt.username,
				Password: tt.password,
			}
			err := validateUser(user)
			if err == nil || err.Error() != tt.expectedError {
				t.Errorf("validateUser() error = %v, expectedError %v", err, tt.expectedError)
			}
		})
	}
}

func TestIsValidUsername(t *testing.T) {
	testTable := []struct {
		name     string
		username string
		expected bool
	}{
		{
			name:     "Valid Username",
			username: "validUser123",
			expected: true,
		},
		{
			name:     "Invalid Username with Special Characters",
			username: "invalid@user",
			expected: false,
		},
		{
			name:     "Invalid Username with Spaces",
			username: "invalid user",
			expected: false,
		},
		{
			name:     "Valid Username with Digits",
			username: "user123",
			expected: true,
		},
		{
			name:     "Invalid Username with Non-English Characters",
			username: "Пользователь",
			expected: false,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidUsername(tt.username)
			if result != tt.expected {
				t.Errorf("isValidUsername(%s) = %v, want %v", tt.username, result, tt.expected)
			}
		})
	}
}

func TestIsValidPassword(t *testing.T) {
	testTable := []struct {
		name     string
		password string
		expected bool
	}{
		{
			name:     "Valid Password",
			password: "ValidPassword123!",
			expected: true,
		},
		{
			name:     "Invalid Password with Special Characters",
			password: "Invalid@Password",
			expected: false,
		},
		{
			name:     "Invalid Password with Spaces",
			password: "Invalid Password",
			expected: false,
		},
		{
			name:     "Valid Password with Digits and Symbols",
			password: "Valid_Password!123",
			expected: true,
		},
		{
			name:     "Invalid Password with Non-English Characters",
			password: "Пароль123",
			expected: false,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidPassword(tt.password)
			if result != tt.expected {
				t.Errorf("isValidPassword(%s) = %v, want %v", tt.password, result, tt.expected)
			}
		})
	}
}

func TestGeneratePasswordHash(t *testing.T) {
	testTable := []struct {
		name     string
		password string
	}{
		{
			name:     "Valid Password",
			password: "ValidPassword123!",
		},
		{
			name:     "Short Password",
			password: "Short1!",
		},
		{
			name:     "Password with Special Characters",
			password: "Special@Password!123",
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword := generatePasswordHash(tt.password)
			if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(tt.password)); err != nil {
				t.Errorf("generatePasswordHash() failed for %s: %v", tt.name, err)
			}
		})
	}
}
