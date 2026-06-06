package persistence

import (
	"context"
)

// GetPasswordHashByLogin достаёт password_hash и id по логину
func (r *Repository) GetPasswordHashByLogin(ctx context.Context, login string) (string, int, error) {
	var passwordHash string
	var userID int

	err := r.DB.Pool.QueryRow(
		ctx,
		`SELECT password_hash, id FROM users WHERE login = $1`,
		login,
	).Scan(&passwordHash, &userID)

	if err != nil {
		return "", 0, err
	}

	return passwordHash, userID, nil
}
