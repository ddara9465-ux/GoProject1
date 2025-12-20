package persistence

import (
	"GoProject1/internal/infrastructure/db"
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

func A_GetAllAppointments() ([]Appointment, error) {
	rows, err := db.Pool.Query(context.Background(), `
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
