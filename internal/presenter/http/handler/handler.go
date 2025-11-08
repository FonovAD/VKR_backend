package handler

import (
	activity "vkr/internal/presenter/http/handler/activities"
	"vkr/internal/presenter/http/handler/museum"
	"vkr/internal/presenter/http/handler/organisation"
)

type AppHandler interface {
	organisation.OrganizationHandler
	museum.MuseumHandler
	activity.ActivityHandler
}
