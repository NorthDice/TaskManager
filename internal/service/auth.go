package service

import (
	"TaskManager/internal/domain/model"
	"TaskManager/internal/repository"
)

// AuthService defines the interface for user authentication and authorization operations.
type AuthService struct {
	repo repository.Authorization
}

// NewAuthService initializes a new AuthService instance with the provided repository.
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser creates a new user in the system.
func (s *AuthService) CreateUser(user model.User) (int, error) {
	return s.repo.CreateUser(user)
}

// GenerateToken generates a JWT token for the user based on their username and password.
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	return s.repo.GenerateToken(username, password)
}

// ParseToken parses a JWT token and returns the user ID associated with it.
func (s *AuthService) ParseToken(tokenString string) (int, error) {
	return s.repo.ParseToken(tokenString)
}
