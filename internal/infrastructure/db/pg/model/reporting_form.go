package model

type ReportingForm struct {
	ID             int    `db:"id"`
	OrganizationID int    `db:"organization_id"`
	Year           int16  `db:"year"`
	Status         string `db:"status"`
}
