package admin

import "context"

// DeleteAppointment удаляет запись
func (s *AdminService) DeleteAppointment(ctx context.Context, appointmentID int) error {
	return s.repo.DeleteAppointment(ctx, appointmentID)
}
