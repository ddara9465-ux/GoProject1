package auth

import "GoProject1/internal/adapter/persistence"

func UC_Register(password string, first_name string, last_name string, phone string, email string) (int, bool) {
	userID, err := persistence.A_RegisterTransaction(password, first_name, last_name, phone, email)
	if err {
		return 0, true // если ошибка = true не получаем userID
	}
	return userID, false
}
