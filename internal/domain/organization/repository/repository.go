package organization

import (
	"context"
	"vkr/internal/domain/entity"
)

type Repository interface {
	Create(ctx context.Context, org *entity.Organization) error
	GetByID(ctx context.Context, id entity.OrganizationID) (*entity.Organization, error)
	Update(ctx context.Context, org *entity.Organization) error
	Delete(ctx context.Context, id entity.OrganizationID) error
	FindByINN(ctx context.Context, inn string) (*entity.Organization, error)
	List(ctx context.Context) ([]*entity.Organization, error)
}
