package handler

import (
	"vkr/internal/presenter/http/handler/museum"
	"vkr/internal/presenter/http/handler/organisation"
)

type AppHandler interface {
	organisation.OrganizationHandler
	museum.MuseumHandler
}
