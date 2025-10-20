package value

type LaborData struct {
	TotalStaffAnnual             float64
	ResearchStaffInternal        float64
	CoreOperationalStaffInternal float64
	AdminSupportStaffInternal    float64
	ResearchStaffExternal        float64
	CoreOperationalStaffExternal float64
	AdminSupportStaffExternal    float64

	FOT LaborFOT
}

type LaborFOT struct {
	TotalStaffAnnual             float64
	ResearchStaffInternal        float64
	CoreOperationalStaffInternal float64
	AdminSupportStaffInternal    float64
	ResearchStaffExternal        float64
	CoreOperationalStaffExternal float64
	AdminSupportStaffExternal    float64
}
