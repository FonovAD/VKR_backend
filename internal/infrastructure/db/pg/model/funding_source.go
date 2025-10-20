package model

type FundingSource struct {
	ID                    int     `db:"id"`
	RecordID              int     `db:"record_id"`
	Revenue               float64 `db:"revenue"`
	TotalVolume           *int    `db:"total_volume"`
	StateAssignmentVolume *int    `db:"state_assignment_volume"`
	IncomeActivityVolume  *int    `db:"income_activity_volume"`
}
