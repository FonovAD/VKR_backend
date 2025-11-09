package handler

import (
	activity "vkr/internal/presenter/http/handler/activities"
	"vkr/internal/presenter/http/handler/museum"
	"vkr/internal/presenter/http/handler/organisation"
	reportingform "vkr/internal/presenter/http/handler/reporting_form"
)

type AppHandler interface {
	organisation.OrganizationHandler
	museum.MuseumHandler
	activity.ActivityHandler
	reportingform.ReportingFormHandler
}
