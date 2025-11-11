package model

type Museum struct {
	IDOwner                          int     `db:"id_owner"`
	Id                               int     `db:"id"`
	INN                              string  `db:"inn"`
	KPP                              *string `db:"kpp"`
	Founder                          *string `db:"founder"`
	MuseumActivityInCharter          bool    `db:"museum_activity_in_charter"`
	Name                             string  `db:"name"`
	MuseumLegalStatus                string  `db:"museum_legal_status"`
	IsMemorialReserveMuseum          bool    `db:"is_memorial_reserve_museum"`
	IsHistoricalMemorialReserve      bool    `db:"is_historical_memorial_reserve"`
	IsArtMuseum                      bool    `db:"is_art_museum"`
	IsMuseumReserve                  bool    `db:"is_museum_reserve"`
	IsEstateMuseum                   bool    `db:"is_estate_museum"`
	IsPalaceParkEnsemble             bool    `db:"is_palace_park_ensemble"`
	IsHistoricalArchitecturalReserve bool    `db:"is_historical_architectural_reserve"`
	AnnualVisitorCapacity            *int    `db:"annual_visitor_capacity"`
	InternalVisitorsCount            *int    `db:"internal_visitors_count"`
	ExternalVisitorsCount            *int    `db:"external_visitors_count"`
	IsValuableCulturalHeritage       bool    `db:"is_valuable_cultural_heritage"`
	ValuableMuseumItemsCount         *int    `db:"valuable_museum_items_count"`
}
