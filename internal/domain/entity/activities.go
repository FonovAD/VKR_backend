package entity

type Activity struct {
	INN                  string
	ActivityTypeID       int
	VisitorCategory      VisitorCategory
	CostSharePercent     *float64
	RevenueAmount        *float64
	TotalCount           *int64
	StateTaskCount       *int64
	RevenueActivityCount *int64
	Year                 int16
}

type VisitorCategory string

const (
	VisitorInternal VisitorCategory = "internal"
	VisitorExternal VisitorCategory = "external"
	VisitorTotal    VisitorCategory = "total"
)
