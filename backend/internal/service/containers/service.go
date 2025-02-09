package containers

import (
	"context"
	"pingerio/backend/internal/model"
	"pingerio/backend/internal/repository"
	"pingerio/backend/internal/service"
)

type containerService struct {
	repo repository.ContainerRepository
}

func NewContainerService(repo repository.ContainerRepository) service.ContainerService {
	return &containerService{
		repo: repo,
	}
}

func (s *containerService) SaveContainers(ctx context.Context, containers []model.Container) error {
	return s.repo.Store(ctx, containers)
}

func (s *containerService) GetContainers(ctx context.Context) ([]model.Container, error) {
	return s.repo.GetAll(ctx)
}
