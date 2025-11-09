package reportingform

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"vkr/internal/domain/entity"
	"vkr/internal/logger"
	reportingUseCase "vkr/internal/usecase/reporting_form"
)

type ReportingFormHandler interface {
	CreateReportingForm(ctx echo.Context) error
	GetReportingFormByID(ctx echo.Context) error
	GetReportingFormByOrganizationAndYear(ctx echo.Context) error
	UpdateReportingForm(ctx echo.Context) error
	DeleteReportingForm(ctx echo.Context) error
	ListReportingFormsByOrganization(ctx echo.Context) error
	ListReportingFormsByYear(ctx echo.Context) error
	ListReportingForms(ctx echo.Context) error
}

type handler struct {
	useCase reportingUseCase.UseCase
	logger  logger.Logger
}

func NewReportingFormHandler(useCase reportingUseCase.UseCase, logger logger.Logger) ReportingFormHandler {
	return &handler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *handler) CreateReportingForm(ctx echo.Context) error {
	var req ReportingFormRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	if !isValidStatus(req.Status) {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid status"})
	}

	form, err := req.ToEntity()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	created, err := h.useCase.Create(ctx.Request().Context(), form)
	if err != nil {
		h.logger.LogError("reportingFormHandler - CreateReportingForm", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	return ctx.JSON(http.StatusCreated, created)
}

func (h *handler) GetReportingFormByID(ctx echo.Context) error {
	id, err := parseFormID(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid form id"})
	}

	form, err := h.useCase.GetByID(ctx.Request().Context(), id)
	if err != nil {
		h.logger.LogError("reportingFormHandler - GetReportingFormByID", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	if form == nil {
		return ctx.JSON(http.StatusNotFound, NotFoundResponse{ErrorMsg: "reporting form not found"})
	}

	return ctx.JSON(http.StatusOK, form)
}

func (h *handler) GetReportingFormByOrganizationAndYear(ctx echo.Context) error {
	orgID, err := parseOrganizationID(ctx.Param("org_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid organization id"})
	}

	year, err := parseYear(ctx.Param("year"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid year"})
	}

	form, err := h.useCase.GetByOrganizationAndYear(ctx.Request().Context(), orgID, year)
	if err != nil {
		h.logger.LogError("reportingFormHandler - GetReportingFormByOrganizationAndYear", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	if form == nil {
		return ctx.JSON(http.StatusNotFound, NotFoundResponse{ErrorMsg: "reporting form not found"})
	}

	return ctx.JSON(http.StatusOK, form)
}

func (h *handler) UpdateReportingForm(ctx echo.Context) error {
	id, err := parseFormID(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid form id"})
	}

	var req ReportingFormUpdateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	if req.ID != id {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "id mismatch between path and payload"})
	}

	if !isValidStatus(req.Status) {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid status"})
	}

	form, err := req.ReportingFormRequest.ToEntity()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}
	form.ID = id

	updated, err := h.useCase.Update(ctx.Request().Context(), form)
	if err != nil {
		h.logger.LogError("reportingFormHandler - UpdateReportingForm", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	return ctx.JSON(http.StatusOK, updated)
}

func (h *handler) DeleteReportingForm(ctx echo.Context) error {
	id, err := parseFormID(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid form id"})
	}

	if err := h.useCase.Delete(ctx.Request().Context(), id); err != nil {
		h.logger.LogError("reportingFormHandler - DeleteReportingForm", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (h *handler) ListReportingFormsByOrganization(ctx echo.Context) error {
	orgID, err := parseOrganizationID(ctx.Param("org_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid organization id"})
	}

	forms, err := h.useCase.ListByOrganization(ctx.Request().Context(), orgID)
	if err != nil {
		h.logger.LogError("reportingFormHandler - ListReportingFormsByOrganization", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	return ctx.JSON(http.StatusOK, forms)
}

func (h *handler) ListReportingFormsByYear(ctx echo.Context) error {
	year, err := parseYear(ctx.Param("year"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid year"})
	}

	forms, err := h.useCase.ListByYear(ctx.Request().Context(), year)
	if err != nil {
		h.logger.LogError("reportingFormHandler - ListReportingFormsByYear", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	return ctx.JSON(http.StatusOK, forms)
}

func (h *handler) ListReportingForms(ctx echo.Context) error {
	forms, err := h.useCase.List(ctx.Request().Context())
	if err != nil {
		h.logger.LogError("reportingFormHandler - ListReportingForms", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	return ctx.JSON(http.StatusOK, forms)
}

func isValidStatus(status entity.FormStatus) bool {
	switch status {
	case entity.FormStatusDraft, entity.FormStatusSubmitted, entity.FormStatusApproved:
		return true
	default:
		return false
	}
}

func parseFormID(value string) (entity.FormID, error) {
	id, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return entity.FormID(id), nil
}

func parseOrganizationID(value string) (entity.OrganizationID, error) {
	id, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return entity.OrganizationID(id), nil
}

func parseYear(value string) (int16, error) {
	parsed, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(parsed), nil
}
