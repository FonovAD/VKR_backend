package model

type Organization struct {
	ID          int    `db:"id"`
	INN         string `db:"inn"`
	Name        string `db:"name"`
	ExistMuseum bool   `db:"exist_museum"`
}
