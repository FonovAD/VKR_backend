package repository

import (
	"context"
	"vkr/internal/domain/entity"
)

type Repository interface {
	Create(ctx context.Context, activity *entity.Activity) error
	GetByINN(ctx context.Context, inn string) ([]entity.Activity, error)
	Update(ctx context.Context, activity *entity.Activity) error
	Delete(ctx context.Context, inn string, activityTypeID int, visitorCategory entity.VisitorCategory, year int16) error
	List(ctx context.Context) ([]entity.Activity, error)
}
