package persistence

import (
	"GoProject1/internal/infrastructure/db"
	"context"
)

func GetPasswordHash_User(ctx context.Context, login string) (string, error) {
	var truePassword string
	err := db.Pool.QueryRow(ctx,
		"SELECT password_hash FROM users WHERE login=$1",
		login,
	).Scan(&truePassword)
	if err != nil {
		return "", err
	}
	return truePassword, nil
}
