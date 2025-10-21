package repository

import (
	"context"

	"vkr/internal/domain/entity"
)

type ReportingFormRepository interface {
	Create(ctx context.Context, form *entity.ReportingForm) error
	GetByID(ctx context.Context, id entity.FormID) (*entity.ReportingForm, error)
	GetByOrganizationAndYear(ctx context.Context, orgID entity.OrganizationID, year int16) (*entity.ReportingForm, error)
	Update(ctx context.Context, form *entity.ReportingForm) error
	Delete(ctx context.Context, id entity.FormID) error
	ListByOrganization(ctx context.Context, orgID entity.OrganizationID) ([]*entity.ReportingForm, error)
	ListByYear(ctx context.Context, year int16) ([]*entity.ReportingForm, error)
	List(ctx context.Context) ([]*entity.ReportingForm, error)
}
