package admin

import (
	"GoProject1/internal/adapter/persistence"
)

func UC_DeleteAppointment(AppointmentID int) error {
	err := persistence.P_DeleteAppointment(AppointmentID)
	return err
}
