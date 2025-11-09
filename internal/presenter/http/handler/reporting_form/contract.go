package reportingform

import (
	"vkr/internal/domain/entity"
	"vkr/internal/domain/value"
)

type BadRequestResponse struct {
	ErrorMsg string `json:"error"`
}

type InternalServerErrorResponse struct {
	ErrorMsg string `json:"error"`
}

type NotFoundResponse struct {
	ErrorMsg string `json:"error"`
}

var (
	ErrInternalServer = &InternalServerErrorResponse{ErrorMsg: "internal server error"}
)

type ReportingFormRequest struct {
	OrganizationID entity.OrganizationID         `json:"organization_id"`
	Year           int16                         `json:"year"`
	Status         entity.FormStatus             `json:"status"`
	GeneralInfo    ReportingFormGeneralInfoDTO   `json:"general_info"`
	MuseumData     ReportingFormMuseumDetailsDTO `json:"museum_data"`
	Labor          ReportingFormLaborDTO         `json:"labor"`
	Finances       ReportingFormFinancialsDTO    `json:"finances"`
}

type ReportingFormUpdateRequest struct {
	ID entity.FormID `json:"id"`
	ReportingFormRequest
}

type ReportingFormGeneralInfoDTO struct {
	INN         string `json:"inn"`
	Name        string `json:"name"`
	ExistMuseum bool   `json:"exist_museum"`
}

type ReportingFormMuseumDetailsDTO struct {
	INN                              string  `json:"inn"`
	KPP                              *string `json:"kpp,omitempty"`
	Founder                          *string `json:"founder,omitempty"`
	MuseumActivityInCharter          bool    `json:"museum_activity_in_charter"`
	Name                             string  `json:"name"`
	MuseumLegalStatus                *string `json:"museum_legal_status,omitempty"`
	IsMemorialReserveMuseum          bool    `json:"is_memorial_reserve_museum"`
	IsHistoricalMemorialReserve      bool    `json:"is_historical_memorial_reserve"`
	IsArtMuseum                      bool    `json:"is_art_museum"`
	IsMuseumReserve                  bool    `json:"is_museum_reserve"`
	IsEstateMuseum                   bool    `json:"is_estate_museum"`
	IsPalaceParkEnsemble             bool    `json:"is_palace_park_ensemble"`
	IsHistoricalArchitecturalReserve bool    `json:"is_historical_architectural_reserve"`
	AnnualVisitorCapacity            *int    `json:"annual_visitor_capacity,omitempty"`
	InternalVisitorsCount            *int    `json:"internal_visitors_count,omitempty"`
	ExternalVisitorsCount            *int    `json:"external_visitors_count,omitempty"`
	IsValuableCulturalHeritage       bool    `json:"is_valuable_cultural_heritage"`
	ValuableMuseumItemsCount         *int    `json:"valuable_museum_items_count,omitempty"`
}

type ReportingFormLaborDTO struct {
	TotalStaffAnnual             float64                  `json:"total_staff_annual"`
	ResearchStaffInternal        float64                  `json:"research_staff_internal"`
	CoreOperationalStaffInternal float64                  `json:"core_operational_staff_internal"`
	AdminSupportStaffInternal    float64                  `json:"admin_support_staff_internal"`
	ResearchStaffExternal        float64                  `json:"research_staff_external"`
	CoreOperationalStaffExternal float64                  `json:"core_operational_staff_external"`
	AdminSupportStaffExternal    float64                  `json:"admin_support_staff_external"`
	FOT                          ReportingFormLaborFOTDTO `json:"fot"`
}

type ReportingFormLaborFOTDTO struct {
	TotalStaffAnnual             float64 `json:"total_staff_annual"`
	ResearchStaffInternal        float64 `json:"research_staff_internal"`
	CoreOperationalStaffInternal float64 `json:"core_operational_staff_internal"`
	AdminSupportStaffInternal    float64 `json:"admin_support_staff_internal"`
	ResearchStaffExternal        float64 `json:"research_staff_external"`
	CoreOperationalStaffExternal float64 `json:"core_operational_staff_external"`
	AdminSupportStaffExternal    float64 `json:"admin_support_staff_external"`
}

type ReportingFormFinancialsDTO struct {
	TotalRevenue               float64 `json:"total_revenue"`
	StateAssignmentSubsidy     float64 `json:"state_assignment_subsidy"`
	EarnedRevenue              float64 `json:"earned_revenue"`
	OtherFundingSources        float64 `json:"other_funding_sources"`
	TotalExpenses              float64 `json:"total_expenses"`
	InventoryAssetsAcquisition float64 `json:"inventory_assets_acquisition"`
	UtilityServices            float64 `json:"utility_services"`
	CommunicationServices      float64 `json:"communication_services"`
	TransportationServices     float64 `json:"transportation_services"`
	ValuableAssetsAcquisition  float64 `json:"valuable_assets_acquisition"`
	RealEstateMaintenance      float64 `json:"real_estate_maintenance"`
	ValuableAssetsMaintenance  float64 `json:"valuable_assets_maintenance"`
	GeneralAdministrativeCosts float64 `json:"general_administrative_costs"`
	TaxPayments                float64 `json:"tax_payments"`
	OtherExpenses              float64 `json:"other_expenses"`
}

