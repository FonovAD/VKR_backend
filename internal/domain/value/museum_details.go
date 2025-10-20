package value

type MuseumDetails struct {
	INN                              string
	KPP                              *string
	Founder                          *string
	MuseumActivityInCharter          bool
	Name                             string
	MuseumLegalStatus                *string
	IsMemorialReserveMuseum          bool
	IsHistoricalMemorialReserve      bool
	IsArtMuseum                      bool
	IsMuseumReserve                  bool
	IsEstateMuseum                   bool
	IsPalaceParkEnsemble             bool
	IsHistoricalArchitecturalReserve bool
	AnnualVisitorCapacity            *int
	InternalVisitorsCount            *int
	ExternalVisitorsCount            *int
	IsValuableCulturalHeritage       bool
	ValuableMuseumItemsCount         *int
}
