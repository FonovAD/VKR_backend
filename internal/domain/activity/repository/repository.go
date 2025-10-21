package repository

import (
	"context"
	"vkr/internal/domain/entity"
)

type ActivityRepository interface {
	Create(ctx context.Context, activity *entity.Activity) error
	GetByID(ctx context.Context, id entity.ActivityId) (*entity.Activity, error)
	Update(ctx context.Context, activity *entity.Activity) error
	Delete(ctx context.Context, id entity.ActivityId) error
	List(ctx context.Context) ([]*entity.Activity, error)
	FindByLocationType(ctx context.Context, locationType entity.LocationType) ([]*entity.Activity, error)
}
