package admin

import "GoProject1/internal/adapter/persistence"

// UC_UpdateAppointmentDetails обновляет запись
func UC_UpdateAppointmentDetails(appointmentID int, date, master, service string) error {
	err := persistence.A_UpdateAppointmentDetails(appointmentID, date, master, service)
	return err
}
