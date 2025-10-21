package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"vkr/internal/presenter/http/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.GET("/ping", Ping)

	orgRouter := e.Group("/api/v1/organization")
	orgRouter.POST("", h.CreateOrg)
	orgRouter.GET("/:id", h.GetOrgByID)
	orgRouter.PUT("/:id", h.UpdateOrg)
	orgRouter.DELETE("/:id", h.DeleteOrg)
	orgRouter.GET("/search/by-inn", h.FindOrgByINN)
	orgRouter.GET("", h.OrgList)
}

func Ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}
