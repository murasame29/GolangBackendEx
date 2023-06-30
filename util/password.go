package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashedPasswordはbcryptによってハッシュ化されたパスワードをかえす
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// bcryptによってハッシュ化されたパスワードを検証する
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
