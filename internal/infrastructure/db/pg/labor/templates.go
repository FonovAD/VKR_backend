package labor

const (
	getLaborResourcesByOwnerQuery = `
		SELECT
			lr.id,
			lr.museum_id,
			lr.total_staff_annual,
			lr.research_staff_internal,
			lr.core_operational_staff_internal,
			lr.admin_support_staff_internal,
			lr.research_staff_external,
			lr.core_operational_staff_external,
			lr.admin_support_staff_external
		FROM labor_resource lr
		JOIN museum m ON m.id = lr.museum_id
		WHERE m.organization_id = $1
		LIMIT 1`
	
	getLaborResourcesByINNQuery = `
		SELECT
			lr.id,
			lr.museum_id,
			lr.total_staff_annual,
			lr.research_staff_internal,
			lr.core_operational_staff_internal,
			lr.admin_support_staff_internal,
			lr.research_staff_external,
			lr.core_operational_staff_external,
			lr.admin_support_staff_external
		FROM labor_resource lr
		JOIN museum m ON m.id = lr.museum_id
		JOIN organization o ON o.id = m.organization_id
		WHERE LOWER(o.inn) = LOWER($1)
		LIMIT 1`

	getLaborFotByLaborIDQuery = `
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

	getOrganizationByINNQuery = `
		SELECT
			id
		FROM organization
		WHERE LOWER(inn) = LOWER($1)
		LIMIT 1`
)
