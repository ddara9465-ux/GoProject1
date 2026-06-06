package auth

import (
	"context"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// AdminAuth проверка администратора
func (s *AuthService) AdminAuth(ctx context.Context, login, password string) (int, error) {
	passwordHash, userID, isAdmin := s.repo.AdminAuth(ctx, login)

	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		log.Printf("Login func | Неверный пароль для пользователя login=%v | Ошибка %v", login, err)
		return 0, err
	}

	log.Printf("Login func | Пароль для login = %v подтвержден", login)

	if !isAdmin {
		return 0, errors.New("пользователь не является администратором")
	}

	return userID, nil
}
