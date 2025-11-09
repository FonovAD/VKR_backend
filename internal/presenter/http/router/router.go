package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"vkr/internal/presenter/http/handler"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.GET("/ping", Ping)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	orgRouter := e.Group("/api/v1/organization")
	orgRouter.POST("", h.CreateOrg)
	orgRouter.GET("/:id", h.GetOrgByID)
	orgRouter.PUT("/:id", h.UpdateOrg)
	orgRouter.DELETE("/:id", h.DeleteOrg)
	orgRouter.GET("/search/by-inn", h.FindOrgByINN)
	orgRouter.GET("", h.OrgList)

	museumRouter := e.Group("/api/v1/museum")
	museumRouter.POST("", h.CreateMuseum)
	museumRouter.GET("/:id", h.GetMuseumByID)
	museumRouter.PUT("/:id", h.UpdateMuseum)
	museumRouter.DELETE("/:id", h.DeleteMuseum)
	museumRouter.GET("/search/by-inn", h.FindMuseumByINN)
	museumRouter.GET("/owner/:owner_id", h.FindMuseumsByOwner)
	museumRouter.GET("", h.ListMuseums)

	activityRouter := e.Group("/api/v1/activity")
	activityRouter.POST("", h.CreateActivity)
	activityRouter.GET("/search/by-inn", h.GetActivitiesByINN) // получение по INN
	activityRouter.PUT("", h.UpdateActivity)
	activityRouter.DELETE("", h.DeleteActivity) // параметры в query
	activityRouter.GET("", h.ListActivities)

	formRouter := e.Group("/api/v1/reporting-form")
	formRouter.GET("/organization/:org_id/year/:year", h.GetReportingFormByOrganizationAndYear)
	formRouter.GET("/organization/:org_id", h.ListReportingFormsByOrganization)
	formRouter.GET("/year/:year", h.ListReportingFormsByYear)
	formRouter.POST("", h.CreateReportingForm)
	formRouter.GET("/:id", h.GetReportingFormByID)
	formRouter.PUT("/:id", h.UpdateReportingForm)
	formRouter.DELETE("/:id", h.DeleteReportingForm)
	formRouter.GET("", h.ListReportingForms)
}

func Ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}
