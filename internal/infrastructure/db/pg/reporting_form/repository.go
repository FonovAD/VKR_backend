package reportingform

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"vkr/internal/domain/entity"
	reportingRepo "vkr/internal/domain/reporting_form/repository"
	value "vkr/internal/domain/value"
	"vkr/internal/infrastructure/db/pg/model"
	"vkr/internal/logger"

	"github.com/jmoiron/sqlx"
)

type reportingFormRepository struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewReportingFormRepository(db *sqlx.DB, logger logger.Logger) reportingRepo.Repository {
	return &reportingFormRepository{
		db:     db,
		logger: logger,
	}
}

func (r *reportingFormRepository) Create(ctx context.Context, form *entity.ReportingForm) error {
	r.logger.LogInfo(fmt.Sprintf("reportingFormRepository - Create - OrgID: %d, Year: %d", form.OrganizationID, form.Year), nil, nil)

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	if err = r.syncFormDetails(ctx, tx, form); err != nil {
		return err
	}

	var id int
	if err = tx.QueryRowxContext(ctx, createReportingFormQuery, form.OrganizationID, form.Year, string(form.Status)).Scan(&id); err != nil {
		return err
	}

	form.ID = entity.FormID(id)

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *reportingFormRepository) GetByID(ctx context.Context, id entity.FormID) (*entity.ReportingForm, error) {
	r.logger.LogInfo(fmt.Sprintf("reportingFormRepository - GetByID - %d", id), nil, nil)

	var dbModel model.ReportingForm
	err := r.db.GetContext(ctx, &dbModel, getReportingFormByIDQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return r.buildEntity(ctx, &dbModel)
}

func (r *reportingFormRepository) GetByOrganizationAndYear(ctx context.Context, orgID entity.OrganizationID, year int16) (*entity.ReportingForm, error) {
	r.logger.LogInfo(fmt.Sprintf("reportingFormRepository - GetByOrganizationAndYear - OrgID: %d, Year: %d", orgID, year), nil, nil)

	var dbModel model.ReportingForm
	err := r.db.GetContext(ctx, &dbModel, getReportingFormByOrganizationAndYearQuery, orgID, year)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return r.buildEntity(ctx, &dbModel)
}

func (r *reportingFormRepository) Update(ctx context.Context, form *entity.ReportingForm) error {
	r.logger.LogInfo(fmt.Sprintf("reportingFormRepository - Update - ID: %d", form.ID), nil, nil)

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	if err = r.syncFormDetails(ctx, tx, form); err != nil {
		return err
	}

	updateParams := map[string]interface{}{
		"id":     form.ID,
		"year":   form.Year,
		"status": string(form.Status),
	}

	if _, err = tx.NamedExecContext(ctx, updateReportingFormQuery, updateParams); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *reportingFormRepository) Delete(ctx context.Context, id entity.FormID) error {
	r.logger.LogInfo(fmt.Sprintf("reportingFormRepository - Delete - ID: %d", id), nil, nil)
	_, err := r.db.ExecContext(ctx, deleteReportingFormQuery, id)
	return err
}

func (r *reportingFormRepository) ListByOrganization(ctx context.Context, orgID entity.OrganizationID) ([]*entity.ReportingForm, error) {
	r.logger.LogInfo(fmt.Sprintf("reportingFormRepository - ListByOrganization - OrgID: %d", orgID), nil, nil)

	var dbModels []model.ReportingForm
	if err := r.db.SelectContext(ctx, &dbModels, listReportingFormsByOrganizationQuery, orgID); err != nil {
		return nil, err
	}

	return r.buildEntities(ctx, dbModels)
}

func (r *reportingFormRepository) ListByYear(ctx context.Context, year int16) ([]*entity.ReportingForm, error) {
	r.logger.LogInfo(fmt.Sprintf("reportingFormRepository - ListByYear - Year: %d", year), nil, nil)

	var dbModels []model.ReportingForm
	if err := r.db.SelectContext(ctx, &dbModels, listReportingFormsByYearQuery, year); err != nil {
		return nil, err
	}

	return r.buildEntities(ctx, dbModels)
}

func (r *reportingFormRepository) List(ctx context.Context) ([]*entity.ReportingForm, error) {
	r.logger.LogInfo("reportingFormRepository - List", nil, nil)

	var dbModels []model.ReportingForm
	if err := r.db.SelectContext(ctx, &dbModels, listReportingFormsQuery); err != nil {
		return nil, err
	}

	return r.buildEntities(ctx, dbModels)
}

func (r *reportingFormRepository) buildEntities(ctx context.Context, dbModels []model.ReportingForm) ([]*entity.ReportingForm, error) {
	forms := make([]*entity.ReportingForm, 0, len(dbModels))
	for i := range dbModels {
		form, err := r.buildEntity(ctx, &dbModels[i])
		if err != nil {
			return nil, err
		}
		forms = append(forms, form)
	}

	return forms, nil
}

func (r *reportingFormRepository) buildEntity(ctx context.Context, dbModel *model.ReportingForm) (*entity.ReportingForm, error) {
	generalInfo, err := r.fetchGeneralInfo(ctx, dbModel.OrganizationID)
	if err != nil {
		return nil, err
	}

	museumDetails, err := r.fetchMuseumDetails(ctx, dbModel.OrganizationID)
	if err != nil {
		return nil, err
	}

	laborData, err := r.fetchLaborData(ctx, dbModel.OrganizationID)
	if err != nil {
		return nil, err
	}

	financialData, err := r.fetchFinancialData(ctx, dbModel.OrganizationID, dbModel.Year)
	if err != nil {
		return nil, err
	}

	return &entity.ReportingForm{
		ID:             entity.FormID(dbModel.ID),
		OrganizationID: entity.OrganizationID(dbModel.OrganizationID),
		Year:           dbModel.Year,
		GeneralInfo:    generalInfo,
		MuseumData:     museumDetails,
		Labor:          laborData,
		Finances:       financialData,
		Status:         entity.FormStatus(dbModel.Status),
	}, nil
}

func (r *reportingFormRepository) fetchGeneralInfo(ctx context.Context, orgID int) (value.FormGeneralInfo, error) {
	var org model.Organization
	if err := r.db.GetContext(ctx, &org, getOrganizationGeneralInfoQuery, orgID); err != nil {
		return value.FormGeneralInfo{}, err
	}

	info, err := value.NewFormGeneralInfo(org.INN, org.Name, org.ExistMuseum)
	if err != nil {
		return value.FormGeneralInfo{}, err
	}

	return info, nil
}

func (r *reportingFormRepository) fetchMuseumDetails(ctx context.Context, orgID int) (value.MuseumDetails, error) {
	var museum model.Museum
	err := r.db.GetContext(ctx, &museum, getPrimaryMuseumByOwnerQuery, orgID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return value.MuseumDetails{}, nil
		}
		return value.MuseumDetails{}, err
	}

	return value.MuseumDetails{
		INN:                              museum.INN,
		KPP:                              museum.KPP,
		Founder:                          museum.Founder,
		MuseumActivityInCharter:          museum.MuseumActivityInCharter,
		Name:                             museum.Name,
		MuseumLegalStatus:                museum.MuseumLegalStatus,
		IsMemorialReserveMuseum:          museum.IsMemorialReserveMuseum,
		IsHistoricalMemorialReserve:      museum.IsHistoricalMemorialReserve,
		IsArtMuseum:                      museum.IsArtMuseum,
		IsMuseumReserve:                  museum.IsMuseumReserve,
		IsEstateMuseum:                   museum.IsEstateMuseum,
		IsPalaceParkEnsemble:             museum.IsPalaceParkEnsemble,
		IsHistoricalArchitecturalReserve: museum.IsHistoricalArchitecturalReserve,
		AnnualVisitorCapacity:            museum.AnnualVisitorCapacity,
		InternalVisitorsCount:            museum.InternalVisitorsCount,
		ExternalVisitorsCount:            museum.ExternalVisitorsCount,
		IsValuableCulturalHeritage:       museum.IsValuableCulturalHeritage,
		ValuableMuseumItemsCount:         museum.ValuableMuseumItemsCount,
	}, nil
}

