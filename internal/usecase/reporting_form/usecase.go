package reporting_form

import (
	"context"
	activRepo "vkr/internal/domain/activities/repository"
	laborRepo "vkr/internal/domain/labor/repository"
	reportingFormRepo "vkr/internal/domain/reporting_form/repository"
	value "vkr/internal/domain/value"
)

// LaborFormUseCase предоставляет данные для формы трудовых ресурсов
type LaborFormUseCase interface {
	GetLaborFormData(ctx context.Context, organizationID int) (LaborFormData, error)
	GetLaborFormDataByINN(ctx context.Context, inn string) (LaborFormData, error)
}

// ActivitiesFormUseCase предоставляет данные для формы объемов (активностей)
type ActivitiesFormUseCase interface {
	GetActivitiesFormData(ctx context.Context, organizationID int, year int16) (ActivitiesFormData, error)
	GetActivitiesFormDataByINN(ctx context.Context, inn string, year int16) (ActivitiesFormData, error)
}

// FinancialFormUseCase предоставляет данные для формы поступлений и расходов
type FinancialFormUseCase interface {
	GetFinancialFormData(ctx context.Context, organizationID int, year int16) (FinancialFormData, error)
	GetFinancialFormDataByINN(ctx context.Context, inn string, year int16) (FinancialFormData, error)
}

// LaborFormData структура данных для формы трудовых ресурсов
type LaborFormData struct {
	Year      int16           `json:"year"`
	LaborData value.LaborData `json:"labor_data"`
}

// ActivitiesFormData структура данных для формы объемов
type ActivitiesFormData struct {
	Year       int16                    `json:"year"`
	Activities []ActivityFormItem       `json:"activities"`
}

// ActivityFormItem элемент формы активности
type ActivityFormItem struct {
	ID                   *int64   `json:"id,omitempty"`
	ActivityTypeID       *int     `json:"activity_type_id,omitempty"`
	ActivityTypeName     *string  `json:"activity_type_name,omitempty"`
	CustomActivityID     *int     `json:"custom_activity_id,omitempty"`
	CustomActivityName   *string  `json:"custom_activity_name,omitempty"`
	VolumeIndicator      *string  `json:"volume_indicator,omitempty"`
	CostSharePercent     *float64 `json:"cost_share_percent,omitempty"`
	RevenueAmount        *float64 `json:"revenue_amount,omitempty"`
	VisitorCategory      string   `json:"visitor_category"`
	TotalCount           *int64   `json:"total_count,omitempty"`
	StateTaskCount       *int64   `json:"state_task_count,omitempty"`
	RevenueActivityCount *int64   `json:"revenue_activity_count,omitempty"`
}

// FinancialFormData структура данных для формы поступлений и расходов
type FinancialFormData struct {
	Year      int16            `json:"year"`
	Financial value.FinancialData `json:"financial"`
}

type laborFormUseCase struct {
	laborRepo laborRepo.Repository
}

func NewLaborFormUseCase(laborRepo laborRepo.Repository) LaborFormUseCase {
	return &laborFormUseCase{
		laborRepo: laborRepo,
	}
}

func (u *laborFormUseCase) GetLaborFormData(ctx context.Context, organizationID int) (LaborFormData, error) {
	laborData, err := u.laborRepo.GetByOrganizationID(ctx, organizationID)
	if err != nil {
		return LaborFormData{}, err
	}

	// По умолчанию используем текущий год, можно добавить параметр года в будущем
	return LaborFormData{
		Year:      2022, // TODO: сделать параметром или получать из конфига
		LaborData: laborData,
	}, nil
}

func (u *laborFormUseCase) GetLaborFormDataByINN(ctx context.Context, inn string) (LaborFormData, error) {
	laborData, err := u.laborRepo.GetByOrganizationINN(ctx, inn)
	if err != nil {
		return LaborFormData{}, err
	}

	return LaborFormData{
		Year:      2022, // TODO: сделать параметром или получать из конфига
		LaborData: laborData,
	}, nil
}

