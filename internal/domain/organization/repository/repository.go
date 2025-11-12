package organization

import (
	"context"
	"vkr/internal/domain/entity"
)

type ListParams struct {
	Name     *string // Filter by name (partial match, case-insensitive)
	Limit    int     // Number of items to return
	Offset   int     // Number of items to skip
}

type ListResult struct {
	Organizations []*entity.Organization
	TotalCount    int64
}

type Repository interface {
	Create(ctx context.Context, org *entity.Organization) error
	GetByID(ctx context.Context, id entity.OrganizationID) (*entity.Organization, error)
	Update(ctx context.Context, org *entity.Organization) error
	Delete(ctx context.Context, id entity.OrganizationID) error
	FindByINN(ctx context.Context, inn string) (*entity.Organization, error)
	List(ctx context.Context, params ListParams) (*ListResult, error)
}
