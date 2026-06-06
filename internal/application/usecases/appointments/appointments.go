package appointments

import (
	"GoProject1/internal/adapter/persistence"
	"context"
)

type AppointmentService struct {
	repo *persistence.Repository
}

// NewAppointmentService создаёт новый сервис записей
func NewAppointmentService(repo *persistence.Repository) *AppointmentService {
	return &AppointmentService{repo: repo}
}

// CreateRequestAppointments создаёт заявку
func (s *AppointmentService) CreateRequestAppointments(ctx context.Context, date, employee, procedure, notes string, userID int) error {
	return s.repo.CreateRequestAppointments(ctx, date, employee, procedure, notes, userID)
}
