package activities

const (
	createActivityQuery = `
		INSERT INTO museum_activities (
			inn, activity_type_id, visitor_category, cost_share_percent,
			revenue_amount, total_count, state_task_count, revenue_activity_count, year
		) VALUES (
			:inn, :activity_type_id, :visitor_category, :cost_share_percent,
			:revenue_amount, :total_count, :state_task_count, :revenue_activity_count, :year
		)`

	getActivityByINNQuery = `
		SELECT 
			inn, activity_type_id, visitor_category, cost_share_percent,
			revenue_amount, total_count, state_task_count, revenue_activity_count, year
		FROM museum_activities
		WHERE inn = $1
		ORDER BY activity_type_id, visitor_category`

	updateActivityQuery = `
		UPDATE museum_activities SET
			cost_share_percent = :cost_share_percent,
			revenue_amount = :revenue_amount,
			total_count = :total_count,
			state_task_count = :state_task_count,
			revenue_activity_count = :revenue_activity_count,
			year = :year
		WHERE inn = :inn
			AND activity_type_id = :activity_type_id
			AND visitor_category = :visitor_category`

	deleteActivityQuery = `
		DELETE FROM museum_activities
		WHERE inn = $1
			AND activity_type_id = $2
			AND visitor_category = $3
			AND year = $4`

	listActivitiesQuery = `
		SELECT 
			inn, activity_type_id, visitor_category, cost_share_percent,
			revenue_amount, total_count, state_task_count, revenue_activity_count, year
		FROM museum_activities
		ORDER BY inn, activity_type_id, visitor_category`
)