func (r *reportingFormRepository) fetchLaborData(ctx context.Context, orgID int) (value.LaborData, error) {
	var labor value.LaborData

	var laborModel model.LaborResources
	err := r.db.GetContext(ctx, &laborModel, selectLaborResourcesByOwnerQuery, orgID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return labor, nil
		}
		return labor, err
	}

	labor.TotalStaffAnnual = laborModel.TotalStaffAnnual
	labor.ResearchStaffInternal = laborModel.ResearchStaffInternal
	labor.CoreOperationalStaffInternal = laborModel.CoreOperationalStaffInternal
	labor.AdminSupportStaffInternal = laborModel.AdminSupportStaffInternal
	labor.ResearchStaffExternal = laborModel.ResearchStaffExternal
	labor.CoreOperationalStaffExternal = laborModel.CoreOperationalStaffExternal
	labor.AdminSupportStaffExternal = laborModel.AdminSupportStaffExternal

	var fotModel model.LaborResourcesFOT
	err = r.db.GetContext(ctx, &fotModel, selectLaborFOTByLaborIDQuery, laborModel.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return labor, nil
		}
		return labor, err
	}

	labor.FOT = value.LaborFOT{
		TotalStaffAnnual:             fotModel.TotalStaffAnnual,
		ResearchStaffInternal:        fotModel.ResearchStaffInternal,
		CoreOperationalStaffInternal: fotModel.CoreOperationalStaffInternal,
		AdminSupportStaffInternal:    fotModel.AdminSupportStaffInternal,
		ResearchStaffExternal:        fotModel.ResearchStaffExternal,
		CoreOperationalStaffExternal: fotModel.CoreOperationalStaffExternal,
		AdminSupportStaffExternal:    fotModel.AdminSupportStaffExternal,
	}

	return labor, nil
}

