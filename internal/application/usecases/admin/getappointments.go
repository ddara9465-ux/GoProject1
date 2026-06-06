package admin

import (
	"GoProject1/internal/adapter/persistence"
	"context"
	"log"
)

type AdminService struct {
	repo *persistence.Repository
}

// NewAdminService создаёт новый сервис администратора
func NewAdminService(repo *persistence.Repository) *AdminService {
	return &AdminService{repo: repo}
}

// GetAppointments получает список записей для админки
func (s *AdminService) GetAppointments(ctx context.Context) []persistence.Appointment {
	apps, err := s.repo.GetAllAppointments(ctx)
	if err != nil {
		log.Printf("GetAppointments error: %v", err)
		return nil
	}
	return apps
}
