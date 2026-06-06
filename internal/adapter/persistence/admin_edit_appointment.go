package persistence

import (
	"context"
)

// UpdateAppointmentDetails обновляет дату, мастера и услугу
func (r *Repository) UpdateAppointmentDetails(ctx context.Context, appointmentID int, date, master, service string) error {
	query := `UPDATE appointments
              SET appointment_date = $1,
                  employee = $2,
                  procedure = $3
              WHERE id = $4`

	_, err := r.DB.Pool.Exec(ctx, query, date, master, service, appointmentID)
	return err
}
