package persistence

import (
	"GoProject1/internal/infrastructure/db"
	"context"
	"log"
)

func A_adminaccess(userID int) bool {
	var is_admin bool
	err := db.Pool.QueryRow(context.Background(),
		`SELECT is_admin FROM users WHERE id = $1`,
		userID,
	).Scan(&is_admin)
	log.Print(is_admin)
	if err != nil {
		return false
	}
	if is_admin {
		return true // ошибки в чтении нет, права подтверждены
	}
	return false
}
