package activities

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	activity "vkr/internal/domain/activities/repository"
	"vkr/internal/domain/entity"
	"vkr/internal/infrastructure/db/pg/model"
	"vkr/internal/logger"

	"github.com/jmoiron/sqlx"
)

type activityRepository struct {
	db     *sqlx.DB
	logger logger.Logger
}


func NewActivityRepository(db *sqlx.DB, logger logger.Logger) activity.Repository {
	return &activityRepository{
		db:     db,
		logger: logger,
	}
}

func (r *activityRepository) Create(ctx context.Context, activity *entity.Activity) error {
	r.logger.LogInfo(fmt.Sprintf("activityRepository - Create - INN: %s, ActivityTypeID: %v, Category: %s",
		activity.INN, activity.ActivityTypeID, activity.VisitorCategory), nil, nil)

	// TODO: нужно получать museum_id по INN через organization -> museum
	// Пока возвращаем ошибку, так как нужно знать museum_id
	return fmt.Errorf("museum_id is required for creating activity. Need to resolve museum_id from INN first")
}

func (r *activityRepository) GetByINN(ctx context.Context, inn string) ([]entity.Activity, error) {
	r.logger.LogInfo(fmt.Sprintf("activityRepository - GetByINN - %s", inn), nil, nil)

	var dbModels []model.Activity
	err := r.db.SelectContext(ctx, &dbModels, getActivityByINNQuery, inn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []entity.Activity{}, nil
		}
		r.logger.LogError(fmt.Sprintf("activityRepository - GetByINN - %s", inn), nil, err)
		return nil, err
	}

	activities := make([]entity.Activity, len(dbModels))
	for i, m := range dbModels {
		activities[i] = r.toEntityActivity(&m)
	}
	return activities, nil
}

func (r *activityRepository) GetByMuseumID(ctx context.Context, museumID entity.MuseumID) ([]entity.Activity, error) {
	r.logger.LogInfo(fmt.Sprintf("activityRepository - GetByMuseumID - %d", museumID), nil, nil)

	var dbModels []model.Activity
	err := r.db.SelectContext(ctx, &dbModels, getActivityByMuseumIDQuery, museumID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []entity.Activity{}, nil
		}
		r.logger.LogError(fmt.Sprintf("activityRepository - GetByMuseumID - %d", museumID), nil, err)
		return nil, err
	}

	activities := make([]entity.Activity, len(dbModels))
	for i, m := range dbModels {
		activities[i] = r.toEntityActivity(&m)
	}
	return activities, nil
}

func (r *activityRepository) GetByOrganizationIDAndYear(ctx context.Context, organizationID int, year int16) ([]entity.Activity, error) {
	r.logger.LogInfo(fmt.Sprintf("activityRepository - GetByOrganizationIDAndYear - orgID: %d, year: %d", organizationID, year), nil, nil)

	var dbModels []model.Activity
	err := r.db.SelectContext(ctx, &dbModels, getActivityByOrganizationIDAndYearQuery, organizationID, year)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []entity.Activity{}, nil
		}
		r.logger.LogError(fmt.Sprintf("activityRepository - GetByOrganizationIDAndYear - orgID: %d, year: %d", organizationID, year), nil, err)
		return nil, err
	}

	activities := make([]entity.Activity, len(dbModels))
	for i, m := range dbModels {
		activities[i] = r.toEntityActivity(&m)
	}
	return activities, nil
}

func (r *activityRepository) Update(ctx context.Context, activity *entity.Activity) error {
	r.logger.LogInfo(fmt.Sprintf("activityRepository - Update - ID: %v", activity.ID), nil, nil)

	if activity.ID == nil {
		return fmt.Errorf("activity ID is required for update")
	}

	dbModel := &model.Activity{
		ID:                  activity.ID,
		ActivityTypeID:      activity.ActivityTypeID,
		VolumeIndicatorID:   nil, // TODO: нужно получать volume_indicator_id из VolumeIndicator
		VisitorCategory:     string(activity.VisitorCategory),
		CostSharePercent:    activity.CostSharePercent,
		RevenueAmount:       activity.RevenueAmount,
		TotalCount:          activity.TotalCount,
		StateTaskCount:      activity.StateTaskCount,
		RevenueActivityCount: activity.RevenueActivityCount,
		Year:                activity.Year,
	}
	_, err := r.db.NamedExecContext(ctx, updateActivityQuery, dbModel)
	return err
}

func (r *activityRepository) Delete(ctx context.Context, inn string, activityTypeID int, visitorCategory entity.VisitorCategory, year int16) error {
	// Этот метод устарел, так как теперь используется ID для удаления
	// Оставляем для обратной совместимости, но лучше использовать DeleteByID
	return fmt.Errorf("delete by composite key is deprecated, use DeleteByID instead")
}

func (r *activityRepository) DeleteByID(ctx context.Context, id int64) error {
	r.logger.LogInfo(fmt.Sprintf("activityRepository - DeleteByID - %d", id), nil, nil)
	_, err := r.db.ExecContext(ctx, deleteActivityQuery, id)
	return err
}

func (r *activityRepository) List(ctx context.Context, params activity.ListParams) (*activity.ListResult, error) {
	r.logger.LogInfo("activityRepository - List", nil, nil)

	// Get total count
	var totalCount int64
	err := r.db.GetContext(ctx, &totalCount, countActivitiesQuery)
	if err != nil {
		r.logger.LogError("activityRepository - List - Count", nil, err)
		return nil, err
	}

	// Get paginated results
	var dbModels []model.Activity
	err = r.db.SelectContext(ctx, &dbModels, listActivitiesQuery, params.Limit, params.Offset)
	if err != nil {
		r.logger.LogError("activityRepository - List - Select", nil, err)
		return nil, err
	}

	entities := make([]entity.Activity, len(dbModels))
	for i, m := range dbModels {
		entities[i] = r.toEntityActivity(&m)
	}

	return &activity.ListResult{
		Activities: entities,
		TotalCount: totalCount,
	}, nil
}


// Маппинг в доменную сущность
func (r *activityRepository) toEntityActivity(m *model.Activity) entity.Activity {
	return entity.Activity{
		ID:                   m.ID,
		IDOwner:              0, // TODO: нужно получать organization_id через museum
		INN:                  "", // TODO: нужно получать inn через museum
		ActivityTypeID:      m.ActivityTypeID,
		ActivityTypeName:     m.ActivityTypeName,
		VolumeIndicator:      m.VolumeIndicatorName,
		CustomActivityID:     nil,
		CustomActivityName:   nil,
		VisitorCategory:      entity.VisitorCategory(m.VisitorCategory),
		CostSharePercent:     m.CostSharePercent,
		RevenueAmount:        m.RevenueAmount,
		TotalCount:           m.TotalCount,
		StateTaskCount:       m.StateTaskCount,
		RevenueActivityCount: m.RevenueActivityCount,
		Year:                 m.Year,
	}
}

func (r *activityRepository) resolveOrganizationID(ctx context.Context, inn string) (int, error) {
	var id int
	if err := r.db.GetContext(ctx, &id, selectOrganizationIDByINNQuery, inn); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("organization with inn %s not found: %w", inn, err)
		}
		return 0, err
	}
	return id, nil
}
