package persistence

import (
	"context"
	"fmt"
)

// UpdateAppointmentStatus обновляет статус записи appointments по её id
func (r *Repository) UpdateAppointmentStatus(ctx context.Context, id int, status string) error {
	cmd, err := r.DB.Pool.Exec(ctx, `
        UPDATE appointments
        SET status = $1
        WHERE id = $2
    `, status, id)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("appointment id=%d not found", id)
	}

	return nil
}
