package admin

import "context"

// UpdateAppointmentStatus обновляет статус записи
func (s *AdminService) UpdateAppointmentStatus(ctx context.Context, id int, status string) error {
	return s.repo.UpdateAppointmentStatus(ctx, id, status)
}
