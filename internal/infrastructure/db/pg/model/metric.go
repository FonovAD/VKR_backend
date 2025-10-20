package model

type Metric struct {
	ID           int    `db:"id"`
	Name         string `db:"name"`
	AudienceType string `db:"audience_type"`
}
