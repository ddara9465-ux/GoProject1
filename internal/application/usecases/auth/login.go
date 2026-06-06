package auth

import (
	"GoProject1/internal/adapter/persistence"
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *persistence.Repository
}

// NewAuthService создаёт новый сервис аутентификации
func NewAuthService(repo *persistence.Repository) *AuthService {
	return &AuthService{repo: repo}
}

// Login - логин: (userID, true) если ок, иначе (0, false)
func (s *AuthService) Login(ctx context.Context, login, password string) (int, bool) {
	trueHashPassword, userID, err := s.repo.GetPasswordHashByLogin(ctx, login)

	if err != nil {
		log.Printf("Log err: %v", err)
		return 0, false
	}

	err = bcrypt.CompareHashAndPassword([]byte(trueHashPassword), []byte(password))
	if err != nil {
		log.Printf("Login func | Неверный пароль для пользователя login=%v", login)
		return 0, false
	}

	log.Print("Login func | Данные авторизации подтверждены")
	return userID, true
}
