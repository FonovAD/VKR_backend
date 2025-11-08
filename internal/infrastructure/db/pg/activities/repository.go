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
	r.logger.LogInfo(fmt.Sprintf("activityRepository - Create - INN: %s, ActivityTypeID: %d, Category: %s",
		activity.INN, activity.ActivityTypeID, activity.VisitorCategory), nil, nil)

	dbModel := r.toDBActivity(activity)
	_, err := r.db.NamedExecContext(ctx, createActivityQuery, dbModel)
	return err
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

func (r *activityRepository) Update(ctx context.Context, activity *entity.Activity) error {
	r.logger.LogInfo(fmt.Sprintf("activityRepository - Update - INN: %s, ActivityTypeID: %d, Category: %s",
		activity.INN, activity.ActivityTypeID, activity.VisitorCategory), nil, nil)

	dbModel := r.toDBActivity(activity)
	_, err := r.db.NamedExecContext(ctx, updateActivityQuery, dbModel)
	return err
}

func (r *activityRepository) Delete(ctx context.Context, inn string, activityTypeID int, visitorCategory entity.VisitorCategory, year int16) error {
	r.logger.LogInfo(fmt.Sprintf("activityRepository - Delete - INN: %s, Type: %d, Category: %s, Year: %d",
		inn, activityTypeID, visitorCategory, year), nil, nil)

	_, err := r.db.ExecContext(ctx, deleteActivityQuery, inn, activityTypeID, visitorCategory, year)
	return err
}

func (r *activityRepository) List(ctx context.Context) ([]entity.Activity, error) {
	r.logger.LogInfo("activityRepository - List", nil, nil)

	var dbModels []model.Activity
	err := r.db.SelectContext(ctx, &dbModels, listActivitiesQuery)
	if err != nil {
		r.logger.LogError("activityRepository - List", nil, err)
		return nil, err
	}

	entities := make([]entity.Activity, len(dbModels))
	for i, m := range dbModels {
		entities[i] = r.toEntityActivity(&m)
	}
	return entities, nil
}

// Маппинг в модель БД
func (r *activityRepository) toDBActivity(a *entity.Activity) *model.Activity {
	return &model.Activity{
		INN:                  a.INN,
		ActivityTypeID:       a.ActivityTypeID,
		VisitorCategory:      string(a.VisitorCategory),
		CostSharePercent:     a.CostSharePercent,
		RevenueAmount:        a.RevenueAmount,
		TotalCount:           a.TotalCount,
		StateTaskCount:       a.StateTaskCount,
		RevenueActivityCount: a.RevenueActivityCount,
		Year:                 a.Year,
	}
}

// Маппинг в доменную сущность
func (r *activityRepository) toEntityActivity(m *model.Activity) entity.Activity {
	return entity.Activity{
		INN:                  m.INN,
		ActivityTypeID:       m.ActivityTypeID,
		VisitorCategory:      entity.VisitorCategory(m.VisitorCategory),
		CostSharePercent:     m.CostSharePercent,
		RevenueAmount:        m.RevenueAmount,
		TotalCount:           m.TotalCount,
		StateTaskCount:       m.StateTaskCount,
		RevenueActivityCount: m.RevenueActivityCount,
		Year:                 m.Year,
	}
}
