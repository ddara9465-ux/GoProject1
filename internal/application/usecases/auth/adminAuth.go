package auth

import (
	"GoProject1/internal/adapter/persistence"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func UC_AdmimAuth(login string, password string) (userID int, err error) {
	//Вызываем функцию запроса к БД
	passwordHash, userID, isAdmin := persistence.A_AdminAuth(login)
	//Сравниваем хешированынй пароль с настоящим
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		log.Printf("Login func | Неверный пароль для пользователя login=%v | Ошибка %v", login, err)
		return 0, err
	}

	log.Printf("Login func | Пароль для login = %v подтвержден", err)

	//Проверяем админ ли он
	if isAdmin != true {
		err := errors.New("Пользователь не является администратором")
		return 0, err
	}
	return userID, nil
}
