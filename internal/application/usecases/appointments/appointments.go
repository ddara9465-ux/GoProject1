package appointments

import "GoProject1/internal/adapter/persistence"

func UC_CreateRequestAppointments(date string, employee string, procedure string, notes string, userID int) {
	persistence.A_CreateRequestAppointments(date, employee, procedure, notes, userID)
}
