package appointments

import "GoProject1/internal/adapter/persistence"

// UC_CreateRequestAppointments просто прокидывает создание заявки в persistence.
func UC_CreateRequestAppointments(date, employee, procedure, notes string, userID int) {
	persistence.A_CreateRequestAppointments(date, employee, procedure, notes, userID)
}
