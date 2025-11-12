package organisation

import (
	"context"

	"vkr/internal/domain/entity"
	orgRepo "vkr/internal/domain/organization/repository"
	orgService "vkr/internal/domain/organization/service"
)

type UseCase interface {
	Create(ctx context.Context, inn, name string, existMuseum bool) (*entity.Organization, error)
	GetByID(ctx context.Context, id entity.OrganizationID) (*entity.Organization, error)
	Update(ctx context.Context, id entity.OrganizationID, inn, name string, existMuseum bool) (*entity.Organization, error)
	Delete(ctx context.Context, id entity.OrganizationID) error
	FindByINN(ctx context.Context, inn string) (*entity.Organization, error)
	List(ctx context.Context, name *string, limit, offset int) (*orgRepo.ListResult, error)
}

type organizationUseCase struct {
	orgRepo    orgRepo.Repository
	orgService orgService.Service
}

func NewOrganizationUseCase(orgRepo orgRepo.Repository, orgService orgService.Service) UseCase {
	return &organizationUseCase{
		orgRepo:    orgRepo,
		orgService: orgService,
	}
}

func (u *organizationUseCase) Create(
	ctx context.Context,
	inn, name string,
	existMuseum bool,
) (*entity.Organization, error) {
	org := &entity.Organization{
		INN:         inn,
		Name:        name,
		ExistMuseum: existMuseum,
	}

	if err := u.orgRepo.Create(ctx, org); err != nil {
		return nil, err
	}

	return org, nil
}

func (u *organizationUseCase) GetByID(ctx context.Context, id entity.OrganizationID) (*entity.Organization, error) {
	return u.orgRepo.GetByID(ctx, id)
}

func (u *organizationUseCase) Update(
	ctx context.Context,
	id entity.OrganizationID,
	inn, name string,
	existMuseum bool,
) (*entity.Organization, error) {
	org := &entity.Organization{
		ID:          id,
		INN:         inn,
		Name:        name,
		ExistMuseum: existMuseum,
	}

	if err := u.orgRepo.Update(ctx, org); err != nil {
		return nil, err
	}

	return org, nil
}

func (u *organizationUseCase) Delete(ctx context.Context, id entity.OrganizationID) error {
	return u.orgRepo.Delete(ctx, id)
}

func (u *organizationUseCase) FindByINN(ctx context.Context, inn string) (*entity.Organization, error) {
	return u.orgRepo.FindByINN(ctx, inn)
}

func (u *organizationUseCase) List(ctx context.Context, name *string, limit, offset int) (*orgRepo.ListResult, error) {
	return u.orgRepo.List(ctx, orgRepo.ListParams{
		Name:   name,
		Limit:  limit,
		Offset: offset,
	})
}
