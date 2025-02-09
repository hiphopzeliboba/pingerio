package repository

import (
	"context"
	"pingerio/backend/internal/model"
)

type ContainerRepository interface {
	Store(ctx context.Context, containers []model.Container) error
	GetAll(ctx context.Context) ([]model.Container, error)
	CreateTable(ctx context.Context) error
}
