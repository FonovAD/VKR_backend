package repository

import (
	"context"
	"vkr/internal/domain/entity"
)

type MetricRepository interface {
	Create(ctx context.Context, metric *entity.Metric) error
	GetByID(ctx context.Context, id int) (*entity.Metric, error)
	Update(ctx context.Context, metric *entity.Metric) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]*entity.Metric, error)
	FindByAudienceType(ctx context.Context, audienceType entity.AudienceType) ([]*entity.Metric, error)
}
