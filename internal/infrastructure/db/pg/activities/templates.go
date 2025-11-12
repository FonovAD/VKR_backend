package activities

const (
	createActivityQuery = `
		INSERT INTO museum_activities (
			id_owner, activity_type_id, custom_activity_id, visitor_category, cost_share_percent,
			revenue_amount, total_count, state_task_count, revenue_activity_count, year
		) VALUES (
			:id_owner, :activity_type_id, :custom_activity_id, :visitor_category, :cost_share_percent,
			:revenue_amount, :total_count, :state_task_count, :revenue_activity_count, :year
		)
		RETURNING id`

	getActivityByINNQuery = `
		SELECT 
			ma.id,
			ma.id_owner,
			o.inn,
			ma.activity_type_id,
			at.name AS activity_type_name,
			ma.custom_activity_id,
			cat.name AS custom_activity_name,
			ma.visitor_category,
			ma.cost_share_percent,
			ma.revenue_amount,
			ma.total_count,
			ma.state_task_count,
			ma.revenue_activity_count,
			ma.year
		FROM museum_activities ma
		JOIN organization o ON o.id = ma.id_owner
		LEFT JOIN activity_types at ON at.id = ma.activity_type_id
		LEFT JOIN activity_custom_types cat ON cat.id = ma.custom_activity_id
		WHERE LOWER(o.inn) = LOWER($1)
		ORDER BY COALESCE(ma.activity_type_id, ma.custom_activity_id), ma.visitor_category`

	updateActivityQuery = `
		UPDATE museum_activities SET
			cost_share_percent = :cost_share_percent,
			revenue_amount = :revenue_amount,
			total_count = :total_count,
			state_task_count = :state_task_count,
			revenue_activity_count = :revenue_activity_count,
			year = :year
		WHERE id = :id`

	deleteActivityQuery = `
		DELETE FROM museum_activities
		WHERE id = $1`

	listActivitiesQuery = `
		SELECT 
			ma.id,
			ma.id_owner,
			o.inn,
			ma.activity_type_id,
			at.name AS activity_type_name,
			ma.custom_activity_id,
			cat.name AS custom_activity_name,
			ma.visitor_category,
			ma.cost_share_percent,
			ma.revenue_amount,
			ma.total_count,
			ma.state_task_count,
			ma.revenue_activity_count,
			ma.year
		FROM museum_activities ma
		JOIN organization o ON o.id = ma.id_owner
		LEFT JOIN activity_types at ON at.id = ma.activity_type_id
		LEFT JOIN activity_custom_types cat ON cat.id = ma.custom_activity_id
		ORDER BY o.inn, COALESCE(ma.activity_type_id, ma.custom_activity_id), ma.visitor_category
		LIMIT $1 OFFSET $2`

	countActivitiesQuery = `
		SELECT COUNT(*)
		FROM museum_activities`

	getActivityByMuseumIDQuery = `
		SELECT 
			ma.id,
			ma.id_owner,
			o.inn,
			ma.activity_type_id,
			at.name AS activity_type_name,
			ma.custom_activity_id,
			cat.name AS custom_activity_name,
			ma.visitor_category,
			ma.cost_share_percent,
			ma.revenue_amount,
			ma.total_count,
			ma.state_task_count,
			ma.revenue_activity_count,
			ma.year
		FROM museum_activities ma
		JOIN organization o ON o.id = ma.id_owner
		LEFT JOIN activity_types at ON at.id = ma.activity_type_id
		LEFT JOIN activity_custom_types cat ON cat.id = ma.custom_activity_id
		WHERE ma.id_owner = (SELECT id_owner FROM museum WHERE id = $1)
		ORDER BY COALESCE(ma.activity_type_id, ma.custom_activity_id), ma.visitor_category`

	selectOrganizationIDByINNQuery = `
		SELECT id
		FROM organization
		WHERE inn = $1`
)
