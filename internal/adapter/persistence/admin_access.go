package persistence

import (
	"GoProject1/internal/infrastructure/db"
	"context"
)

// A_adminaccess проверяет, является ли пользователь админом (берём users.is_admin по userID).
func A_adminaccess(userID int) bool {
	var isAdmin bool

	// QueryRow - ожидаем максимум одну строку
	// Scan положит значение is_admin прямо в переменную.
	err := db.Pool.QueryRow(
		context.Background(),
		`SELECT is_admin FROM users WHERE id = $1`,
		userID,
	).Scan(&isAdmin)

	// Если ошибка  считаем, что не админ.
	if err != nil {
		return false
	}

	// Если is_admin = true, значит доступ есть
	return isAdmin
}
