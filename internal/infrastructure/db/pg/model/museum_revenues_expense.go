package model

type MuseumRevenuesExpenses struct {
	ID                       int      `db:"id"`
	MuseumID                 int      `db:"museum_id"`
	Year                     int16    `db:"year"`
	TotalRevenue             *float64 `db:"total_revenue"`
	StateAssignmentSubsidy   *float64 `db:"state_assignment_subsidy"`
	EarnedRevenue            *float64 `db:"earned_revenue"`

	TotalExpenses              *float64 `db:"total_expenses"`
	InventoryAssetsAcquisition *float64 `db:"inventory_assets_acquisition"`
	UtilityServices            *float64 `db:"utility_services"`
	CommunicationServices      *float64 `db:"communication_services"`
	TransportationServices     *float64 `db:"transportation_services"`
	ValuableAssetsAcquisition  *float64 `db:"valuable_assets_acquisition"`
	RealEstateMaintenance      *float64 `db:"real_estate_maintenance"`
	ValuableAssetsMaintenance  *float64 `db:"valuable_assets_maintenance"`
	GeneralAdministrativeCosts *float64 `db:"general_administrative_costs"`
	TaxPayments                *float64 `db:"tax_payments"`
}

type OtherFundingSource struct {
	ID                     int     `db:"id"`
	MuseumRevenuesExpenseID int     `db:"museum_revenues_expense_id"`
	Amount                 float64 `db:"amount"`
}

type OtherExpense struct {
	ID                     int     `db:"id"`
	MuseumRevenuesExpenseID int     `db:"museum_revenues_expense_id"`
	Amount                 float64 `db:"amount"`
}
