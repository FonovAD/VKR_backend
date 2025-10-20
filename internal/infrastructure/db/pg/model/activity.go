package model

type Activity struct {
	ID           int    `db:"id"`
	Name         string `db:"name"`
	LocationType string `db:"location_type"`
}
