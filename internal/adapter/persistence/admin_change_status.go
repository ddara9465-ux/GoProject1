package persistence

import (
	"GoProject1/internal/infrastructure/db"
	"context"
	"fmt"
)

// A_UpdateAppointmentStatus обновляет статус записи appointments по её id.
func A_UpdateAppointmentStatus(ctx context.Context, id int, status string) error {
	// Exec, нам не нужно читать строки из результата — только выполнить UPDATE.
	cmd, err := db.Pool.Exec(ctx, `
		UPDATE appointments
		SET status = $1
		WHERE id = $2
	`, status, id)
	if err != nil {
		return err
	}

	// Если RowsAffected() == 0, значит WHERE id = $2 не нашёл запись.
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("appointment id=%d not found", id)
	}

	return nil
}
