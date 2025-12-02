package repository

import (
	"context"
	value "vkr/internal/domain/value"
)

type Repository interface {
	GetFinancialDataByOrganizationIDAndYear(ctx context.Context, organizationID int, year int16) (value.FinancialData, error)
	GetFinancialDataByOrganizationINNAndYear(ctx context.Context, inn string, year int16) (value.FinancialData, error)
}