type activitiesFormUseCase struct {
	activityRepo activRepo.Repository
}

func NewActivitiesFormUseCase(activityRepo activRepo.Repository) ActivitiesFormUseCase {
	return &activitiesFormUseCase{
		activityRepo: activityRepo,
	}
}

func (u *activitiesFormUseCase) GetActivitiesFormData(ctx context.Context, organizationID int, year int16) (ActivitiesFormData, error) {
	activities, err := u.activityRepo.GetByOrganizationIDAndYear(ctx, organizationID, year)
	if err != nil {
		return ActivitiesFormData{}, err
	}

	var formItems []ActivityFormItem
	for _, activity := range activities {
		item := ActivityFormItem{
			ID:                   activity.ID,
			ActivityTypeID:       activity.ActivityTypeID,
			ActivityTypeName:     activity.ActivityTypeName,
			VolumeIndicator:      activity.VolumeIndicator,
			CustomActivityID:     activity.CustomActivityID,
			CustomActivityName:   activity.CustomActivityName,
			CostSharePercent:     activity.CostSharePercent,
			RevenueAmount:        activity.RevenueAmount,
			VisitorCategory:      string(activity.VisitorCategory),
			TotalCount:           activity.TotalCount,
			StateTaskCount:       activity.StateTaskCount,
			RevenueActivityCount: activity.RevenueActivityCount,
		}
		formItems = append(formItems, item)
	}

	return ActivitiesFormData{
		Year:       year,
		Activities: formItems,
	}, nil
}

func (u *activitiesFormUseCase) GetActivitiesFormDataByINN(ctx context.Context, inn string, year int16) (ActivitiesFormData, error) {
	activities, err := u.activityRepo.GetByINN(ctx, inn)
	if err != nil {
		return ActivitiesFormData{}, err
	}

	// Фильтруем по году
	var filteredActivities []ActivityFormItem
	for _, activity := range activities {
		if activity.Year == year {
			item := ActivityFormItem{
				ID:                   activity.ID,
				ActivityTypeID:       activity.ActivityTypeID,
				ActivityTypeName:     activity.ActivityTypeName,
				VolumeIndicator:      activity.VolumeIndicator,
				CustomActivityID:     activity.CustomActivityID,
				CustomActivityName:   activity.CustomActivityName,
				CostSharePercent:     activity.CostSharePercent,
				RevenueAmount:        activity.RevenueAmount,
				VisitorCategory:      string(activity.VisitorCategory),
				TotalCount:           activity.TotalCount,
				StateTaskCount:       activity.StateTaskCount,
				RevenueActivityCount: activity.RevenueActivityCount,
			}
			filteredActivities = append(filteredActivities, item)
		}
	}

	return ActivitiesFormData{
		Year:       year,
		Activities: filteredActivities,
	}, nil
}

type financialFormUseCase struct {
	financialRepo reportingFormRepo.Repository
}

func NewFinancialFormUseCase(financialRepo reportingFormRepo.Repository) FinancialFormUseCase {
	return &financialFormUseCase{
		financialRepo: financialRepo,
	}
}

func (u *financialFormUseCase) GetFinancialFormData(ctx context.Context, organizationID int, year int16) (FinancialFormData, error) {
	financialData, err := u.financialRepo.GetFinancialDataByOrganizationIDAndYear(ctx, organizationID, year)
	if err != nil {
		return FinancialFormData{}, err
	}

	return FinancialFormData{
		Year:      year,
		Financial: financialData,
	}, nil
}

func (u *financialFormUseCase) GetFinancialFormDataByINN(ctx context.Context, inn string, year int16) (FinancialFormData, error) {
	financialData, err := u.financialRepo.GetFinancialDataByOrganizationINNAndYear(ctx, inn, year)
	if err != nil {
		return FinancialFormData{}, err
	}

	return FinancialFormData{
		Year:      year,
		Financial: financialData,
	}, nil
}

