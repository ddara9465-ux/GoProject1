package auth

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Register - регистрация, возвращает (userID, ok)
func (s *AuthService) Register(ctx context.Context, password, first_name, last_name, phone, email string) (int, bool) {
	hashed_password, err := hashPassword(password)
	if err != nil {
		return 0, false
	}

	userID, err2 := s.repo.RegisterTransaction(ctx, hashed_password, first_name, last_name, phone, email)
	if err2 {
		return 0, false
	}

	return userID, true
}
