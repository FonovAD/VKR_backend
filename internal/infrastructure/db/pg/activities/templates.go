package activities

const (
	createActivityQuery = `
		INSERT INTO museum_activities (
			museum_id, activity_type_id, volume_indicator_id, visitor_category, cost_share_percent,
			revenue_amount, total_count, state_task_count, revenue_activity_count, year
		) VALUES (
			:museum_id, :activity_type_id, :volume_indicator_id, :visitor_category, :cost_share_percent,
			:revenue_amount, :total_count, :state_task_count, :revenue_activity_count, :year
		)
		RETURNING id`

	getActivityByINNQuery = `
		SELECT 
			ma.id,
			ma.museum_id,
			ma.activity_type_id,
			at.name AS activity_type_name,
			ma.volume_indicator_id,
			vi.name AS volume_indicator_name,
			ma.visitor_category,
			ma.cost_share_percent,
			ma.revenue_amount,
			ma.total_count,
			ma.state_task_count,
			ma.revenue_activity_count,
			ma.year
		FROM museum_activities ma
		JOIN museum m ON m.id = ma.museum_id
		JOIN organization o ON o.id = m.organization_id
		LEFT JOIN activity_types at ON at.id = ma.activity_type_id
		LEFT JOIN volume_indicators vi ON vi.id = ma.volume_indicator_id
		WHERE LOWER(o.inn) = LOWER($1)
		ORDER BY ma.activity_type_id, ma.visitor_category`

	updateActivityQuery = `
		UPDATE museum_activities SET
			activity_type_id = :activity_type_id,
			volume_indicator_id = :volume_indicator_id,
			visitor_category = :visitor_category,
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
			ma.museum_id,
			ma.activity_type_id,
			at.name AS activity_type_name,
			ma.volume_indicator_id,
			vi.name AS volume_indicator_name,
			ma.visitor_category,
			ma.cost_share_percent,
			ma.revenue_amount,
			ma.total_count,
			ma.state_task_count,
			ma.revenue_activity_count,
			ma.year
		FROM museum_activities ma
		JOIN museum m ON m.id = ma.museum_id
		JOIN organization o ON o.id = m.organization_id
		LEFT JOIN activity_types at ON at.id = ma.activity_type_id
		LEFT JOIN volume_indicators vi ON vi.id = ma.volume_indicator_id
		ORDER BY o.inn, ma.activity_type_id, ma.visitor_category
		LIMIT $1 OFFSET $2`

	countActivitiesQuery = `
		SELECT COUNT(*)
		FROM museum_activities`

	getActivityByMuseumIDQuery = `
		SELECT 
			ma.id,
			ma.museum_id,
			ma.activity_type_id,
			at.name AS activity_type_name,
			ma.volume_indicator_id,
			vi.name AS volume_indicator_name,
			ma.visitor_category,
			ma.cost_share_percent,
			ma.revenue_amount,
			ma.total_count,
			ma.state_task_count,
			ma.revenue_activity_count,
			ma.year
		FROM museum_activities ma
		LEFT JOIN activity_types at ON at.id = ma.activity_type_id
		LEFT JOIN volume_indicators vi ON vi.id = ma.volume_indicator_id
		WHERE ma.museum_id = $1
		ORDER BY ma.activity_type_id, ma.visitor_category`

	selectOrganizationIDByINNQuery = `
		SELECT id
		FROM organization
		WHERE inn = $1`

	getActivityByOrganizationIDAndYearQuery = `
		SELECT 
			ma.id,
			ma.museum_id,
			ma.activity_type_id,
			at.name AS activity_type_name,
			ma.volume_indicator_id,
			vi.name AS volume_indicator_name,
			ma.visitor_category,
			ma.cost_share_percent,
			ma.revenue_amount,
			ma.total_count,
			ma.state_task_count,
			ma.revenue_activity_count,
			ma.year
		FROM museum_activities ma
		JOIN museum m ON m.id = ma.museum_id
		LEFT JOIN activity_types at ON at.id = ma.activity_type_id
		LEFT JOIN volume_indicators vi ON vi.id = ma.volume_indicator_id
		WHERE m.organization_id = $1 AND ma.year = $2
		ORDER BY ma.activity_type_id, ma.visitor_category`
)
