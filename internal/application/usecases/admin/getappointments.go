package admin

import (
	"GoProject1/internal/adapter/persistence"
	"log"
)

func UC_getAppointments() []persistence.Appointment {
	apps, err := persistence.A_GetAllAppointments()
	if err != nil {
		log.Printf("UC_getAppointments() error: %v", err)
		return []persistence.Appointment{}
	}
	return apps
}
