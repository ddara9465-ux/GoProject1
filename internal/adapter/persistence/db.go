package persistence

import (
	"GoProject1/internal/infrastructure/db"
)

// Repository - базовая структура для всех репозиториев
type Repository struct {
	DB *db.DB
}

// NewRepository создаёт новый репозиторий
func NewRepository(db *db.DB) *Repository {
	return &Repository{DB: db}
}
