package auth

import (
	"auth/internal/models"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const (
	createUseQuery = `INSERT INTO users (username, email, password, status) VALUES ($1, $2, $3, $4) RETURNING id`
)

func (r *PostgresAuthRepository) CreateUser(user models.User) (int, error) {
	_, err := r.FindUser(user.Name, user.Password)
	if err != nil && err.Error() != "user not found" {
		return 0, fmt.Errorf("error creating user: %w", err)
	}

	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("error hashing password: %w", err)
	}

	// Save user with the hashed password
	var userID int
	err = r.conn.QueryRow(context.Background(), createUseQuery, user.Name, user.Email, string(hashedPassword), user.Status).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}

	return userID, nil
}
