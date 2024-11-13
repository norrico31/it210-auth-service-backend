package user

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/norrico31/it210-auth-service-backend/entities"
	"golang.org/x/crypto/bcrypt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// GenerateJWT generates a JWT for the authenticated user
func GenerateJWT(user entities.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", fmt.Errorf("JWT_SECRET is not set")
	}

	claims := jwt.MapClaims{
		"user_id":    user.ID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"exp":        time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// Login authenticates a user by email and password, returning a JWT and user details if successful
func (s *Store) Login(payload entities.UserLoginPayload) (string, entities.User, error) {
	var user entities.User
	err := s.db.QueryRow(`
		SELECT id, firstName, lastName, email, password, age, lastActiveAt, createdAt, updatedAt, deletedAt
		FROM users WHERE email = $1 AND deletedAt IS NULL`, payload.Email).Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password,
		&user.Age, &user.LastActiveAt, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
	)
	if err == sql.ErrNoRows {
		return "", user, fmt.Errorf("user not found")
	} else if err != nil {
		return "", user, fmt.Errorf("failed to query user: %v", err)
	}

	// Verify the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return "", user, fmt.Errorf("invalid password")
	}

	// Generate JWT token
	token, err := GenerateJWT(user)
	if err != nil {
		return "", user, fmt.Errorf("failed to generate token: %v", err)
	}

	// Update user's last active timestamp

	_, err = s.db.Exec(`UPDATE users SET lastActiveAt = NULL WHERE id = $1`, user.ID)
	if err != nil {
		log.Printf("Failed to update last active timestamp for user %d: %v", user.ID, err)
	}

	return token, user, nil
}
