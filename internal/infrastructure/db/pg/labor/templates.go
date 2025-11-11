package labor

const (
	getLaborResourcesByOwnerQuery = `
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
			id,
			inn,
			name,
			exist_museum
		FROM organization
		WHERE inn = $1
		LIMIT 1`
)
