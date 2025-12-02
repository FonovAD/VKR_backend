package reporting_form

const (
	getFinancialDataByOrgIDAndYearQuery = `
		SELECT
			mre.id,
			mre.museum_id,
			mre.year,
			mre.total_revenue,
			mre.state_assignment_subsidy,
			mre.earned_revenue,
			mre.total_expenses,
			mre.inventory_assets_acquisition,
			mre.utility_services,
			mre.communication_services,
			mre.transportation_services,
			mre.valuable_assets_acquisition,
			mre.real_estate_maintenance,
			mre.valuable_assets_maintenance,
			mre.general_administrative_costs,
			mre.tax_payments
		FROM museum_revenues_expense mre
		JOIN museum m ON m.id = mre.museum_id
		WHERE m.organization_id = $1 AND mre.year = $2
		LIMIT 1`

	getOtherFundingSourcesSumQuery = `
		SELECT COALESCE(SUM(amount), 0) as total
		FROM other_funding_sources
		WHERE museum_revenues_expense_id = $1`

	getOtherExpensesSumQuery = `
		SELECT COALESCE(SUM(amount), 0) as total
		FROM other_expenses
		WHERE museum_revenues_expense_id = $1`

	getFinancialDataByINNAndYearQuery = `
		SELECT
			mre.id,
			mre.museum_id,
			mre.year,
			mre.total_revenue,
			mre.state_assignment_subsidy,
			mre.earned_revenue,
			mre.total_expenses,
			mre.inventory_assets_acquisition,
			mre.utility_services,
			mre.communication_services,
			mre.transportation_services,
			mre.valuable_assets_acquisition,
			mre.real_estate_maintenance,
			mre.valuable_assets_maintenance,
			mre.general_administrative_costs,
			mre.tax_payments
		FROM museum_revenues_expense mre
		JOIN museum m ON m.id = mre.museum_id
		JOIN organization o ON o.id = m.organization_id
		WHERE LOWER(o.inn) = LOWER($1) AND mre.year = $2
		LIMIT 1`

	getOrganizationByINNQuery = `
		SELECT
			id
		FROM organization
		WHERE LOWER(inn) = LOWER($1)
		LIMIT 1`
)

