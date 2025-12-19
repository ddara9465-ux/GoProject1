package persistence

import (
	"GoProject1/internal/infrastructure/db"
	"context"
)

func A_GetPasswordHash_User(ctx context.Context, login string) (string, int, error) {
	var truePassword string
	var userID int
	err := db.Pool.QueryRow(ctx,
		"SELECT password_hash, id FROM users WHERE login=$1",
		login,
	).Scan(&truePassword, &userID)
	if err != nil {
		return "", 0, err
	}
	return truePassword, userID, nil
}
