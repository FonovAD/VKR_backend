package entity

import (
	value "vkr/internal/domain/value"
)

type FormID int

type ReportingForm struct {
	ID             FormID
	OrganizationID OrganizationID
	Year           int16

	GeneralInfo value.FormGeneralInfo
	MuseumData  value.MuseumDetails
	Labor       value.LaborData
	Finances    value.FinancialData
	Status      FormStatus
}
