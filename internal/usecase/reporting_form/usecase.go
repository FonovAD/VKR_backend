package reportingform

import (
	"context"

	"vkr/internal/domain/entity"
	reportingRepo "vkr/internal/domain/reporting_form/repository"
	reportingService "vkr/internal/domain/reporting_form/service"
)

type UseCase interface {
	Create(ctx context.Context, form *entity.ReportingForm) (*entity.ReportingForm, error)
	GetByID(ctx context.Context, id entity.FormID) (*entity.ReportingForm, error)
	GetByOrganizationAndYear(ctx context.Context, orgID entity.OrganizationID, year int16) (*entity.ReportingForm, error)
	Update(ctx context.Context, form *entity.ReportingForm) (*entity.ReportingForm, error)
	Delete(ctx context.Context, id entity.FormID) error
	ListByOrganization(ctx context.Context, orgID entity.OrganizationID) ([]*entity.ReportingForm, error)
	ListByYear(ctx context.Context, year int16) ([]*entity.ReportingForm, error)
	List(ctx context.Context) ([]*entity.ReportingForm, error)
}

type useCase struct {
	repo    reportingRepo.Repository
	service *reportingService.Service
}

func NewUseCase(repo reportingRepo.Repository, service *reportingService.Service) UseCase {
	return &useCase{
		repo:    repo,
		service: service,
	}
}

func (u *useCase) Create(ctx context.Context, form *entity.ReportingForm) (*entity.ReportingForm, error) {
	if err := u.repo.Create(ctx, form); err != nil {
		return nil, err
	}
	return form, nil
}

func (u *useCase) GetByID(ctx context.Context, id entity.FormID) (*entity.ReportingForm, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *useCase) GetByOrganizationAndYear(ctx context.Context, orgID entity.OrganizationID, year int16) (*entity.ReportingForm, error) {
	return u.repo.GetByOrganizationAndYear(ctx, orgID, year)
}

func (u *useCase) Update(ctx context.Context, form *entity.ReportingForm) (*entity.ReportingForm, error) {
	if err := u.repo.Update(ctx, form); err != nil {
		return nil, err
	}
	return form, nil
}

func (u *useCase) Delete(ctx context.Context, id entity.FormID) error {
	return u.repo.Delete(ctx, id)
}

func (u *useCase) ListByOrganization(ctx context.Context, orgID entity.OrganizationID) ([]*entity.ReportingForm, error) {
	return u.repo.ListByOrganization(ctx, orgID)
}

func (u *useCase) ListByYear(ctx context.Context, year int16) ([]*entity.ReportingForm, error) {
	return u.repo.ListByYear(ctx, year)
}

func (u *useCase) List(ctx context.Context) ([]*entity.ReportingForm, error) {
	return u.repo.List(ctx)
}
