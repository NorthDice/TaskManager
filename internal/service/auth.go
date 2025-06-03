package service

import (
	"TaskManager/internal/domain/model"
	"TaskManager/internal/repository"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

// AuthService defines the interface for user authentication and authorization operations.
type AuthService struct {
	repo repository.Authorization
}

const (
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId string `json:"user_id" bson:"user_id"`
}

// NewAuthService initializes a new AuthService instance with the provided repository.
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser creates a new user in the system.
func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// GenerateToken generates a JWT token for the user based on their username and password.
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, password)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: user.Id.Hex(),
	})

	return token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
}

// ParseToken parses a JWT token and returns the user ID associated with it.
func (s *AuthService) ParseToken(tokenString string) (string, error) {
	parsedToken, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SIGNING_KEY")), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := parsedToken.Claims.(*tokenClaims)
	if !ok || !parsedToken.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}
