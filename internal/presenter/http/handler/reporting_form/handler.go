package reporting_form

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"vkr/internal/logger"
	reportingFormUseCase "vkr/internal/usecase/reporting_form"
)

type ReportingFormHandler interface {
	GetLaborForm(ctx echo.Context) error
	GetLaborFormByINN(ctx echo.Context) error
	GetActivitiesForm(ctx echo.Context) error
	GetActivitiesFormByINN(ctx echo.Context) error
	GetFinancialForm(ctx echo.Context) error
	GetFinancialFormByINN(ctx echo.Context) error
}

type reportingFormHandler struct {
	laborFormUseCase      reportingFormUseCase.LaborFormUseCase
	activitiesFormUseCase reportingFormUseCase.ActivitiesFormUseCase
	financialFormUseCase  reportingFormUseCase.FinancialFormUseCase
	logger                logger.Logger
}

func NewReportingFormHandler(
	laborFormUseCase reportingFormUseCase.LaborFormUseCase,
	activitiesFormUseCase reportingFormUseCase.ActivitiesFormUseCase,
	financialFormUseCase reportingFormUseCase.FinancialFormUseCase,
	logger logger.Logger,
) ReportingFormHandler {
	return &reportingFormHandler{
		laborFormUseCase:      laborFormUseCase,
		activitiesFormUseCase: activitiesFormUseCase,
		financialFormUseCase:  financialFormUseCase,
		logger:                logger,
	}
}

// GetLaborForm возвращает данные формы трудовых ресурсов по ID организации
func (h *reportingFormHandler) GetLaborForm(ctx echo.Context) error {
	orgIDParam := ctx.Param("org_id")
	orgID, err := strconv.Atoi(orgIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid organization id"})
	}

	formData, err := h.laborFormUseCase.GetLaborFormData(ctx.Request().Context(), orgID)
	if err != nil {
		h.logger.LogError("reportingFormHandler - GetLaborForm", nil, err)
		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: "internal server error"})
	}

	return ctx.JSON(http.StatusOK, formData)
}

// GetLaborFormByINN возвращает данные формы трудовых ресурсов по INN организации
func (h *reportingFormHandler) GetLaborFormByINN(ctx echo.Context) error {
	inn := ctx.QueryParam("inn")
	if inn == "" {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "query parameter 'inn' is required"})
	}

	formData, err := h.laborFormUseCase.GetLaborFormDataByINN(ctx.Request().Context(), inn)
	if err != nil {
		h.logger.LogError("reportingFormHandler - GetLaborFormByINN", nil, err)
		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: "internal server error"})
	}

	return ctx.JSON(http.StatusOK, formData)
}

// GetActivitiesForm возвращает данные формы объемов (активностей) по ID организации и году
func (h *reportingFormHandler) GetActivitiesForm(ctx echo.Context) error {
	orgIDParam := ctx.Param("org_id")
	orgID, err := strconv.Atoi(orgIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid organization id"})
	}

	yearParam := ctx.QueryParam("year")
	if yearParam == "" {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "query parameter 'year' is required"})
	}

	year, err := strconv.ParseInt(yearParam, 10, 16)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid year format"})
	}

	formData, err := h.activitiesFormUseCase.GetActivitiesFormData(ctx.Request().Context(), orgID, int16(year))
	if err != nil {
		h.logger.LogError("reportingFormHandler - GetActivitiesForm", nil, err)
		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: "internal server error"})
	}

	return ctx.JSON(http.StatusOK, formData)
}

// GetActivitiesFormByINN возвращает данные формы объемов (активностей) по INN организации и году
func (h *reportingFormHandler) GetActivitiesFormByINN(ctx echo.Context) error {
	inn := ctx.QueryParam("inn")
	if inn == "" {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "query parameter 'inn' is required"})
	}

	yearParam := ctx.QueryParam("year")
	if yearParam == "" {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "query parameter 'year' is required"})
	}

	year, err := strconv.ParseInt(yearParam, 10, 16)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid year format"})
	}

	formData, err := h.activitiesFormUseCase.GetActivitiesFormDataByINN(ctx.Request().Context(), inn, int16(year))
	if err != nil {
		h.logger.LogError("reportingFormHandler - GetActivitiesFormByINN", nil, err)
		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: "internal server error"})
	}

	return ctx.JSON(http.StatusOK, formData)
}

// GetFinancialForm возвращает данные формы поступлений и расходов по ID организации и году
func (h *reportingFormHandler) GetFinancialForm(ctx echo.Context) error {
	orgIDParam := ctx.Param("org_id")
	orgID, err := strconv.Atoi(orgIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid organization id"})
	}

	yearParam := ctx.QueryParam("year")
	if yearParam == "" {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "query parameter 'year' is required"})
	}

	year, err := strconv.ParseInt(yearParam, 10, 16)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid year format"})
	}

	formData, err := h.financialFormUseCase.GetFinancialFormData(ctx.Request().Context(), orgID, int16(year))
	if err != nil {
		h.logger.LogError("reportingFormHandler - GetFinancialForm", nil, err)
		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: "internal server error"})
	}

	return ctx.JSON(http.StatusOK, formData)
}

// GetFinancialFormByINN возвращает данные формы поступлений и расходов по INN организации и году
func (h *reportingFormHandler) GetFinancialFormByINN(ctx echo.Context) error {
	inn := ctx.QueryParam("inn")
	if inn == "" {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "query parameter 'inn' is required"})
	}

	yearParam := ctx.QueryParam("year")
	if yearParam == "" {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "query parameter 'year' is required"})
	}

	year, err := strconv.ParseInt(yearParam, 10, 16)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid year format"})
	}

	formData, err := h.financialFormUseCase.GetFinancialFormDataByINN(ctx.Request().Context(), inn, int16(year))
	if err != nil {
		h.logger.LogError("reportingFormHandler - GetFinancialFormByINN", nil, err)
		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: "internal server error"})
	}

	return ctx.JSON(http.StatusOK, formData)
}



