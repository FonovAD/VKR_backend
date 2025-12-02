package entity

type Activity struct {
	ID                  *int64
	IDOwner             OrganizationID
	INN                 string
	ActivityTypeID      *int
	ActivityTypeName    *string
	VolumeIndicator     *string
	CustomActivityID    *int
	CustomActivityName  *string
	VisitorCategory     VisitorCategory
	CostSharePercent    *float64
	RevenueAmount       *float64
	TotalCount          *int64
	StateTaskCount      *int64
	RevenueActivityCount *int64
	Year                int16
}

type VisitorCategory string

const (
	VisitorInternal VisitorCategory = "internal"
	VisitorExternal VisitorCategory = "external"
)
