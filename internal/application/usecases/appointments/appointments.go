package appointments

import (
	"GoProject1/internal/adapter/persistence"
	"context"
)

// UC_CreateRequestAppointments просто прокидывает создание заявки в persistence.
func UC_CreateRequestAppointments(ctx context.Context, date, employee, procedure, notes string, userID int) error {
	return persistence.A_CreateRequestAppointments(ctx, date, employee, procedure, notes, userID)
}
