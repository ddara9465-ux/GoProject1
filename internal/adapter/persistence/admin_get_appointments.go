package persistence

import (
	"context"
)

type Appointment struct {
	ID              int
	ClientID        int
	FirstName       string
	Phone           string
	Procedure       string
	Employee        string
	AppointmentDate string
	Status          string
	Notes           string
}

// GetAllAppointments достаёт все записи + данные клиента через JOIN
func (r *Repository) GetAllAppointments(ctx context.Context) ([]Appointment, error) {
	rows, err := r.DB.Pool.Query(ctx, `
        SELECT
            a.id,
            a.client_id,
            c.first_name,
            c.phone,
            a.procedure,
            a.employee,
            a.appointment_date,
            a.status,
            a.notes
        FROM appointments a
        INNER JOIN clients c ON c.id = a.client_id
        ORDER BY a.id DESC
    `)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	apps := make([]Appointment, 0)
	for rows.Next() {
		var a Appointment
		if err := rows.Scan(
			&a.ID,
			&a.ClientID,
			&a.FirstName,
			&a.Phone,
			&a.Procedure,
			&a.Employee,
			&a.AppointmentDate,
			&a.Status,
			&a.Notes,
		); err != nil {
			return nil, err
		}
		apps = append(apps, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return apps, nil
}
