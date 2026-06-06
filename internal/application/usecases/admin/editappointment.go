package admin

import "context"

// UpdateAppointmentDetails обновляет запись
func (s *AdminService) UpdateAppointmentDetails(ctx context.Context, appointmentID int, date, master, service string) error {
	return s.repo.UpdateAppointmentDetails(ctx, appointmentID, date, master, service)
}
