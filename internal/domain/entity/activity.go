package entity

type ActivityId int

type Activity struct {
	ID           ActivityId
	Name         string
	LocationType LocationType // enum: internal, external, online
}
