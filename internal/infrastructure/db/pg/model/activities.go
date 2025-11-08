package model

type Activity struct {
	INN                  string   `db:"inn" json:"inn"`
	ActivityTypeID       int      `db:"activity_type_id" json:"activity_type_id"`
	VisitorCategory      string   `db:"visitor_category" json:"visitor_category"`
	CostSharePercent     *float64 `db:"cost_share_percent" json:"cost_share_percent"`
	RevenueAmount        *float64 `db:"revenue_amount" json:"revenue_amount"`
	TotalCount           *int64   `db:"total_count" json:"total_count"`
	StateTaskCount       *int64   `db:"state_task_count" json:"state_task_count"`
	RevenueActivityCount *int64   `db:"revenue_activity_count" json:"revenue_activity_count"`
	Year                 int16    `db:"year" json:"year"`
}
