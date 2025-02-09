package containers

import (
	"context"
	"github.com/Masterminds/squirrel"
	"pingerio/backend/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
	"pingerio/backend/internal/repository"
)

const (
	tableName   = "Containers"
	idColumn    = "id"
	nameColumn  = "name"
	imageColumn = "image"
	//ipColumn       = "ip"
	statusColumn   = "status"
	pingTimeColumn = "ping_time"
)

type repo struct {
	db *pgxpool.Pool
	sq squirrel.StatementBuilderType
}

func NewContainerRepository(db *pgxpool.Pool) repository.ContainerRepository {
	return &repo{
		db: db,
		//sb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *repo) CreateTable(ctx context.Context) error {
	query := `
	CREATE TABLE IF NOT EXISTS containers (
		id VARCHAR(255) PRIMARY KEY,
		name VARCHAR(255),
	    image VARCHAR(255),
		status VARCHAR(255),
	    created TIMESTAMP,
		ping_time TIMESTAMP
	)`

	// Выполняем SQL-запрос через пул соединений
	_, err := r.db.Exec(ctx, query)
	return err
}

func (r *repo) Store(ctx context.Context, containers []model.Container) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, c := range containers {
		query, args, err := r.sq.Insert(tableName).
			Columns(idColumn, nameColumn, imageColumn, statusColumn, pingTimeColumn).
			Values(c.ID, c.Name, c.Image, c.Status, c.PingTime).
			Suffix("ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name, image = EXCLUDED.image, status = EXCLUDED.status, ping_time = EXCLUDED.ping_time").
			ToSql()
		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx, query, args...)
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

// Получает список контейнеров из бд ->
func (r *repo) GetAll(ctx context.Context) ([]model.Container, error) {
	query, args, err := r.sq.Select(idColumn, nameColumn, imageColumn, statusColumn, pingTimeColumn).
		From(tableName).
		OrderBy("ping_time DESC").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var containers []model.Container
	for rows.Next() {
		var c model.Container
		if err := rows.Scan(&c.ID, &c.Name, &c.Image, &c.Status, &c.PingTime); err != nil {
			return nil, err
		}
		containers = append(containers, c)
	}

	return containers, nil
}
