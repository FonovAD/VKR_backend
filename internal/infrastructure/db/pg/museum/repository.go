package museum

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	museumRepo "vkr/internal/domain/museum/repository"
	"vkr/internal/infrastructure/db/pg/model"

	"github.com/jmoiron/sqlx"
	"vkr/internal/domain/entity"
	"vkr/internal/logger"
)

type museumRepository struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewMuseumRepository(db *sqlx.DB, logger logger.Logger) museumRepo.Repository {
	return &museumRepository{
		db:     db,
		logger: logger,
	}
}

func (r *museumRepository) Create(ctx context.Context, museum *entity.Museum) error {
	r.logger.LogInfo(fmt.Sprintf("museumRepository - Create - INN: %s", museum.INN), nil, nil)

	dbModel := r.toDB(museum)
	var id int
	err := r.db.QueryRowxContext(ctx, createMuseumQuery,
		dbModel.IDOwner,
		dbModel.INN,
		dbModel.KPP,
		dbModel.Founder,
		dbModel.MuseumActivityInCharter,
		dbModel.Name,
		dbModel.MuseumLegalStatus,
		dbModel.IsMemorialReserveMuseum,
		dbModel.IsHistoricalMemorialReserve,
		dbModel.IsArtMuseum,
		dbModel.IsMuseumReserve,
		dbModel.IsEstateMuseum,
		dbModel.IsPalaceParkEnsemble,
		dbModel.IsHistoricalArchitecturalReserve,
		dbModel.AnnualVisitorCapacity,
		dbModel.InternalVisitorsCount,
		dbModel.ExternalVisitorsCount,
		dbModel.IsValuableCulturalHeritage,
		dbModel.ValuableMuseumItemsCount,
	).Scan(&id)
	if err != nil {
		return err
	}
	museum.Id = entity.MuseumID(id)
	return nil
}

func (r *museumRepository) GetByID(ctx context.Context, id entity.MuseumID) (*entity.Museum, error) {
	r.logger.LogInfo(fmt.Sprintf("museumRepository - GetByID - %d", id), nil, nil)
	var dbModel model.Museum
	err := r.db.GetContext(ctx, &dbModel, getMuseumByIDQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		r.logger.LogError(fmt.Sprintf("museumRepository - GetByID - %d", id), nil, err)
		return nil, err
	}
	dbModel.INN = "" // TODO: исправить костыль
	return r.toEntity(&dbModel), nil
}

func (r *museumRepository) Update(ctx context.Context, museum *entity.Museum) error {
	r.logger.LogInfo(fmt.Sprintf("museumRepository - Update - ID: %d", museum.Id), nil, nil)
	dbModel := r.toDB(museum)
	_, err := r.db.NamedExecContext(ctx, updateMuseumQuery, dbModel)
	return err
}

func (r *museumRepository) Delete(ctx context.Context, id entity.MuseumID) error {
	r.logger.LogInfo(fmt.Sprintf("museumRepository - Delete - %d", id), nil, nil)
	_, err := r.db.ExecContext(ctx, deleteMuseumQuery, id)
	return err
}

func (r *museumRepository) FindByINN(ctx context.Context, inn string) (*entity.Museum, error) {
	r.logger.LogInfo(fmt.Sprintf("museumRepository - FindByINN - %s", inn), nil, nil)
	var dbModel model.Museum
	err := r.db.GetContext(ctx, &dbModel, findMuseumByINNQuery, inn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		r.logger.LogError(fmt.Sprintf("museumRepository - FindByINN - %s", inn), nil, err)
		return nil, err
	}
	return r.toEntity(&dbModel), nil
}

func (r *museumRepository) FindByOwner(ctx context.Context, ownerID entity.OrganizationID) ([]*entity.Museum, error) {
	r.logger.LogInfo(fmt.Sprintf("museumRepository - FindByOwner - %d", ownerID), nil, nil)
	var dbModels []model.Museum
	err := r.db.SelectContext(ctx, &dbModels, findMuseumsByOwnerQuery, ownerID)
	if err != nil {
		r.logger.LogError(fmt.Sprintf("museumRepository - FindByOwner - %d", ownerID), nil, err)
		return nil, err
	}

	entities := make([]*entity.Museum, len(dbModels))
	for i, m := range dbModels {
		entities[i] = r.toEntity(&m)
	}
	return entities, nil
}

func (r *museumRepository) List(ctx context.Context, params museumRepo.ListParams) (*museumRepo.ListResult, error) {
	r.logger.LogInfo("museumRepository - List", nil, nil)
	
	// Get total count
	var totalCount int64
	err := r.db.GetContext(ctx, &totalCount, countMuseumsQuery, params.Name, params.MuseumType)
	if err != nil {
		r.logger.LogError("museumRepository - List - Count", nil, err)
		return nil, err
	}

	// Get paginated results
	var dbModels []model.Museum
	err = r.db.SelectContext(ctx, &dbModels, listMuseumsQuery, params.Name, params.MuseumType, params.Limit, params.Offset)
	if err != nil {
		r.logger.LogError("museumRepository - List - Select", nil, err)
		return nil, err
	}

	entities := make([]*entity.Museum, len(dbModels))
	for i, m := range dbModels {
		entities[i] = r.toEntity(&m)
	}

	return &museumRepo.ListResult{
		Museums:    entities,
		TotalCount: totalCount,
	}, nil
}
