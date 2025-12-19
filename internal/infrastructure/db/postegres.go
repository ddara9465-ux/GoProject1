package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitPostgresPool(ctx context.Context, dsn string) error {
	p, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return fmt.Errorf("pgxpool new: %w", err)
	}
	if err := p.Ping(ctx); err != nil {
		p.Close()
		return fmt.Errorf("pgx ping: %w", err)
	}
	Pool = p
	return nil
}
