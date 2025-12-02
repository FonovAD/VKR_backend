package museum

import (
	"vkr/internal/domain/entity"
	"vkr/internal/infrastructure/db/pg/model"
)

func (r *museumRepository) toDB(m *entity.Museum) *model.Museum {
	return &model.Museum{
		OrganizationID:                   int(m.IdOwner),
		Id:                               int(m.Id),
		INN:                              m.INN,
		KPP:                              m.KPP,
		Founder:                          m.Founder,
		MuseumActivityInCharter:          m.MuseumActivityInCharter,
		Name:                             m.Name,
		MuseumLegalStatus:                m.MuseumLegalStatus,
		IsMemorialReserveMuseum:          m.IsMemorialReserveMuseum,
		IsHistoricalMemorialReserve:      m.IsHistoricalMemorialReserve,
		IsArtMuseum:                      m.IsArtMuseum,
		IsMuseumReserve:                  m.IsMuseumReserve,
		IsEstateMuseum:                   m.IsEstateMuseum,
		IsPalaceParkEnsemble:             m.IsPalaceParkEnsemble,
		IsHistoricalArchitecturalReserve: m.IsHistoricalArchitecturalReserve,
		AnnualVisitorCapacity:            m.AnnualVisitorCapacity,
		InternalVisitorsCount:            m.InternalVisitorsCount,
		ExternalVisitorsCount:            m.ExternalVisitorsCount,
		IsValuableCulturalHeritage:       m.IsValuableCulturalHeritage,
		ValuableMuseumItemsCount:         m.ValuableMuseumItemsCount,
	}
}

func (r *museumRepository) toEntity(dbModel *model.Museum) *entity.Museum {
	return &entity.Museum{
		IdOwner:                          entity.OrganizationID(dbModel.OrganizationID),
		Id:                               entity.MuseumID(dbModel.Id),
		INN:                              dbModel.INN,
		KPP:                              dbModel.KPP,
		Founder:                          dbModel.Founder,
		MuseumActivityInCharter:          dbModel.MuseumActivityInCharter,
		Name:                             dbModel.Name,
		MuseumLegalStatus:                dbModel.MuseumLegalStatus,
		IsMemorialReserveMuseum:          dbModel.IsMemorialReserveMuseum,
		IsHistoricalMemorialReserve:      dbModel.IsHistoricalMemorialReserve,
		IsArtMuseum:                      dbModel.IsArtMuseum,
		IsMuseumReserve:                  dbModel.IsMuseumReserve,
		IsEstateMuseum:                   dbModel.IsEstateMuseum,
		IsPalaceParkEnsemble:             dbModel.IsPalaceParkEnsemble,
		IsHistoricalArchitecturalReserve: dbModel.IsHistoricalArchitecturalReserve,
		AnnualVisitorCapacity:            dbModel.AnnualVisitorCapacity,
		InternalVisitorsCount:            dbModel.InternalVisitorsCount,
		ExternalVisitorsCount:            dbModel.ExternalVisitorsCount,
		IsValuableCulturalHeritage:       dbModel.IsValuableCulturalHeritage,
		ValuableMuseumItemsCount:         dbModel.ValuableMuseumItemsCount,
	}
}
