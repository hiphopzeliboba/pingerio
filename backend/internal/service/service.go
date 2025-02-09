package service

import (
	"context"
	"pingerio/backend/internal/model"
)

type ContainerService interface {
	SaveContainers(ctx context.Context, containers []model.Container) error
	GetContainers(ctx context.Context) ([]model.Container, error)
}
