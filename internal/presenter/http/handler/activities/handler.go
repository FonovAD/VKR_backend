package activity

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"vkr/internal/domain/entity"
	"vkr/internal/logger"
	"vkr/internal/presenter/http/pagination"
	"vkr/internal/usecase/activities"
)

type ActivityHandler interface {
	CreateActivity(ctx echo.Context) error
	GetActivitiesByINN(ctx echo.Context) error
	GetActivitiesByMuseumID(ctx echo.Context) error
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
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid visitor_category: must be 'internal' or 'external'"})
	}

	// Проверяем, что указан либо ActivityTypeID, либо CustomActivityID
	if (req.ActivityTypeID == nil && req.CustomActivityID == nil) || 
		(req.ActivityTypeID != nil && req.CustomActivityID != nil) {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{
			ErrorMsg: "exactly one of 'activity_type_id' or 'custom_activity_id' must be provided",
		})
	}

	activityEntity := &entity.Activity{
		INN:                  req.INN,
		ActivityTypeID:       req.ActivityTypeID,
		CustomActivityID:     req.CustomActivityID,
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

	responses := make([]ActivityResponse, len(activities))
	for i, activity := range activities {
		responses[i] = NewActivityResponse(activity)
	}

	return ctx.JSON(http.StatusOK, responses)
}

func (h *activityHandler) GetActivitiesByMuseumID(ctx echo.Context) error {
	museumIDParam := ctx.Param("museum_id")
	museumID, err := strconv.Atoi(museumIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid museum_id"})
	}

	activities, err := h.useCase.GetByMuseumID(ctx.Request().Context(), entity.MuseumID(museumID))
	if err != nil {
		h.logger.LogError("activityHandler - GetActivitiesByMuseumID", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	responses := make([]ActivityResponse, len(activities))
	for i, activity := range activities {
		responses[i] = NewActivityResponse(activity)
	}

	return ctx.JSON(http.StatusOK, responses)
}

func (h *activityHandler) UpdateActivity(ctx echo.Context) error {
	var req UpdateActivityDTO
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	if req.ID == nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "id is required"})
	}

	// Получаем текущую активность для получения остальных полей
	// Используем большой лимит для получения всех активностей (можно оптимизировать позже, добавив GetByID)
	result, err := h.useCase.List(ctx.Request().Context(), 10000, 0)
	if err != nil {
		h.logger.LogError("activityHandler - UpdateActivity - List", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	var currentActivity *entity.Activity
	for i := range result.Activities {
		if result.Activities[i].ID != nil && *result.Activities[i].ID == *req.ID {
			currentActivity = &result.Activities[i]
			break
		}
	}

	if currentActivity == nil {
		return ctx.JSON(http.StatusNotFound, BadRequestResponse{ErrorMsg: "activity not found"})
	}

	// Обновляем только переданные поля
	activityEntity := &entity.Activity{
		ID:                   req.ID,
		IDOwner:              currentActivity.IDOwner,
		INN:                  currentActivity.INN,
		ActivityTypeID:       currentActivity.ActivityTypeID,
		CustomActivityID:     currentActivity.CustomActivityID,
		VisitorCategory:      currentActivity.VisitorCategory,
		CostSharePercent:     req.CostSharePercent,
		RevenueAmount:        req.RevenueAmount,
		TotalCount:           req.TotalCount,
		StateTaskCount:       req.StateTaskCount,
		RevenueActivityCount: req.RevenueActivityCount,
		Year:                 req.Year,
	}

	_, err = h.useCase.Update(ctx.Request().Context(), activityEntity)
	if err != nil {
		h.logger.LogError("activityHandler - UpdateActivity", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *activityHandler) DeleteActivity(ctx echo.Context) error {
	idStr := ctx.QueryParam("id")
	if idStr == "" {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{
			ErrorMsg: "query param 'id' is required",
		})
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid id"})
	}

	err = h.useCase.DeleteByID(ctx.Request().Context(), id)
	if err != nil {
		h.logger.LogError("activityHandler - DeleteActivity", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (h *activityHandler) ListActivities(ctx echo.Context) error {
	// Parse pagination parameters
	var paginationParams pagination.PaginationParams
	if err := ctx.Bind(&paginationParams); err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid pagination parameters"})
	}
	paginationParams.ValidateAndSetDefaults()

	// Get paginated list
	result, err := h.useCase.List(ctx.Request().Context(), paginationParams.Limit(), paginationParams.Offset())
	if err != nil {
		h.logger.LogError("activityHandler - ListActivities", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	// Convert to response format
	responses := make([]ActivityResponse, len(result.Activities))
	for i, activity := range result.Activities {
		responses[i] = NewActivityResponse(activity)
	}

	// Build paginated response
	response := pagination.PaginatedResponse[ActivityResponse]{
		Data:       responses,
		Page:       paginationParams.Page,
		PageSize:   paginationParams.PageSize,
		TotalCount: result.TotalCount,
		TotalPages: pagination.CalculateTotalPages(result.TotalCount, paginationParams.PageSize),
	}

	return ctx.JSON(http.StatusOK, response)
}

// Вспомогательная функция валидации категории
func isValidVisitorCategory(v VisitorCategory) bool {
	switch v {
	case "internal", "external":
		return true
	default:
		return false
	}
}
