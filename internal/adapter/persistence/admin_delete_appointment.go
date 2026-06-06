package persistence

import (
	"context"
	"fmt"
)

// DeleteAppointment удаляет запись по ID
func (r *Repository) DeleteAppointment(ctx context.Context, appointmentID int) error {
	cmd, err := r.DB.Pool.Exec(
		ctx,
		`DELETE FROM appointments WHERE id = $1`,
		appointmentID,
	)

	if err != nil {
		return fmt.Errorf("ошибка удаления записи: %w", err)
	}

	rowsAffected := cmd.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("запись с ID %d не найдена", appointmentID)
	}

	return nil
}
