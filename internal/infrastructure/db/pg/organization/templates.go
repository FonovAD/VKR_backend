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
		WHERE inn = $1`

	listOrganizationsQuery = `
		SELECT id, inn, name, exist_museum
		FROM organization
		ORDER BY id`
)
