package persistence

import (
	"context"
	"log"
)

// AdminAuth проверка администратора по логину
func (r *Repository) AdminAuth(ctx context.Context, login string) (passwordHash string, userID int, isAdmin bool) {
	query := `SELECT id, password_hash, is_admin FROM users WHERE login = $1`

	err := r.DB.Pool.QueryRow(ctx, query, login).Scan(&userID, &passwordHash, &isAdmin)
	if err != nil {
		log.Printf("Error persistence AdminAuth: %v", err)
		return "", 0, false
	}

	return passwordHash, userID, isAdmin
}
