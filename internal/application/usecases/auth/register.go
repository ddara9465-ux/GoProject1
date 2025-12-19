package auth

import "GoProject1/internal/adapter/persistence"

func Register(password string, first_name string, last_name string, phone string, email string) bool {
	if persistence.RegisterTransaction(password, first_name, last_name, phone, email) {
		return true
	}
	return false
}
