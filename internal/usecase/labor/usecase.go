package labor

import (
	"context"

	laborRepo "vkr/internal/domain/labor/repository"
	laborService "vkr/internal/domain/labor/service"
	value "vkr/internal/domain/value"
)

type UseCase interface {
	GetByOrganizationID(ctx context.Context, organizationID int) (value.LaborData, error)
	GetByOrganizationINN(ctx context.Context, inn string) (value.LaborData, error)
}

type useCase struct {
	repo    laborRepo.Repository
	service laborService.Service
}

func NewUseCase(repo laborRepo.Repository, service laborService.Service) UseCase {
	return &useCase{
		repo:    repo,
		service: service,
	}
}

func (u *useCase) GetByOrganizationID(ctx context.Context, organizationID int) (value.LaborData, error) {
	return u.repo.GetByOrganizationID(ctx, organizationID)
}

func (u *useCase) GetByOrganizationINN(ctx context.Context, inn string) (value.LaborData, error) {
	return u.repo.GetByOrganizationINN(ctx, inn)
}
