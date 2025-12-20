package auth

import (
	"GoProject1/internal/adapter/persistence"
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// UC_Login — логин: (userID, true) если ок, иначе (0, false).
func UC_Login(login string, password string) (int, bool) {
	// Берём хеш пароля и id по логину.
	trueHashPassword, userID, err := persistence.A_GetPasswordHash_User(context.Background(), login)
	if err != nil {
		// Ошибка БД/пользователь не найден.
		log.Printf("Log err: %v", err)
		return 0, false
	}

	// Проверяем пароль через bcrypt: nil = пароль подошёл, error = не подошёл/хеш битый.
	err = bcrypt.CompareHashAndPassword([]byte(trueHashPassword), []byte(password))
	if err != nil {
		log.Printf("Login func | Неверный пароль для пользователя login=%v", login)
		return 0, false
	}

	log.Print("Login func | Данные авторизации подтверждены")
	return userID, true
}
