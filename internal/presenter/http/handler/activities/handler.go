package activity

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"vkr/internal/domain/entity"
	"vkr/internal/logger"
	"vkr/internal/usecase/activities"
)

type ActivityHandler interface {
	CreateActivity(ctx echo.Context) error
	GetActivitiesByINN(ctx echo.Context) error
	UpdateActivity(ctx echo.Context) error
	DeleteActivity(ctx echo.Context) error
	ListActivities(ctx echo.Context) error
}

type activityHandler struct {
	useCase activities.ActivityUseCase
	logger  logger.Logger
}

func NewActivityHandler(uc activities.ActivityUseCase, logger logger.Logger) ActivityHandler {
	return &activityHandler{useCase: uc, logger: logger}
}

func (h *activityHandler) CreateActivity(ctx echo.Context) error {
	var req CreateActivityDTO
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	if !isValidVisitorCategory(req.VisitorCategory) {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid visitor_category: must be 'internal', 'external' or 'total'"})
	}

	activityEntity := &entity.Activity{
		INN:                  req.INN,
		ActivityTypeID:       req.ActivityTypeID,
		VisitorCategory:      req.VisitorCategory,
		CostSharePercent:     req.CostSharePercent,
		RevenueAmount:        req.RevenueAmount,
		TotalCount:           req.TotalCount,
		StateTaskCount:       req.StateTaskCount,
		RevenueActivityCount: req.RevenueActivityCount,
		Year:                 req.Year,
	}

	_, err := h.useCase.Create(ctx.Request().Context(), activityEntity)
	if err != nil {
		h.logger.LogError("activityHandler - CreateActivity", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	return ctx.NoContent(http.StatusCreated)
}

func (h *activityHandler) GetActivitiesByINN(ctx echo.Context) error {
	inn := ctx.QueryParam("inn")
	if inn == "" {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "query parameter 'inn' is required"})
	}

	activities, err := h.useCase.GetByINN(ctx.Request().Context(), inn)
	if err != nil {
		h.logger.LogError("activityHandler - GetActivitiesByINN", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	return ctx.JSON(http.StatusOK, activities)
}

func (h *activityHandler) UpdateActivity(ctx echo.Context) error {
	var req UpdateActivityDTO
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	if !isValidVisitorCategory(req.VisitorCategory) {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid visitor_category"})
	}

	activityEntity := &entity.Activity{
		INN:                  req.INN,
		ActivityTypeID:       req.ActivityTypeID,
		VisitorCategory:      req.VisitorCategory,
		CostSharePercent:     req.CostSharePercent,
		RevenueAmount:        req.RevenueAmount,
		TotalCount:           req.TotalCount,
		StateTaskCount:       req.StateTaskCount,
		RevenueActivityCount: req.RevenueActivityCount,
		Year:                 req.Year,
	}

	_, err := h.useCase.Update(ctx.Request().Context(), activityEntity)
	if err != nil {
		h.logger.LogError("activityHandler - UpdateActivity", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *activityHandler) DeleteActivity(ctx echo.Context) error {
	// Получаем параметры из query — так удобнее для составного ключа
	inn := ctx.QueryParam("inn")
	typeIDStr := ctx.QueryParam("activity_type_id")
	category := ctx.QueryParam("visitor_category")
	yearStr := ctx.QueryParam("year")

	if inn == "" || typeIDStr == "" || category == "" || yearStr == "" {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{
			ErrorMsg: "query params 'inn', 'activity_type_id', 'visitor_category', 'year' are required",
		})
	}

	activityTypeID, err := strconv.Atoi(typeIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid activity_type_id"})
	}

	var visitorCategory entity.VisitorCategory
	switch category {
	case "internal", "external", "total":
		visitorCategory = entity.VisitorCategory(category)
	default:
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid visitor_category"})
	}

	year, err := strconv.ParseInt(yearStr, 10, 16)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid year"})
	}

	err = h.useCase.Delete(ctx.Request().Context(), inn, activityTypeID, visitorCategory, int16(year))
	if err != nil {
		h.logger.LogError("activityHandler - DeleteActivity", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (h *activityHandler) ListActivities(ctx echo.Context) error {
	activities, err := h.useCase.List(ctx.Request().Context())
	if err != nil {
		h.logger.LogError("activityHandler - ListActivities", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	return ctx.JSON(http.StatusOK, activities)
}

// Вспомогательная функция валидации категории
func isValidVisitorCategory(v VisitorCategory) bool {
	switch v {
	case "internal", "external", "total":
		return true
	default:
		return false
	}
}
