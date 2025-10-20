package model

type MuseumRevenuesExpenses struct {
	IDOwner                int     `db:"id_owner"`
	Year                   int16   `db:"year"`
	TotalRevenue           float64 `db:"total_revenue"`
	StateAssignmentSubsidy float64 `db:"state_assignment_subsidy"`
	EarnedRevenue          float64 `db:"earned_revenue"`
	OtherFundingSources    float64 `db:"other_funding_sources"`

	TotalExpenses              float64 `db:"total_expenses"`
	InventoryAssetsAcquisition float64 `db:"inventory_assets_acquisition"`
	UtilityServices            float64 `db:"utility_services"`
	CommunicationServices      float64 `db:"communication_services"`
	TransportationServices     float64 `db:"transportation_services"`
	ValuableAssetsAcquisition  float64 `db:"valuable_assets_acquisition"`
	RealEstateMaintenance      float64 `db:"real_estate_maintenance"`
	ValuableAssetsMaintenance  float64 `db:"valuable_assets_maintenance"`
	GeneralAdministrativeCosts float64 `db:"general_administrative_costs"`
	TaxPayments                float64 `db:"tax_payments"`
	OtherExpenses              float64 `db:"other_expenses"`
}
