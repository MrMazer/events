package auth

import (
	"auth/internal/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	findUserQuery = `SELECT id, username, email, password, status FROM users WHERE username=$1`
)

func (r *PostgresAuthRepository) FindUser(username, password string) (*models.User, error) {
	var user models.User
	err := r.conn.QueryRow(context.Background(), findUserQuery, username).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error finding user: %w", err)
	}

	// Compare the provided password with the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	return &user, nil
}
