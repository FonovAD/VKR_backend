package repository

import (
	"context"
	"vkr/internal/domain/entity"
)

type ListParams struct {
	Name       *string // Filter by name (partial match, case-insensitive)
	MuseumType *string // Filter by museum type (e.g., "art", "memorial_reserve", "estate", etc.)
	Limit      int     // Number of items to return
	Offset     int     // Number of items to skip
}

type ListResult struct {
	Museums    []*entity.Museum
	TotalCount int64
}

type Repository interface {
	Create(ctx context.Context, museum *entity.Museum) error
	GetByID(ctx context.Context, id entity.MuseumID) (*entity.Museum, error)
	Update(ctx context.Context, museum *entity.Museum) error
	Delete(ctx context.Context, id entity.MuseumID) error
	FindByINN(ctx context.Context, inn string) (*entity.Museum, error)
	FindByOwner(ctx context.Context, ownerID entity.OrganizationID) ([]*entity.Museum, error)
	List(ctx context.Context, params ListParams) (*ListResult, error)
}
