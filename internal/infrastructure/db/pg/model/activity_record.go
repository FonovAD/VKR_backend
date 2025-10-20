package model

type ActivityRecord struct {
	ID         int     `db:"id"`
	ActivityID int     `db:"activity_id"`
	MetricID   int     `db:"metric_id"`
	CostShare  float64 `db:"cost_share"`
}
