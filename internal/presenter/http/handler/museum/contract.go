package museum

import "vkr/internal/domain/entity"

type BadRequestResponse struct {
	ErrorMsg string `json:"error"`
}

type InternalServerErrorResponse struct {
	ErrorMsg string `json:"error"`
}

type MuseumID = entity.MuseumID
type OrganizationID = entity.OrganizationID

type CreateMuseumDTO struct {
	IdOwner                          OrganizationID `json:"id_owner"`
	INN                              string         `json:"inn"`
	KPP                              *string        `json:"kpp,omitempty"`
	Founder                          *string        `json:"founder,omitempty"`
	MuseumActivityInCharter          bool           `json:"museum_activity_in_charter"`
	Name                             string         `json:"name"`
	MuseumLegalStatus                string         `json:"museum_legal_status"`
	IsMemorialReserveMuseum          bool           `json:"is_memorial_reserve_museum"`
	IsHistoricalMemorialReserve      bool           `json:"is_historical_memorial_reserve"`
	IsArtMuseum                      bool           `json:"is_art_museum"`
	IsMuseumReserve                  bool           `json:"is_museum_reserve"`
	IsEstateMuseum                   bool           `json:"is_estate_museum"`
	IsPalaceParkEnsemble             bool           `json:"is_palace_park_ensemble"`
	IsHistoricalArchitecturalReserve bool           `json:"is_historical_architectural_reserve"`
	AnnualVisitorCapacity            *int           `json:"annual_visitor_capacity,omitempty"`
	InternalVisitorsCount            *int           `json:"internal_visitors_count,omitempty"`
	ExternalVisitorsCount            *int           `json:"external_visitors_count,omitempty"`
	IsValuableCulturalHeritage       bool           `json:"is_valuable_cultural_heritage"`
	ValuableMuseumItemsCount         *int           `json:"valuable_museum_items_count,omitempty"`
}

// DTO для обновления музея (аналогичен Create, но с ID)
type UpdateMuseumDTO struct {
	Id                               MuseumID       `json:"id"`
	IdOwner                          OrganizationID `json:"id_owner"`
	INN                              string         `json:"inn"`
	KPP                              *string        `json:"kpp,omitempty"`
	Founder                          *string        `json:"founder,omitempty"`
	MuseumActivityInCharter          bool           `json:"museum_activity_in_charter"`
	Name                             string         `json:"name"`
	MuseumLegalStatus                string         `json:"museum_legal_status"`
	IsMemorialReserveMuseum          bool           `json:"is_memorial_reserve_museum"`
	IsHistoricalMemorialReserve      bool           `json:"is_historical_memorial_reserve"`
	IsArtMuseum                      bool           `json:"is_art_museum"`
	IsMuseumReserve                  bool           `json:"is_museum_reserve"`
	IsEstateMuseum                   bool           `json:"is_estate_museum"`
	IsPalaceParkEnsemble             bool           `json:"is_palace_park_ensemble"`
	IsHistoricalArchitecturalReserve bool           `json:"is_historical_architectural_reserve"`
	AnnualVisitorCapacity            *int           `json:"annual_visitor_capacity,omitempty"`
	InternalVisitorsCount            *int           `json:"internal_visitors_count,omitempty"`
	ExternalVisitorsCount            *int           `json:"external_visitors_count,omitempty"`
	IsValuableCulturalHeritage       bool           `json:"is_valuable_cultural_heritage"`
	ValuableMuseumItemsCount         *int           `json:"valuable_museum_items_count,omitempty"`
}
