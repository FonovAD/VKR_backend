package model

type LaborResources struct {
	MuseumID                     int      `db:"museum_id"`
	ID                           int      `db:"id"`
	TotalStaffAnnual             *float64 `db:"total_staff_annual"`
	ResearchStaffInternal        *float64 `db:"research_staff_internal"`
	CoreOperationalStaffInternal *float64 `db:"core_operational_staff_internal"`
	AdminSupportStaffInternal    *float64 `db:"admin_support_staff_internal"`
	ResearchStaffExternal        *float64 `db:"research_staff_external"`
	CoreOperationalStaffExternal *float64 `db:"core_operational_staff_external"`
	AdminSupportStaffExternal    *float64 `db:"admin_support_staff_external"`
}
