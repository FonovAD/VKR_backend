package repository

import (
	"context"
	"vkr/internal/domain/entity"
)

type MuseumRepository interface {
	Create(ctx context.Context, museum *entity.Museum) error
	GetByID(ctx context.Context, id entity.MuseumID) (*entity.Museum, error)
	Update(ctx context.Context, museum *entity.Museum) error
	Delete(ctx context.Context, id entity.MuseumID) error
	FindByINN(ctx context.Context, inn string) (*entity.Museum, error)
	FindByOwner(ctx context.Context, ownerID entity.OrganizationID) ([]*entity.Museum, error)
	List(ctx context.Context) ([]*entity.Museum, error)
}
