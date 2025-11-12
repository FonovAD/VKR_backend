package organization

const (
	createOrganizationQuery = `
		INSERT INTO organization (inn, name, exist_museum)
		VALUES ($1, $2, $3)
		RETURNING id`

	getOrganizationByIDQuery = `
		SELECT id, inn, name, exist_museum
		FROM organization
		WHERE id = $1`

	updateOrganizationQuery = `
		UPDATE organization
		SET inn = :inn, name = :name, exist_museum = :exist_museum
		WHERE id = :id`

	deleteOrganizationQuery = `
		DELETE FROM organization
		WHERE id = $1`

	findOrganizationByINNQuery = `
		SELECT id, inn, name, exist_museum
		FROM organization
		WHERE LOWER(inn) = LOWER($1)`

	listOrganizationsQuery = `
		SELECT id, inn, name, exist_museum
		FROM organization
		WHERE ($1::text IS NULL OR name ILIKE '%' || $1 || '%')
		ORDER BY id
		LIMIT $2 OFFSET $3`

	countOrganizationsQuery = `
		SELECT COUNT(*)
		FROM organization
		WHERE ($1::text IS NULL OR name ILIKE '%' || $1 || '%')`
)
