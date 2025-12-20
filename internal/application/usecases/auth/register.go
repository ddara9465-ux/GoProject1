package auth

import (
	"GoProject1/internal/adapter/persistence"

	"golang.org/x/crypto/bcrypt"
)

// "засекречиваем" пароль
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// UC_Register: регистрация, возвращает (userID, ok).
func UC_Register(password string, first_name string, last_name string, phone string, email string) (int, bool) {
	hashed_password, err := hashPassword(password)
	if err != nil {
		return 0, false
	}

	userID, err2 := persistence.A_RegisterTransaction(hashed_password, first_name, last_name, phone, email)
	if err2 {
		return 0, false
	}

	return userID, true
}
