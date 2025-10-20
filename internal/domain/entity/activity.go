package entity

type Activity struct {
	ID           int
	Name         string
	LocationType LocationType // enum: internal, external, online
}
