package persistence

import (
	"GoProject1/internal/infrastructure/db"
	"context"
	"fmt"
)

func P_DeleteAppointment(AppointmentID int) error {
	// Exec выполняет SQL команду без возврата строк
	cmd, err := db.Pool.Exec(
		context.Background(),
		`DELETE FROM appointments WHERE id = $1`,
		AppointmentID,
	)

	// Если ошибка при выполнении запроса
	if err != nil {
		return fmt.Errorf("ошибка удаления записи: %w", err)
	}

	// Проверяем, была ли удалена хотя бы одна запись
	rowsAffected := cmd.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("запись с ID %d не найдена", AppointmentID)
	}

	// Успешно удалено
	return nil
}
