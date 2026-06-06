package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB - структура, содержащая пул соединений
type DB struct {
	Pool *pgxpool.Pool
}

// NewDB создаёт новое подключение к БД
func NewDB(ctx context.Context, dsn string) (*DB, error) {
	p, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("pgxpool new: %w", err)
	}

	if err := p.Ping(ctx); err != nil {
		p.Close()
		return nil, fmt.Errorf("pgx ping: %w", err)
	}

	return &DB{Pool: p}, nil
}

// Close закрывает пул соединений
func (db *DB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}
