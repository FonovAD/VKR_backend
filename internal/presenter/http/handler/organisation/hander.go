package organisation

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"vkr/internal/logger"
	"vkr/internal/usecase/organisation"
)

type OrganizationHandler interface {
	CreateOrg(ctx echo.Context) error
	GetOrgByID(ctx echo.Context) error
	UpdateOrg(ctx echo.Context) error
	DeleteOrg(ctx echo.Context) error
	FindOrgByINN(ctx echo.Context) error
	OrgList(ctx echo.Context) error
}

type organizationHandler struct {
	useCase organisation.UseCase
	logger  logger.Logger
}

func NewOrganizationHandler(uc organisation.UseCase, logger logger.Logger) OrganizationHandler {
	return &organizationHandler{useCase: uc, logger: logger}
}

func (h *organizationHandler) CreateOrg(ctx echo.Context) error {
	var req CreateOrganizationDTO
	if err := ctx.Bind(&req); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	org, err := h.useCase.Create(
		ctx.Request().Context(),
		req.INN,
		req.Name,
		req.ExistMuseum,
	)
	if err != nil {
		h.logger.LogError("organizationHandler - CreateOrg", err, nil)
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusCreated, org)
}

func (h *organizationHandler) GetOrgByID(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := parseOrganizationID(idParam)
	if err != nil {
		h.logger.LogError("GetByID", err, "invalid ID format")
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid ID"})
	}

	org, err := h.useCase.GetByID(ctx.Request().Context(), id)
	if err != nil {
		h.logger.LogError("GetByID", err, "failed to get organization")
		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: "internal error"})
	}
	if org == nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "organization not found"})
	}

	return ctx.JSON(http.StatusOK, org)
}

func (h *organizationHandler) UpdateOrg(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := parseOrganizationID(idParam)
	if err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid organization ID"})
	}

	var req UpdateOrganizationDTO
	if err := ctx.Bind(&req); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	org, err := h.useCase.Update(
		ctx.Request().Context(),
		id,
		req.INN,
		req.Name,
		req.ExistMuseum,
	)
	if err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, org)
}

func (h *organizationHandler) DeleteOrg(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := parseOrganizationID(idParam)
	if err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid organization ID"})
	}

	err = h.useCase.Delete(ctx.Request().Context(), id)
	if err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (h *organizationHandler) FindOrgByINN(ctx echo.Context) error {
	inn := ctx.QueryParam("inn")
	if inn == "" {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "query parameter 'inn' is required"})
	}

	org, err := h.useCase.FindByINN(ctx.Request().Context(), inn)
	if err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}
	if org == nil {
		return ctx.JSON(http.StatusNotFound, BadRequestResponse{ErrorMsg: "organization with this INN not found"})
	}

	return ctx.JSON(http.StatusOK, org)
}

func (h *organizationHandler) OrgList(ctx echo.Context) error {
	orgs, err := h.useCase.List(ctx.Request().Context())
	if err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, orgs)
}

func parseOrganizationID(s string) (OrganizationID, error) {

	id, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return OrganizationID(id), nil
}
