package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"time"
)

// NewPostgresPool создаёт и возвращает пул подключений к PostgreSQL
func NewPostgresPool(ctx context.Context) (*pgxpool.Pool, error) {
	conn := os.Getenv("po") // Пример: "postgres://user:password@localhost:5432/dbname"
	if conn == "" {
		return nil, fmt.Errorf("POSTGRES_CONN is not set")
	}

	config, err := pgxpool.ParseConfig(conn)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database URL: %w", err)
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.HealthCheckPeriod = 30 * time.Second

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	return pool, nil
}
