package masters

import (
	"GoProject1/internal/adapter/persistence"
	"context"
)

func UC_CreateMaster(name, specialization string) {
	persistence.A_CreateMaster(name, specialization)
}

func UC_GetMasters() []map[string]interface{} {
	return persistence.A_GetMasters()
}

func UC_DeleteMaster(masterID int) {
	persistence.A_DeleteMaster(masterID)
}

// Структура, которая использует интерфейс
type MasterService struct {
	repo interface {
		GetMasters(ctx context.Context) ([]map[string]interface{}, error)
	}
}

// Конструктор (внедрение зависимости через интерфейс)
func NewMasterService(repo interface {
	GetMasters(ctx context.Context) ([]map[string]interface{}, error)
}) *MasterService {
	return &MasterService{repo: repo}
}

// Метод, который использует интерфейс (а не конкретную реализацию)
func (s *MasterService) GetAllMasters(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo.GetMasters(ctx)
}