func (r *reportingFormRepository) fetchFinancialData(ctx context.Context, orgID int, year int16) (value.FinancialData, error) {
	var finance model.MuseumRevenuesExpenses
	err := r.db.GetContext(ctx, &finance, selectFinancialDataQuery, orgID, year)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return value.FinancialData{}, nil
		}
		return value.FinancialData{}, err
	}

	return value.FinancialData{
		TotalRevenue:               finance.TotalRevenue,
		StateAssignmentSubsidy:     finance.StateAssignmentSubsidy,
		EarnedRevenue:              finance.EarnedRevenue,
		OtherFundingSources:        finance.OtherFundingSources,
		TotalExpenses:              finance.TotalExpenses,
		InventoryAssetsAcquisition: finance.InventoryAssetsAcquisition,
		UtilityServices:            finance.UtilityServices,
		CommunicationServices:      finance.CommunicationServices,
		TransportationServices:     finance.TransportationServices,
		ValuableAssetsAcquisition:  finance.ValuableAssetsAcquisition,
		RealEstateMaintenance:      finance.RealEstateMaintenance,
		ValuableAssetsMaintenance:  finance.ValuableAssetsMaintenance,
		GeneralAdministrativeCosts: finance.GeneralAdministrativeCosts,
		TaxPayments:                finance.TaxPayments,
		OtherExpenses:              finance.OtherExpenses,
	}, nil
}

func (r *reportingFormRepository) syncFormDetails(ctx context.Context, tx *sqlx.Tx, form *entity.ReportingForm) error {
	if err := r.updateOrganization(ctx, tx, form.OrganizationID, form.GeneralInfo); err != nil {
		return err
	}

	if err := r.upsertMuseum(ctx, tx, form.OrganizationID, form.MuseumData); err != nil {
		return err
	}

	laborID, err := r.upsertLaborResources(ctx, tx, form.OrganizationID, form.Labor)
	if err != nil {
		return err
	}

	if laborID != 0 {
		if err := r.upsertLaborFOT(ctx, tx, laborID, form.Labor.FOT); err != nil {
			return err
		}
	}

	if err := r.upsertFinancialData(ctx, tx, form.OrganizationID, form.Year, form.Finances); err != nil {
		return err
	}

	return nil
}

func (r *reportingFormRepository) updateOrganization(ctx context.Context, tx *sqlx.Tx, orgID entity.OrganizationID, info value.FormGeneralInfo) error {
	if info.INN == "" && info.Name == "" {
		return nil
	}

	_, err := tx.ExecContext(ctx, updateOrganizationGeneralInfoQuery, orgID, info.INN, info.Name, info.ExistMuseum)
	return err
}

func (r *reportingFormRepository) upsertMuseum(ctx context.Context, tx *sqlx.Tx, orgID entity.OrganizationID, details value.MuseumDetails) error {
	modelMuseum := model.Museum{
		IDOwner:                          int(orgID),
		INN:                              details.INN,
		KPP:                              details.KPP,
		Founder:                          details.Founder,
		MuseumActivityInCharter:          details.MuseumActivityInCharter,
		Name:                             details.Name,
		MuseumLegalStatus:                details.MuseumLegalStatus,
		IsMemorialReserveMuseum:          details.IsMemorialReserveMuseum,
		IsHistoricalMemorialReserve:      details.IsHistoricalMemorialReserve,
		IsArtMuseum:                      details.IsArtMuseum,
		IsMuseumReserve:                  details.IsMuseumReserve,
		IsEstateMuseum:                   details.IsEstateMuseum,
		IsPalaceParkEnsemble:             details.IsPalaceParkEnsemble,
		IsHistoricalArchitecturalReserve: details.IsHistoricalArchitecturalReserve,
		AnnualVisitorCapacity:            details.AnnualVisitorCapacity,
		InternalVisitorsCount:            details.InternalVisitorsCount,
		ExternalVisitorsCount:            details.ExternalVisitorsCount,
		IsValuableCulturalHeritage:       details.IsValuableCulturalHeritage,
		ValuableMuseumItemsCount:         details.ValuableMuseumItemsCount,
	}

	var existingID sql.NullInt64
	if err := tx.QueryRowxContext(ctx, selectMuseumIDByOwnerQuery, orgID).Scan(&existingID); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}

	if existingID.Valid {
		modelMuseum.Id = int(existingID.Int64)
		_, err := tx.NamedExecContext(ctx, updateMuseumByIDQuery, modelMuseum)
		return err
	}

	query, args, err := sqlx.Named(insertMuseumQuery, modelMuseum)
	if err != nil {
		return err
	}
	query = tx.Rebind(query)
	return tx.QueryRowxContext(ctx, query, args...).Scan(&modelMuseum.Id)
}

