package service

import (
	"TaskManager/internal/domain/model"
	"TaskManager/internal/repository"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
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
	if err := validateUser(user); err != nil {
		return 0, err
	}
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

// validateUser checks if the user data meets the required validation criteria.
func validateUser(user model.User) error {
	if len(user.Username) < 3 || len(user.Username) > 30 {
		return fmt.Errorf("username must be between 3 and 30 characters")
	}

	if !isValidUsername(user.Username) {
		return fmt.Errorf("username can contain only English letters and digits")
	}

	if len(user.Password) < 6 {
		return fmt.Errorf("password must be at least 6 characters")
	}

	if !isValidPassword(user.Password) {
		return fmt.Errorf("password can contain only English letters,digits and symbols (_ , !)")
	}

	return nil
}

// isValidUsername checks if the username contains only valid characters (English letters and digits).
func isValidUsername(username string) bool {
	usernameRegexp := regexp.MustCompile(`^[a-zA-Z0-9]$`)
	return usernameRegexp.MatchString(username)
}

// isValidPassword checks if the password contains only valid characters (English letters, digits, and specific symbols).
func isValidPassword(password string) bool {
	passwordRegexp := regexp.MustCompile(`^[a-zA-Z0-9!_]$`)
	return passwordRegexp.MatchString(password)
}

// generatePasswordHash hashes the user's password using bcrypt.
func generatePasswordHash(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}
