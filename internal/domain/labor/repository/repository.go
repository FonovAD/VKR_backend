package repository

import (
	"context"

	value "vkr/internal/domain/value"
)

type Repository interface {
	GetByOrganizationID(ctx context.Context, organizationID int) (value.LaborData, error)
	GetByOrganizationINN(ctx context.Context, inn string) (value.LaborData, error)
}
