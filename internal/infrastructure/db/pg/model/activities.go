package model

type Activity struct {
	ID                  *int64   `db:"id" json:"id"`
	IDOwner             int      `db:"id_owner" json:"id_owner"`
	INN                 string   `db:"inn" json:"inn"`
	ActivityTypeID      *int     `db:"activity_type_id" json:"activity_type_id,omitempty"`
	ActivityTypeName    *string  `db:"activity_type_name" json:"activity_type_name,omitempty"`
	CustomActivityID    *int     `db:"custom_activity_id" json:"custom_activity_id,omitempty"`
	CustomActivityName  *string  `db:"custom_activity_name" json:"custom_activity_name,omitempty"`
	VisitorCategory     string   `db:"visitor_category" json:"visitor_category"`
	CostSharePercent    *float64 `db:"cost_share_percent" json:"cost_share_percent,omitempty"`
	RevenueAmount       *float64 `db:"revenue_amount" json:"revenue_amount,omitempty"`
	TotalCount          *int64   `db:"total_count" json:"total_count,omitempty"`
	StateTaskCount      *int64   `db:"state_task_count" json:"state_task_count,omitempty"`
	RevenueActivityCount *int64  `db:"revenue_activity_count" json:"revenue_activity_count,omitempty"`
	Year                int16    `db:"year" json:"year"`
}