func (r *reportingFormRepository) upsertLaborResources(ctx context.Context, tx *sqlx.Tx, orgID entity.OrganizationID, data value.LaborData) (int, error) {
	modelLabor := model.LaborResources{
		IDOwner:                      int(orgID),
		TotalStaffAnnual:             data.TotalStaffAnnual,
		ResearchStaffInternal:        data.ResearchStaffInternal,
		CoreOperationalStaffInternal: data.CoreOperationalStaffInternal,
		AdminSupportStaffInternal:    data.AdminSupportStaffInternal,
		ResearchStaffExternal:        data.ResearchStaffExternal,
		CoreOperationalStaffExternal: data.CoreOperationalStaffExternal,
		AdminSupportStaffExternal:    data.AdminSupportStaffExternal,
	}

	var existing model.LaborResources
	err := tx.GetContext(ctx, &existing, selectLaborResourcesByOwnerQuery, orgID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			query, args, err := sqlx.Named(insertLaborResourcesQuery, modelLabor)
			if err != nil {
				return 0, err
			}
			query = tx.Rebind(query)

			var newID int
			if err := tx.QueryRowxContext(ctx, query, args...).Scan(&newID); err != nil {
				return 0, err
			}

			return newID, nil
		}
		return 0, err
	}

	modelLabor.ID = existing.ID
	if _, err = tx.NamedExecContext(ctx, updateLaborResourcesQuery, modelLabor); err != nil {
		return 0, err
	}

	return existing.ID, nil
}

func (r *reportingFormRepository) upsertLaborFOT(ctx context.Context, tx *sqlx.Tx, laborID int, fot value.LaborFOT) error {
	if laborID == 0 {
		return nil
	}

	modelFOT := model.LaborResourcesFOT{
		LaborResourcesID:             laborID,
		TotalStaffAnnual:             fot.TotalStaffAnnual,
		ResearchStaffInternal:        fot.ResearchStaffInternal,
		CoreOperationalStaffInternal: fot.CoreOperationalStaffInternal,
		AdminSupportStaffInternal:    fot.AdminSupportStaffInternal,
		ResearchStaffExternal:        fot.ResearchStaffExternal,
		CoreOperationalStaffExternal: fot.CoreOperationalStaffExternal,
		AdminSupportStaffExternal:    fot.AdminSupportStaffExternal,
	}

	var existing model.LaborResourcesFOT
	err := tx.GetContext(ctx, &existing, selectLaborFOTByLaborIDQuery, laborID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			_, err = tx.NamedExecContext(ctx, insertLaborFOTQuery, modelFOT)
			return err
		}
		return err
	}

	_, err = tx.NamedExecContext(ctx, updateLaborFOTQuery, modelFOT)
	return err
}

func (r *reportingFormRepository) upsertFinancialData(ctx context.Context, tx *sqlx.Tx, orgID entity.OrganizationID, year int16, data value.FinancialData) error {
	modelFinance := model.MuseumRevenuesExpenses{
		IDOwner:                    int(orgID),
		Year:                       year,
		TotalRevenue:               data.TotalRevenue,
		StateAssignmentSubsidy:     data.StateAssignmentSubsidy,
		EarnedRevenue:              data.EarnedRevenue,
		OtherFundingSources:        data.OtherFundingSources,
		TotalExpenses:              data.TotalExpenses,
		InventoryAssetsAcquisition: data.InventoryAssetsAcquisition,
		UtilityServices:            data.UtilityServices,
		CommunicationServices:      data.CommunicationServices,
		TransportationServices:     data.TransportationServices,
		ValuableAssetsAcquisition:  data.ValuableAssetsAcquisition,
		RealEstateMaintenance:      data.RealEstateMaintenance,
		ValuableAssetsMaintenance:  data.ValuableAssetsMaintenance,
		GeneralAdministrativeCosts: data.GeneralAdministrativeCosts,
		TaxPayments:                data.TaxPayments,
		OtherExpenses:              data.OtherExpenses,
	}

	var existing model.MuseumRevenuesExpenses
	err := tx.GetContext(ctx, &existing, selectFinancialDataQuery, orgID, year)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			_, err = tx.NamedExecContext(ctx, insertFinancialDataQuery, modelFinance)
			return err
		}
		return err
	}

	_, err = tx.NamedExecContext(ctx, updateFinancialDataQuery, modelFinance)
	return err
}
