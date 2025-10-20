package entity

type OrganizationID int

type Organization struct {
	ID          OrganizationID
	INN         string
	Name        string
	ExistMuseum bool
}
