package persistence

import (
	"GoProject1/internal/infrastructure/db"
	"context"
	"fmt"
)

// CheckAdminAuth проверяет логин/пароль админа
func CheckAdminAuth(username, password string) (bool, int, error) {
	var adminID int
	var storedPassword string

	// Простая проверка (можно добавить bcrypt позже)
	query := `SELECT id, password_hash FROM admins_login 
	          WHERE username = $1 AND is_active = true`

	err := db.Pool.QueryRow(context.Background(), query, username).Scan(
		&adminID, &storedPassword,
	)

	if err != nil {
		return false, 0, fmt.Errorf("админ не найден")
	}

	// Простая проверка пароля (лучше использовать bcrypt)
	if password != storedPassword {
		return false, 0, fmt.Errorf("неверный пароль")
	}

	return true, adminID, nil
}
