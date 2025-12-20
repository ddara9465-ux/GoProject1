package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

// InitPostgresPool создаёт пул подключений и проверяет, что БД реально доступна (Ping).
func InitPostgresPool(ctx context.Context, dsn string) error {
	// Создаём пул по DSN (строка подключения).
	p, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return fmt.Errorf("pgxpool new: %w", err)
	}

	// Проверяем соединение; если не ок — закрываем пул, чтобы не было утечек.
	if err := p.Ping(ctx); err != nil {
		p.Close()
		return fmt.Errorf("pgx ping: %w", err)
	}

	// Сохраняем пул в глобальную переменную, чтобы использовать в других пакетах.
	Pool = p
	return nil
}
