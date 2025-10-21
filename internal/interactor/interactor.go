package interactor

import (
	"github.com/jmoiron/sqlx"
	"vkr/internal/logger"
	"vkr/internal/presenter/http/handler"
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
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.OrganizationHandler = i.NewOrganizationHandler()
	return appHandler
}
