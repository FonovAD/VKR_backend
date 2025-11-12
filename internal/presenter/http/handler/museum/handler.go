package museum

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"vkr/internal/domain/entity"
	"vkr/internal/logger"
	"vkr/internal/presenter/http/pagination"
	"vkr/internal/usecase/museum"
)

var (
	ErrInternalServer = &InternalServerErrorResponse{ErrorMsg: "internal server error"}
)

type MuseumHandler interface {
	CreateMuseum(ctx echo.Context) error
	GetMuseumByID(ctx echo.Context) error
	UpdateMuseum(ctx echo.Context) error
	DeleteMuseum(ctx echo.Context) error
	FindMuseumByINN(ctx echo.Context) error
	FindMuseumsByOwner(ctx echo.Context) error
	ListMuseums(ctx echo.Context) error
}

type museumHandler struct {
	useCase museum.UseCase
	logger  logger.Logger
}

func NewMuseumHandler(uc museum.UseCase, logger logger.Logger) MuseumHandler {
	return &museumHandler{useCase: uc, logger: logger}
}

func (h *museumHandler) CreateMuseum(ctx echo.Context) error {
	var req CreateMuseumDTO
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	museumEntity := &entity.Museum{
		IdOwner:                          req.IdOwner,
		INN:                              req.INN,
		KPP:                              req.KPP,
		Founder:                          req.Founder,
		MuseumActivityInCharter:          req.MuseumActivityInCharter,
		Name:                             req.Name,
		MuseumLegalStatus:                req.MuseumLegalStatus,
		IsMemorialReserveMuseum:          req.IsMemorialReserveMuseum,
		IsHistoricalMemorialReserve:      req.IsHistoricalMemorialReserve,
		IsArtMuseum:                      req.IsArtMuseum,
		IsMuseumReserve:                  req.IsMuseumReserve,
		IsEstateMuseum:                   req.IsEstateMuseum,
		IsPalaceParkEnsemble:             req.IsPalaceParkEnsemble,
		IsHistoricalArchitecturalReserve: req.IsHistoricalArchitecturalReserve,
		AnnualVisitorCapacity:            req.AnnualVisitorCapacity,
		InternalVisitorsCount:            req.InternalVisitorsCount,
		ExternalVisitorsCount:            req.ExternalVisitorsCount,
		IsValuableCulturalHeritage:       req.IsValuableCulturalHeritage,
		ValuableMuseumItemsCount:         req.ValuableMuseumItemsCount,
	}

	created, err := h.useCase.Create(ctx.Request().Context(), museumEntity)
	if err != nil {
		h.logger.LogError("museumHandler - CreateMuseum", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	return ctx.JSON(http.StatusCreated, created)
}

func (h *museumHandler) GetMuseumByID(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := parseMuseumID(idParam)
	if err != nil {
		h.logger.LogError("museumHandler - GetMuseumByID", err, "invalid ID format")
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid museum ID"})
	}

	m, err := h.useCase.GetByID(ctx.Request().Context(), id)
	if err != nil {
		h.logger.LogError("museumHandler - GetMuseumByID", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}
	if m == nil {
		return ctx.JSON(http.StatusNotFound, BadRequestResponse{ErrorMsg: "museum not found"})
	}

	return ctx.JSON(http.StatusOK, m)
}

func (h *museumHandler) UpdateMuseum(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := parseMuseumID(idParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid museum ID"})
	}

	var req UpdateMuseumDTO
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	if req.Id != id {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "ID in URL does not match ID in body"})
	}

	museumEntity := &entity.Museum{
		Id:                               id,
		IdOwner:                          req.IdOwner,
		INN:                              req.INN,
		KPP:                              req.KPP,
		Founder:                          req.Founder,
		MuseumActivityInCharter:          req.MuseumActivityInCharter,
		Name:                             req.Name,
		MuseumLegalStatus:                req.MuseumLegalStatus,
		IsMemorialReserveMuseum:          req.IsMemorialReserveMuseum,
		IsHistoricalMemorialReserve:      req.IsHistoricalMemorialReserve,
		IsArtMuseum:                      req.IsArtMuseum,
		IsMuseumReserve:                  req.IsMuseumReserve,
		IsEstateMuseum:                   req.IsEstateMuseum,
		IsPalaceParkEnsemble:             req.IsPalaceParkEnsemble,
		IsHistoricalArchitecturalReserve: req.IsHistoricalArchitecturalReserve,
		AnnualVisitorCapacity:            req.AnnualVisitorCapacity,
		InternalVisitorsCount:            req.InternalVisitorsCount,
		ExternalVisitorsCount:            req.ExternalVisitorsCount,
		IsValuableCulturalHeritage:       req.IsValuableCulturalHeritage,
		ValuableMuseumItemsCount:         req.ValuableMuseumItemsCount,
	}

	updated, err := h.useCase.Update(ctx.Request().Context(), museumEntity)
	if err != nil {
		h.logger.LogError("museumHandler - UpdateMuseum", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	return ctx.JSON(http.StatusOK, updated)
}

func (h *museumHandler) DeleteMuseum(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := parseMuseumID(idParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid museum ID"})
	}

	err = h.useCase.Delete(ctx.Request().Context(), id)
	if err != nil {
		h.logger.LogError("museumHandler - DeleteMuseum", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (h *museumHandler) FindMuseumByINN(ctx echo.Context) error {
	inn := ctx.QueryParam("inn")
	if inn == "" {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "query parameter 'inn' is required"})
	}

	m, err := h.useCase.FindByINN(ctx.Request().Context(), inn)
	if err != nil {
		h.logger.LogError("museumHandler - FindMuseumByINN", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}
	if m == nil {
		return ctx.JSON(http.StatusNotFound, BadRequestResponse{ErrorMsg: "museum with this INN not found"})
	}

	return ctx.JSON(http.StatusOK, m)
}

func (h *museumHandler) FindMuseumsByOwner(ctx echo.Context) error {
	ownerIDParam := ctx.Param("owner_id")
	ownerID, err := parseOrganizationID(ownerIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid owner ID"})
	}

	museums, err := h.useCase.FindByOwner(ctx.Request().Context(), ownerID)
	if err != nil {
		h.logger.LogError("museumHandler - FindMuseumsByOwner", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	return ctx.JSON(http.StatusOK, museums)
}

func (h *museumHandler) ListMuseums(ctx echo.Context) error {
	// Parse pagination parameters
	var paginationParams pagination.PaginationParams
	if err := ctx.Bind(&paginationParams); err != nil {
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: "invalid pagination parameters"})
	}
	paginationParams.ValidateAndSetDefaults()

	// Parse name filter
	var nameFilter *string
	if name := ctx.QueryParam("name"); name != "" {
		nameFilter = &name
	}

	// Parse museum type filter
	var museumTypeFilter *string
	if museumType := ctx.QueryParam("museum_type"); museumType != "" {
		museumTypeFilter = &museumType
	}

	// Get paginated list
	result, err := h.useCase.List(ctx.Request().Context(), nameFilter, museumTypeFilter, paginationParams.Limit(), paginationParams.Offset())
	if err != nil {
		h.logger.LogError("museumHandler - ListMuseums", nil, err)
		return ctx.JSON(http.StatusInternalServerError, ErrInternalServer)
	}

	// Build paginated response
	response := pagination.PaginatedResponse[*entity.Museum]{
		Data:       result.Museums,
		Page:       paginationParams.Page,
		PageSize:   paginationParams.PageSize,
		TotalCount: result.TotalCount,
		TotalPages: pagination.CalculateTotalPages(result.TotalCount, paginationParams.PageSize),
	}

	return ctx.JSON(http.StatusOK, response)
}

func parseMuseumID(s string) (MuseumID, error) {
	id, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return MuseumID(id), nil
}

func parseOrganizationID(s string) (OrganizationID, error) {
	id, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return OrganizationID(id), nil
}
