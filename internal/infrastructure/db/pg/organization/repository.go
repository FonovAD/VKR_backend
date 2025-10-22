package organization

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	organization "vkr/internal/domain/organization/repository"
	"vkr/internal/logger"

	"github.com/jmoiron/sqlx"
	"vkr/internal/domain/entity"
)

type OrganizationDB struct {
	ID          int    `db:"id"`
	INN         string `db:"inn"`
	Name        string `db:"name"`
	ExistMuseum bool   `db:"exist_museum"`
}

type organizationRepository struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewOrganizationRepository(db *sqlx.DB, logger logger.Logger) organization.Repository {
	return &organizationRepository{
		db:     db,
		logger: logger,
	}
}

func (r *organizationRepository) Create(ctx context.Context, org *entity.Organization) error {
	r.logger.LogInfo(fmt.Sprintf("organizationRepository - Create - %s", org.ID), nil, nil)
	dbModel := OrganizationDB{
		INN:         org.INN,
		Name:        org.Name,
		ExistMuseum: org.ExistMuseum,
	}
	fmt.Print(dbModel)
	var id int
	err := r.db.QueryRowxContext(ctx, createOrganizationQuery, dbModel.INN, dbModel.Name, dbModel.ExistMuseum).Scan(&id)
	if err != nil {
		return err
	}
	org.ID = entity.OrganizationID(id)
	return nil
}

func (r *organizationRepository) GetByID(ctx context.Context, id entity.OrganizationID) (*entity.Organization, error) {
	r.logger.LogInfo(fmt.Sprintf("organizationRepository - GetByID - %d", id), nil, nil)
	var dbModel OrganizationDB
	err := r.db.GetContext(ctx, &dbModel, getOrganizationByIDQuery, id)
	if err != nil {
		r.logger.LogInfo(fmt.Sprintf("organizationRepository - GetByID - %d", id), err, nil)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return r.toEntity(&dbModel), nil
}

func (r *organizationRepository) Update(ctx context.Context, org *entity.Organization) error {
	r.logger.LogInfo(fmt.Sprintf("organizationRepository - Update - %s", org.ID), nil, nil)
	dbModel := OrganizationDB{
		ID:          int(org.ID),
		INN:         org.INN,
		Name:        org.Name,
		ExistMuseum: org.ExistMuseum,
	}

	_, err := r.db.NamedExecContext(ctx, updateOrganizationQuery, dbModel)
	return err
}

func (r *organizationRepository) Delete(ctx context.Context, id entity.OrganizationID) error {
	r.logger.LogInfo(fmt.Sprintf("organizationRepository - Delete - %d", id), nil, nil)
	_, err := r.db.ExecContext(ctx, deleteOrganizationQuery, id)
	return err
}

func (r *organizationRepository) FindByINN(ctx context.Context, inn string) (*entity.Organization, error) {
	r.logger.LogInfo(fmt.Sprintf("organizationRepository - FindByINN - %s", inn), nil, nil)
	var dbModel OrganizationDB
	err := r.db.GetContext(ctx, &dbModel, findOrganizationByINNQuery, inn)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return r.toEntity(&dbModel), nil
}

func (r *organizationRepository) List(ctx context.Context) ([]*entity.Organization, error) {
	r.logger.LogInfo("organizationRepository - FindByINN", nil, nil)
	var dbModels []OrganizationDB
	err := r.db.SelectContext(ctx, &dbModels, listOrganizationsQuery)
	if err != nil {
		return nil, err
	}

	entities := make([]*entity.Organization, len(dbModels))
	for i, m := range dbModels {
		entities[i] = r.toEntity(&m)
	}
	return entities, nil
}

func (r *organizationRepository) toEntity(dbModel *OrganizationDB) *entity.Organization {
	return &entity.Organization{
		ID:          entity.OrganizationID(dbModel.ID),
		INN:         dbModel.INN,
		Name:        dbModel.Name,
		ExistMuseum: dbModel.ExistMuseum,
	}
}
