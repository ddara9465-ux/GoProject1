package masters

import (
	"GoProject1/internal/adapter/persistence"
	"context"
)

type MasterService struct {
	repo *persistence.Repository
}

// NewMasterService создаёт новый сервис для работы с мастерами
func NewMasterService(repo *persistence.Repository) *MasterService {
	return &MasterService{repo: repo}
}

func (s *MasterService) CreateMaster(ctx context.Context, name, specialization string) error {
	return s.repo.CreateMaster(ctx, name, specialization)
}

func (s *MasterService) GetMasters(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo.GetMasters(ctx)
}

func (s *MasterService) DeleteMaster(ctx context.Context, masterID int) error {
	return s.repo.DeleteMaster(ctx, masterID)
}
