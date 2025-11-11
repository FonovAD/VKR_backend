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

	laborData := value.LaborData{
		TotalStaffAnnual:             laborModel.TotalStaffAnnual,
		ResearchStaffInternal:        laborModel.ResearchStaffInternal,
		CoreOperationalStaffInternal: laborModel.CoreOperationalStaffInternal,
		AdminSupportStaffInternal:    laborModel.AdminSupportStaffInternal,
		ResearchStaffExternal:        laborModel.ResearchStaffExternal,
		CoreOperationalStaffExternal: laborModel.CoreOperationalStaffExternal,
		AdminSupportStaffExternal:    laborModel.AdminSupportStaffExternal,
	}

	var fotModel model.LaborResourcesFOT
	if err := r.db.GetContext(ctx, &fotModel, getLaborFotByLaborIDQuery, laborModel.ID); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return value.LaborData{}, err
		}
	} else {
		laborData.FOT = value.LaborFOT{
			TotalStaffAnnual:             fotModel.TotalStaffAnnual,
			ResearchStaffInternal:        fotModel.ResearchStaffInternal,
			CoreOperationalStaffInternal: fotModel.CoreOperationalStaffInternal,
			AdminSupportStaffInternal:    fotModel.AdminSupportStaffInternal,
			ResearchStaffExternal:        fotModel.ResearchStaffExternal,
			CoreOperationalStaffExternal: fotModel.CoreOperationalStaffExternal,
			AdminSupportStaffExternal:    fotModel.AdminSupportStaffExternal,
		}
	}

	return laborData, nil
}

func (r *repository) GetByOrganizationINN(ctx context.Context, inn string) (value.LaborData, error) {
	var org struct {
		ID int `db:"id"`
	}

	if err := r.db.GetContext(ctx, &org, getOrganizationByINNQuery, inn); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return value.LaborData{}, sql.ErrNoRows
		}
		return value.LaborData{}, err
	}

	return r.GetByOrganizationID(ctx, org.ID)
}
