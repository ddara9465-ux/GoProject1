package auth

import (
	"GoProject1/internal/adapter/persistence"
	"context"
	"log"
)

func UC_Login(login string, password string) bool {
	trueHashPassword, err := persistence.A_GetPasswordHash_User(context.Background(), login)
	if err != nil {
		log.Printf("Log err: %v", err)
	} else {
		if password == trueHashPassword {
			log.Print("Login func | Данные авторизации подтверждены")
			return true
		} else {
			log.Print("Login func | Неверные данные авторизации")
		}
	}
	return false
}
