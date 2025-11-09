package reportingform

const (
	createReportingFormQuery = `
		INSERT INTO reporting_forms (organization_id, year, status)
		VALUES ($1, $2, $3)
		RETURNING id`

	getReportingFormByIDQuery = `
		SELECT id, organization_id, year, status
		FROM reporting_forms
		WHERE id = $1`

	getReportingFormByOrganizationAndYearQuery = `
		SELECT id, organization_id, year, status
		FROM reporting_forms
		WHERE organization_id = $1 AND year = $2`

	updateReportingFormQuery = `
		UPDATE reporting_forms
		SET year = :year,
		    status = :status
		WHERE id = :id`

	deleteReportingFormQuery = `
		DELETE FROM reporting_forms
		WHERE id = $1`

	listReportingFormsQuery = `
		SELECT id, organization_id, year, status
		FROM reporting_forms
		ORDER BY year DESC, id`

	listReportingFormsByOrganizationQuery = `
		SELECT id, organization_id, year, status
		FROM reporting_forms
		WHERE organization_id = $1
		ORDER BY year DESC, id`

	listReportingFormsByYearQuery = `
		SELECT id, organization_id, year, status
		FROM reporting_forms
		WHERE year = $1
		ORDER BY organization_id, id`

	getOrganizationGeneralInfoQuery = `
		SELECT id, inn, name, exist_museum
		FROM organization
		WHERE id = $1`

	updateOrganizationGeneralInfoQuery = `
		UPDATE organization
		SET inn = $2, name = $3, exist_museum = $4
		WHERE id = $1`

	getPrimaryMuseumByOwnerQuery = `
		SELECT
			id,
			id_owner,
			inn,
			kpp,
			founder,
			museum_activity_in_charter,
			name,
			museum_legal_status,
			is_memorial_reserve_museum,
			is_historical_memorial_reserve,
			is_art_museum,
			is_museum_reserve,
			is_estate_museum,
			is_palace_park_ensemble,
			is_historical_architectural_reserve,
			annual_visitor_capacity,
			internal_visitors_count,
			external_visitors_count,
			is_valuable_cultural_heritage,
			valuable_museum_items_count
		FROM museum
		WHERE id_owner = $1
		ORDER BY id
		LIMIT 1`

	selectMuseumIDByOwnerQuery = `
		SELECT id
		FROM museum
		WHERE id_owner = $1
		ORDER BY id
		LIMIT 1`

	updateMuseumByIDQuery = `
		UPDATE museum SET
			id_owner = :id_owner,
			inn = :inn,
			kpp = :kpp,
			founder = :founder,
			museum_activity_in_charter = :museum_activity_in_charter,
			name = :name,
			museum_legal_status = :museum_legal_status,
			is_memorial_reserve_museum = :is_memorial_reserve_museum,
			is_historical_memorial_reserve = :is_historical_memorial_reserve,
			is_art_museum = :is_art_museum,
			is_museum_reserve = :is_museum_reserve,
			is_estate_museum = :is_estate_museum,
			is_palace_park_ensemble = :is_palace_park_ensemble,
			is_historical_architectural_reserve = :is_historical_architectural_reserve,
			annual_visitor_capacity = :annual_visitor_capacity,
			internal_visitors_count = :internal_visitors_count,
			external_visitors_count = :external_visitors_count,
			is_valuable_cultural_heritage = :is_valuable_cultural_heritage,
			valuable_museum_items_count = :valuable_museum_items_count
		WHERE id = :id`

	insertMuseumQuery = `
		INSERT INTO museum (
			id_owner,
			inn,
			kpp,
			founder,
			museum_activity_in_charter,
			name,
			museum_legal_status,
			is_memorial_reserve_museum,
			is_historical_memorial_reserve,
			is_art_museum,
			is_museum_reserve,
			is_estate_museum,
			is_palace_park_ensemble,
			is_historical_architectural_reserve,
			annual_visitor_capacity,
			internal_visitors_count,
			external_visitors_count,
			is_valuable_cultural_heritage,
			valuable_museum_items_count
		)
		VALUES (
			:id_owner,
			:inn,
			:kpp,
			:founder,
			:museum_activity_in_charter,
			:name,
			:museum_legal_status,
			:is_memorial_reserve_museum,
			:is_historical_memorial_reserve,
			:is_art_museum,
			:is_museum_reserve,
			:is_estate_museum,
			:is_palace_park_ensemble,
			:is_historical_architectural_reserve,
			:annual_visitor_capacity,
			:internal_visitors_count,
			:external_visitors_count,
			:is_valuable_cultural_heritage,
			:valuable_museum_items_count
		)
		RETURNING id`

	selectLaborResourcesByOwnerQuery = `
		SELECT
			id,
			id_owner,
			total_staff_annual,
			research_staff_internal,
			core_operational_staff_internal,
			admin_support_staff_internal,
			research_staff_external,
			core_operational_staff_external,
			admin_support_staff_external
		FROM labor_resource
		WHERE id_owner = $1
		ORDER BY id
		LIMIT 1`

	insertLaborResourcesQuery = `
		INSERT INTO labor_resource (
			id_owner,
			total_staff_annual,
			research_staff_internal,
			core_operational_staff_internal,
			admin_support_staff_internal,
			research_staff_external,
			core_operational_staff_external,
			admin_support_staff_external
		)
		VALUES (
			:id_owner,
			:total_staff_annual,
			:research_staff_internal,
			:core_operational_staff_internal,
			:admin_support_staff_internal,
			:research_staff_external,
			:core_operational_staff_external,
			:admin_support_staff_external
		)
		RETURNING id`

	updateLaborResourcesQuery = `
		UPDATE labor_resource SET
			total_staff_annual = :total_staff_annual,
			research_staff_internal = :research_staff_internal,
			core_operational_staff_internal = :core_operational_staff_internal,
			admin_support_staff_internal = :admin_support_staff_internal,
			research_staff_external = :research_staff_external,
			core_operational_staff_external = :core_operational_staff_external,
			admin_support_staff_external = :admin_support_staff_external
		WHERE id = :id`

	selectLaborFOTByLaborIDQuery = `
		SELECT
			labor_resources_id,
			total_staff_annual,
			research_staff_internal,
			core_operational_staff_internal,
			admin_support_staff_internal,
			research_staff_external,
			core_operational_staff_external,
			admin_support_staff_external
		FROM labor_resource_fot
		WHERE labor_resources_id = $1
		LIMIT 1`

	insertLaborFOTQuery = `
		INSERT INTO labor_resource_fot (
			labor_resources_id,
			total_staff_annual,
			research_staff_internal,
			core_operational_staff_internal,
			admin_support_staff_internal,
			research_staff_external,
			core_operational_staff_external,
			admin_support_staff_external
		)
		VALUES (
			:labor_resources_id,
			:total_staff_annual,
			:research_staff_internal,
			:core_operational_staff_internal,
			:admin_support_staff_internal,
			:research_staff_external,
			:core_operational_staff_external,
			:admin_support_staff_external
		)`

	updateLaborFOTQuery = `
		UPDATE labor_resource_fot SET
			total_staff_annual = :total_staff_annual,
			research_staff_internal = :research_staff_internal,
			core_operational_staff_internal = :core_operational_staff_internal,
			admin_support_staff_internal = :admin_support_staff_internal,
			research_staff_external = :research_staff_external,
			core_operational_staff_external = :core_operational_staff_external,
			admin_support_staff_external = :admin_support_staff_external
		WHERE labor_resources_id = :labor_resources_id`

	selectFinancialDataQuery = `
		SELECT
			id_owner,
			year,
			total_revenue,
			state_assignment_subsidy,
			earned_revenue,
			other_funding_sources,
			total_expenses,
			inventory_assets_acquisition,
			utility_services,
			communication_services,
			transportation_services,
			valuable_assets_acquisition,
			real_estate_maintenance,
			valuable_assets_maintenance,
			general_administrative_costs,
			tax_payments,
			other_expenses
		FROM museum_revenues_expense
		WHERE id_owner = $1 AND year = $2
		LIMIT 1`

	insertFinancialDataQuery = `
		INSERT INTO museum_revenues_expense (
			id_owner,
			year,
			total_revenue,
			state_assignment_subsidy,
			earned_revenue,
			other_funding_sources,
			total_expenses,
			inventory_assets_acquisition,
			utility_services,
			communication_services,
			transportation_services,
			valuable_assets_acquisition,
			real_estate_maintenance,
			valuable_assets_maintenance,
			general_administrative_costs,
			tax_payments,
			other_expenses
		)
		VALUES (
			:id_owner,
			:year,
			:total_revenue,
			:state_assignment_subsidy,
			:earned_revenue,
			:other_funding_sources,
			:total_expenses,
			:inventory_assets_acquisition,
			:utility_services,
			:communication_services,
			:transportation_services,
			:valuable_assets_acquisition,
			:real_estate_maintenance,
			:valuable_assets_maintenance,
			:general_administrative_costs,
			:tax_payments,
			:other_expenses
		)`

	updateFinancialDataQuery = `
		UPDATE museum_revenues_expense SET
			total_revenue = :total_revenue,
			state_assignment_subsidy = :state_assignment_subsidy,
			earned_revenue = :earned_revenue,
			other_funding_sources = :other_funding_sources,
			total_expenses = :total_expenses,
			inventory_assets_acquisition = :inventory_assets_acquisition,
			utility_services = :utility_services,
			communication_services = :communication_services,
			transportation_services = :transportation_services,
			valuable_assets_acquisition = :valuable_assets_acquisition,
			real_estate_maintenance = :real_estate_maintenance,
			valuable_assets_maintenance = :valuable_assets_maintenance,
			general_administrative_costs = :general_administrative_costs,
			tax_payments = :tax_payments,
			other_expenses = :other_expenses
		WHERE id_owner = :id_owner AND year = :year`
)
