package labor

import (
	value "vkr/internal/domain/value"
)

type BadRequestResponse struct {
	ErrorMsg string `json:"error"`
}

type InternalServerErrorResponse struct {
	ErrorMsg string `json:"error"`
}

type NotFoundResponse struct {
	ErrorMsg string `json:"error"`
}

type LaborResponse struct {
	TotalStaffAnnual             float64       `json:"total_staff_annual"`
	ResearchStaffInternal        float64       `json:"research_staff_internal"`
	CoreOperationalStaffInternal float64       `json:"core_operational_staff_internal"`
	AdminSupportStaffInternal    float64       `json:"admin_support_staff_internal"`
	ResearchStaffExternal        float64       `json:"research_staff_external"`
	CoreOperationalStaffExternal float64       `json:"core_operational_staff_external"`
	AdminSupportStaffExternal    float64       `json:"admin_support_staff_external"`
	FOT                          LaborFotResponse `json:"fot"`
}

type LaborFotResponse struct {
	TotalStaffAnnual             float64 `json:"total_staff_annual"`
	ResearchStaffInternal        float64 `json:"research_staff_internal"`
	CoreOperationalStaffInternal float64 `json:"core_operational_staff_internal"`
	AdminSupportStaffInternal    float64 `json:"admin_support_staff_internal"`
	ResearchStaffExternal        float64 `json:"research_staff_external"`
	CoreOperationalStaffExternal float64 `json:"core_operational_staff_external"`
	AdminSupportStaffExternal    float64 `json:"admin_support_staff_external"`
}

func NewLaborResponse(data value.LaborData) LaborResponse {
	return LaborResponse{
		TotalStaffAnnual:             data.TotalStaffAnnual,
		ResearchStaffInternal:        data.ResearchStaffInternal,
		CoreOperationalStaffInternal: data.CoreOperationalStaffInternal,
		AdminSupportStaffInternal:    data.AdminSupportStaffInternal,
		ResearchStaffExternal:        data.ResearchStaffExternal,
		CoreOperationalStaffExternal: data.CoreOperationalStaffExternal,
		AdminSupportStaffExternal:    data.AdminSupportStaffExternal,
		FOT: LaborFotResponse{
			TotalStaffAnnual:             data.FOT.TotalStaffAnnual,
			ResearchStaffInternal:        data.FOT.ResearchStaffInternal,
			CoreOperationalStaffInternal: data.FOT.CoreOperationalStaffInternal,
			AdminSupportStaffInternal:    data.FOT.AdminSupportStaffInternal,
			ResearchStaffExternal:        data.FOT.ResearchStaffExternal,
			CoreOperationalStaffExternal: data.FOT.CoreOperationalStaffExternal,
			AdminSupportStaffExternal:    data.FOT.AdminSupportStaffExternal,
		},
	}
}

