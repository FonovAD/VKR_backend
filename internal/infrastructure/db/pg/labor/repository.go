package labor

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	laborRepo "vkr/internal/domain/labor/repository"
	value "vkr/internal/domain/value"
	"vkr/internal/infrastructure/db/pg/model"
	"vkr/internal/logger"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewRepository(db *sqlx.DB, logger logger.Logger) laborRepo.Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) GetByOrganizationID(ctx context.Context, organizationID int) (value.LaborData, error) {
	r.logger.LogInfo(fmt.Sprintf("laborRepository - GetByOrganizationID - %d", organizationID), nil, nil)

	var laborModel model.LaborResources
	if err := r.db.GetContext(ctx, &laborModel, getLaborResourcesByOwnerQuery, organizationID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return value.LaborData{}, sql.ErrNoRows
		}
		return value.LaborData{}, err
	}

	// Вспомогательная функция для получения значения или 0.0 если nil
	getFloatValue := func(v *float64) float64 {
		if v == nil {
			return 0.0
		}
		return *v
	}

	laborData := value.LaborData{
		TotalStaffAnnual:             getFloatValue(laborModel.TotalStaffAnnual),
		ResearchStaffInternal:        getFloatValue(laborModel.ResearchStaffInternal),
		CoreOperationalStaffInternal: getFloatValue(laborModel.CoreOperationalStaffInternal),
		AdminSupportStaffInternal:    getFloatValue(laborModel.AdminSupportStaffInternal),
		ResearchStaffExternal:        getFloatValue(laborModel.ResearchStaffExternal),
		CoreOperationalStaffExternal: getFloatValue(laborModel.CoreOperationalStaffExternal),
		AdminSupportStaffExternal:    getFloatValue(laborModel.AdminSupportStaffExternal),
	}

	var fotModel model.LaborResourcesFOT
	if err := r.db.GetContext(ctx, &fotModel, getLaborFotByLaborIDQuery, laborModel.ID); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return value.LaborData{}, err
		}
	} else {
		laborData.FOT = value.LaborFOT{
			TotalStaffAnnual:             getFloatValue(fotModel.TotalStaffAnnual),
			ResearchStaffInternal:        getFloatValue(fotModel.ResearchStaffInternal),
			CoreOperationalStaffInternal: getFloatValue(fotModel.CoreOperationalStaffInternal),
			AdminSupportStaffInternal:    getFloatValue(fotModel.AdminSupportStaffInternal),
			ResearchStaffExternal:        getFloatValue(fotModel.ResearchStaffExternal),
			CoreOperationalStaffExternal: getFloatValue(fotModel.CoreOperationalStaffExternal),
			AdminSupportStaffExternal:    getFloatValue(fotModel.AdminSupportStaffExternal),
		}
	}

	return laborData, nil
}

func (r *repository) GetByOrganizationINN(ctx context.Context, inn string) (value.LaborData, error) {
	r.logger.LogInfo(fmt.Sprintf("laborRepository - GetByOrganizationINN - %s", inn), nil, nil)

	var laborModel model.LaborResources
	if err := r.db.GetContext(ctx, &laborModel, getLaborResourcesByINNQuery, inn); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return value.LaborData{}, sql.ErrNoRows
		}
		return value.LaborData{}, err
	}

	// Вспомогательная функция для получения значения или 0.0 если nil
	getFloatValue := func(v *float64) float64 {
		if v == nil {
			return 0.0
		}
		return *v
	}

	laborData := value.LaborData{
		TotalStaffAnnual:             getFloatValue(laborModel.TotalStaffAnnual),
		ResearchStaffInternal:        getFloatValue(laborModel.ResearchStaffInternal),
		CoreOperationalStaffInternal: getFloatValue(laborModel.CoreOperationalStaffInternal),
		AdminSupportStaffInternal:    getFloatValue(laborModel.AdminSupportStaffInternal),
		ResearchStaffExternal:        getFloatValue(laborModel.ResearchStaffExternal),
		CoreOperationalStaffExternal: getFloatValue(laborModel.CoreOperationalStaffExternal),
		AdminSupportStaffExternal:    getFloatValue(laborModel.AdminSupportStaffExternal),
	}

	var fotModel model.LaborResourcesFOT
	if err := r.db.GetContext(ctx, &fotModel, getLaborFotByLaborIDQuery, laborModel.ID); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return value.LaborData{}, err
		}
	} else {
		laborData.FOT = value.LaborFOT{
			TotalStaffAnnual:             getFloatValue(fotModel.TotalStaffAnnual),
			ResearchStaffInternal:        getFloatValue(fotModel.ResearchStaffInternal),
			CoreOperationalStaffInternal: getFloatValue(fotModel.CoreOperationalStaffInternal),
			AdminSupportStaffInternal:    getFloatValue(fotModel.AdminSupportStaffInternal),
			ResearchStaffExternal:        getFloatValue(fotModel.ResearchStaffExternal),
			CoreOperationalStaffExternal: getFloatValue(fotModel.CoreOperationalStaffExternal),
			AdminSupportStaffExternal:    getFloatValue(fotModel.AdminSupportStaffExternal),
		}
	}

	return laborData, nil
}
