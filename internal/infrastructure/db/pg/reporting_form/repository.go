package reporting_form

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	reportingFormRepo "vkr/internal/domain/reporting_form/repository"
	value "vkr/internal/domain/value"
	"vkr/internal/infrastructure/db/pg/model"
	"vkr/internal/logger"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewRepository(db *sqlx.DB, logger logger.Logger) reportingFormRepo.Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) GetFinancialDataByOrganizationIDAndYear(ctx context.Context, organizationID int, year int16) (value.FinancialData, error) {
	r.logger.LogInfo(fmt.Sprintf("reportingFormRepository - GetFinancialDataByOrganizationIDAndYear - orgID: %d, year: %d", organizationID, year), nil, nil)

	var financialModel model.MuseumRevenuesExpenses
	if err := r.db.GetContext(ctx, &financialModel, getFinancialDataByOrgIDAndYearQuery, organizationID, year); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return value.FinancialData{}, sql.ErrNoRows
		}
		return value.FinancialData{}, err
	}

	// Вспомогательная функция для получения значения или 0.0 если nil
	getFloatValue := func(v *float64) float64 {
		if v == nil {
			return 0.0
		}
		return *v
	}

	// Получаем сумму из other_funding_sources
	var fundingSum struct {
		Total float64 `db:"total"`
	}
	otherFundingSourcesSum := 0.0
	if err := r.db.GetContext(ctx, &fundingSum, getOtherFundingSourcesSumQuery, financialModel.ID); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return value.FinancialData{}, err
		}
	} else {
		otherFundingSourcesSum = fundingSum.Total
	}

	// Получаем сумму из other_expenses
	var expensesSum struct {
		Total float64 `db:"total"`
	}
	otherExpensesSum := 0.0
	if err := r.db.GetContext(ctx, &expensesSum, getOtherExpensesSumQuery, financialModel.ID); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return value.FinancialData{}, err
		}
	} else {
		otherExpensesSum = expensesSum.Total
	}

	return value.FinancialData{
		TotalRevenue:           getFloatValue(financialModel.TotalRevenue),
		StateAssignmentSubsidy: getFloatValue(financialModel.StateAssignmentSubsidy),
		EarnedRevenue:          getFloatValue(financialModel.EarnedRevenue),
		OtherFundingSources:    otherFundingSourcesSum,
		TotalExpenses:          getFloatValue(financialModel.TotalExpenses),
		InventoryAssetsAcquisition: getFloatValue(financialModel.InventoryAssetsAcquisition),
		UtilityServices:            getFloatValue(financialModel.UtilityServices),
		CommunicationServices:      getFloatValue(financialModel.CommunicationServices),
		TransportationServices:     getFloatValue(financialModel.TransportationServices),
		ValuableAssetsAcquisition:  getFloatValue(financialModel.ValuableAssetsAcquisition),
		RealEstateMaintenance:      getFloatValue(financialModel.RealEstateMaintenance),
		ValuableAssetsMaintenance:  getFloatValue(financialModel.ValuableAssetsMaintenance),
		GeneralAdministrativeCosts: getFloatValue(financialModel.GeneralAdministrativeCosts),
		TaxPayments:                getFloatValue(financialModel.TaxPayments),
		OtherExpenses:              otherExpensesSum,
	}, nil
}

func (r *repository) GetFinancialDataByOrganizationINNAndYear(ctx context.Context, inn string, year int16) (value.FinancialData, error) {
	r.logger.LogInfo(fmt.Sprintf("reportingFormRepository - GetFinancialDataByOrganizationINNAndYear - inn: %s, year: %d", inn, year), nil, nil)

	var financialModel model.MuseumRevenuesExpenses
	if err := r.db.GetContext(ctx, &financialModel, getFinancialDataByINNAndYearQuery, inn, year); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return value.FinancialData{}, sql.ErrNoRows
		}
		return value.FinancialData{}, err
	}

	// Вспомогательная функция для получения значения или 0.0 если nil
	getFloatValue := func(v *float64) float64 {
		if v == nil {
			return 0.0
		}
		return *v
	}

	// Получаем сумму из other_funding_sources
	var fundingSum struct {
		Total float64 `db:"total"`
	}
	otherFundingSourcesSum := 0.0
	if err := r.db.GetContext(ctx, &fundingSum, getOtherFundingSourcesSumQuery, financialModel.ID); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return value.FinancialData{}, err
		}
	} else {
		otherFundingSourcesSum = fundingSum.Total
	}

	// Получаем сумму из other_expenses
	var expensesSum struct {
		Total float64 `db:"total"`
	}
	otherExpensesSum := 0.0
	if err := r.db.GetContext(ctx, &expensesSum, getOtherExpensesSumQuery, financialModel.ID); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return value.FinancialData{}, err
		}
	} else {
		otherExpensesSum = expensesSum.Total
	}

	return value.FinancialData{
		TotalRevenue:           getFloatValue(financialModel.TotalRevenue),
		StateAssignmentSubsidy: getFloatValue(financialModel.StateAssignmentSubsidy),
		EarnedRevenue:          getFloatValue(financialModel.EarnedRevenue),
		OtherFundingSources:    otherFundingSourcesSum,
		TotalExpenses:          getFloatValue(financialModel.TotalExpenses),
		InventoryAssetsAcquisition: getFloatValue(financialModel.InventoryAssetsAcquisition),
		UtilityServices:            getFloatValue(financialModel.UtilityServices),
		CommunicationServices:      getFloatValue(financialModel.CommunicationServices),
		TransportationServices:     getFloatValue(financialModel.TransportationServices),
		ValuableAssetsAcquisition:  getFloatValue(financialModel.ValuableAssetsAcquisition),
		RealEstateMaintenance:      getFloatValue(financialModel.RealEstateMaintenance),
		ValuableAssetsMaintenance:  getFloatValue(financialModel.ValuableAssetsMaintenance),
		GeneralAdministrativeCosts: getFloatValue(financialModel.GeneralAdministrativeCosts),
		TaxPayments:                getFloatValue(financialModel.TaxPayments),
		OtherExpenses:              otherExpensesSum,
	}, nil
}

