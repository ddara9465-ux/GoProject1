package persistence

import (
	"GoProject1/internal/infrastructure/db"
	"context"
)

// UpdateAppointmentDetails обновляет дату, мастера и услугу
func A_UpdateAppointmentDetails(appointmentID int, date, master, service string) error {
	query := `UPDATE appointments 
	          SET appointment_date = $1, 
	              employee = $2, 
	              procedure = $3 
	          WHERE id = $4`

	_, err := db.Pool.Exec(
		context.Background(),
		query, date, master, service, appointmentID,
	)

	return err

}
