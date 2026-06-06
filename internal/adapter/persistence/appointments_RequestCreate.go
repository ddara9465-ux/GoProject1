package persistence

import (
	"context"
)

// GetClientIDByUserID ищет id клиента по user_id
func (r *Repository) GetClientIDByUserID(ctx context.Context, userID int) (int, error) {
	var clientID int
	err := r.DB.Pool.QueryRow(
		ctx,
		`SELECT id FROM clients WHERE user_id = $1`,
		userID,
	).Scan(&clientID)

	if err != nil {
		return 0, err
	}
	return clientID, nil
}

// CreateRequestAppointments создаёт запись в appointments со статусом "Запрос звонка"
func (r *Repository) CreateRequestAppointments(ctx context.Context, date, employee, procedure, notes string, userID int) error {
	clientID, err := r.GetClientIDByUserID(ctx, userID)
	if err != nil {
		return err
	}

	_, err = r.DB.Pool.Exec(ctx,
		`INSERT INTO appointments (client_id, procedure, employee, appointment_date, status, notes)
         VALUES ($1, $2, $3, $4, $5, $6)`,
		clientID, procedure, employee, date, "Запрос звонка", notes)

	return err
}
