package interactor

import (
	"github.com/jmoiron/sqlx"
	"vkr/internal/logger"
	"vkr/internal/presenter/http/handler"
	activity "vkr/internal/presenter/http/handler/activities"
	"vkr/internal/presenter/http/handler/museum"
	"vkr/internal/presenter/http/handler/organisation"
)

type Interactor interface {
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	conn   *sqlx.DB
	logger logger.Logger
}

func NewInteractor(conn *sqlx.DB, logger logger.Logger) Interactor {
	return &interactor{conn: conn, logger: logger}
}

type appHandler struct {
	organisation.OrganizationHandler
	museum.MuseumHandler
	activity.ActivityHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	return &appHandler{
		OrganizationHandler: i.NewOrganizationHandler(),
		MuseumHandler:       i.NewMuseumHandler(),
		ActivityHandler:     i.NewActivityHandler(),
	}
}
