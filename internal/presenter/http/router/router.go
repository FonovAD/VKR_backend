package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"vkr/internal/presenter/http/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.GET("/ping", Ping)

	// Organization routes
	orgRouter := e.Group("/api/v1/organization")
	orgRouter.POST("", h.CreateOrg)
	orgRouter.GET("/:id", h.GetOrgByID)
	orgRouter.PUT("/:id", h.UpdateOrg)
	orgRouter.DELETE("/:id", h.DeleteOrg)
	orgRouter.GET("/search/by-inn", h.FindOrgByINN)
	orgRouter.GET("", h.OrgList)

	// Museum routes
	museumRouter := e.Group("/api/v1/museum")
	museumRouter.POST("", h.CreateMuseum)
	museumRouter.GET("/:id", h.GetMuseumByID)
	museumRouter.PUT("/:id", h.UpdateMuseum)
	museumRouter.DELETE("/:id", h.DeleteMuseum)
	museumRouter.GET("/search/by-inn", h.FindMuseumByINN)      // ?inn=1234567890
	museumRouter.GET("/owner/:owner_id", h.FindMuseumsByOwner) // музеи по ID организации
	museumRouter.GET("", h.ListMuseums)
}

func Ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}
