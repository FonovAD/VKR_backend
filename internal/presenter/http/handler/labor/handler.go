package labor

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"vkr/internal/logger"
	laborUseCase "vkr/internal/usecase/labor"
)

type Handler interface {
	GetByOrganizationID(ctx echo.Context) error
	GetByOrganizationINN(ctx echo.Context) error
}

type handler struct {
	useCase laborUseCase.UseCase
	logger  logger.Logger
}

func NewHandler(useCase laborUseCase.UseCase, logger logger.Logger) Handler {
	return &handler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *handler) GetByOrganizationID(ctx echo.Context) error {
	orgIDStr := ctx.Param("org_id")
	orgID, err := strconv.Atoi(orgIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid organization id"})
	}

	data, err := h.useCase.GetByOrganizationID(ctx.Request().Context(), orgID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, NotFoundResponse{ErrorMsg: "labor data not found"})
		}
		h.logger.LogError("laborHandler - GetByOrganizationID", nil, err)
		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: "internal server error"})
	}

	return ctx.JSON(http.StatusOK, NewLaborResponse(data))
}

func (h *handler) GetByOrganizationINN(ctx echo.Context) error {
	inn := ctx.QueryParam("inn")
	if inn == "" {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "query parameter 'inn' is required"})
	}

	data, err := h.useCase.GetByOrganizationINN(ctx.Request().Context(), inn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, NotFoundResponse{ErrorMsg: "labor data not found"})
		}
		h.logger.LogError("laborHandler - GetByOrganizationINN", nil, err)
		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: "internal server error"})
	}

	return ctx.JSON(http.StatusOK, NewLaborResponse(data))
}
