package persistence

import (
	"GoProject1/internal/infrastructure/db"
	"context"
)

// GetClientIDByUserID ищет id клиента по user_id.
func GetClientIDByUserID(ctx context.Context, userID int) (int, error) {
	var clientID int
	err := db.Pool.QueryRow(
		ctx,
		`SELECT id FROM clients WHERE user_id = $1`,
		userID,
	).Scan(&clientID)
	if err != nil {
		return 0, err
	}
	return clientID, nil
}

// A_CreateRequestAppointments создаёт запись в appointments со статусом "Запрос звонка".
func A_CreateRequestAppointments(ctx context.Context, date, employee, procedure, notes string, userID int) error {
	clientID, err := GetClientIDByUserID(ctx, userID)
	if err != nil {
		return err
	}

	_, err = db.Pool.Exec(ctx,
		`INSERT INTO appointments (client_id, procedure, employee, appointment_date, status, notes)
         VALUES ($1, $2, $3, $4, $5, $6)`,
		clientID, procedure, employee, date, "Запрос звонка", notes)

	return err
}
