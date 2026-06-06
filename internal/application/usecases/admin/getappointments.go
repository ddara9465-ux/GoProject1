package admin

import (
	"GoProject1/internal/adapter/persistence"
	"log"

	"golang.org/x/net/context"
)

// UC_getAppointments получает список записей для админки (если ошибка — возвращаем пустой слайс).
func UC_getAppointments(ctx context.Context) []persistence.Appointment {
	apps, err := persistence.A_GetAllAppointments(ctx)
	if err != nil {
		log.Printf("UC_getAppointments error: %v", err)
		return nil // nil-слайс норм, len(nil) == 0
	}
	return apps
}
