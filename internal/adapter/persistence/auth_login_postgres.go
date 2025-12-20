package persistence

import (
	"GoProject1/internal/infrastructure/db"
	"context"
)

// A_GetPasswordHash_User достаёт password_hash и id по логину (нужно для проверки пароля при входе).
func A_GetPasswordHash_User(ctx context.Context, login string) (string, int, error) {
	var passwordHash string
	var userID int

	// QueryRow ждём 1 строку; ошибка появится на Scan (в т.ч. если пользователя нет).
	err := db.Pool.QueryRow(
		ctx,
		`SELECT password_hash, id FROM users WHERE login = $1`,
		login,
	).Scan(&passwordHash, &userID)
	if err != nil {
		return "", 0, err
	}

	return passwordHash, userID, nil
}
