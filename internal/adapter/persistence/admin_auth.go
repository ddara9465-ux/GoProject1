package persistence

import (
	"GoProject1/internal/infrastructure/db"
	"context"
	"log"
)

func A_AdminAuth(login string) (passwordHash string, userID int, isAdmin bool) {
	query := `SELECT id, password_hash, is_admin FROM users WHERE login = $1`
	err := db.Pool.QueryRow(context.Background(), query, login).Scan(&userID, &passwordHash, &isAdmin)
	if err != nil {
		log.Printf("Error persistence A_AdminAuth: %v", err)
		return "0", 0, false
	}
	return passwordHash, userID, isAdmin
}