type ReportingFormResponse = entity.ReportingForm
type ReportingFormListResponse = []*entity.ReportingForm

func (dto ReportingFormRequest) ToEntity() (*entity.ReportingForm, error) {
	generalInfo, err := value.NewFormGeneralInfo(dto.GeneralInfo.INN, dto.GeneralInfo.Name, dto.GeneralInfo.ExistMuseum)
	if err != nil {
		return nil, err
	}

	labor := value.LaborData{
		TotalStaffAnnual:             dto.Labor.TotalStaffAnnual,
		ResearchStaffInternal:        dto.Labor.ResearchStaffInternal,
		CoreOperationalStaffInternal: dto.Labor.CoreOperationalStaffInternal,
		AdminSupportStaffInternal:    dto.Labor.AdminSupportStaffInternal,
		ResearchStaffExternal:        dto.Labor.ResearchStaffExternal,
		CoreOperationalStaffExternal: dto.Labor.CoreOperationalStaffExternal,
		AdminSupportStaffExternal:    dto.Labor.AdminSupportStaffExternal,
		FOT: value.LaborFOT{
			TotalStaffAnnual:             dto.Labor.FOT.TotalStaffAnnual,
			ResearchStaffInternal:        dto.Labor.FOT.ResearchStaffInternal,
			CoreOperationalStaffInternal: dto.Labor.FOT.CoreOperationalStaffInternal,
			AdminSupportStaffInternal:    dto.Labor.FOT.AdminSupportStaffInternal,
			ResearchStaffExternal:        dto.Labor.FOT.ResearchStaffExternal,
			CoreOperationalStaffExternal: dto.Labor.FOT.CoreOperationalStaffExternal,
			AdminSupportStaffExternal:    dto.Labor.FOT.AdminSupportStaffExternal,
		},
	}

	finances := value.FinancialData{
		TotalRevenue:               dto.Finances.TotalRevenue,
		StateAssignmentSubsidy:     dto.Finances.StateAssignmentSubsidy,
		EarnedRevenue:              dto.Finances.EarnedRevenue,
		OtherFundingSources:        dto.Finances.OtherFundingSources,
		TotalExpenses:              dto.Finances.TotalExpenses,
		InventoryAssetsAcquisition: dto.Finances.InventoryAssetsAcquisition,
		UtilityServices:            dto.Finances.UtilityServices,
		CommunicationServices:      dto.Finances.CommunicationServices,
		TransportationServices:     dto.Finances.TransportationServices,
		ValuableAssetsAcquisition:  dto.Finances.ValuableAssetsAcquisition,
		RealEstateMaintenance:      dto.Finances.RealEstateMaintenance,
		ValuableAssetsMaintenance:  dto.Finances.ValuableAssetsMaintenance,
		GeneralAdministrativeCosts: dto.Finances.GeneralAdministrativeCosts,
		TaxPayments:                dto.Finances.TaxPayments,
		OtherExpenses:              dto.Finances.OtherExpenses,
	}

	return &entity.ReportingForm{
		OrganizationID: dto.OrganizationID,
		Year:           dto.Year,
		Status:         dto.Status,
		GeneralInfo:    generalInfo,
		MuseumData: value.MuseumDetails{
			INN:                              dto.MuseumData.INN,
			KPP:                              dto.MuseumData.KPP,
			Founder:                          dto.MuseumData.Founder,
			MuseumActivityInCharter:          dto.MuseumData.MuseumActivityInCharter,
			Name:                             dto.MuseumData.Name,
			MuseumLegalStatus:                dto.MuseumData.MuseumLegalStatus,
			IsMemorialReserveMuseum:          dto.MuseumData.IsMemorialReserveMuseum,
			IsHistoricalMemorialReserve:      dto.MuseumData.IsHistoricalMemorialReserve,
			IsArtMuseum:                      dto.MuseumData.IsArtMuseum,
			IsMuseumReserve:                  dto.MuseumData.IsMuseumReserve,
			IsEstateMuseum:                   dto.MuseumData.IsEstateMuseum,
			IsPalaceParkEnsemble:             dto.MuseumData.IsPalaceParkEnsemble,
			IsHistoricalArchitecturalReserve: dto.MuseumData.IsHistoricalArchitecturalReserve,
			AnnualVisitorCapacity:            dto.MuseumData.AnnualVisitorCapacity,
			InternalVisitorsCount:            dto.MuseumData.InternalVisitorsCount,
			ExternalVisitorsCount:            dto.MuseumData.ExternalVisitorsCount,
			IsValuableCulturalHeritage:       dto.MuseumData.IsValuableCulturalHeritage,
			ValuableMuseumItemsCount:         dto.MuseumData.ValuableMuseumItemsCount,
		},
		Labor:    labor,
		Finances: finances,
	}, nil
}
