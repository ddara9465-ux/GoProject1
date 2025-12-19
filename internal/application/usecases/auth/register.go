package auth

import "GoProject1/internal/adapter/persistence"

func UC_Register(password string, first_name string, last_name string, phone string, email string) bool {
	if persistence.A_RegisterTransaction(password, first_name, last_name, phone, email) {
		return true
	}
	return false
}
