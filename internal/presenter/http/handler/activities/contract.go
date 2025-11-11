package activity

import (
	"vkr/internal/domain/entity"
)

type BadRequestResponse struct {
	ErrorMsg string `json:"error"`
}

type InternalServerErrorResponse struct {
	ErrorMsg string `json:"error"`
}

var (
	ErrInternalServer = &InternalServerErrorResponse{ErrorMsg: "internal server error"}
)

type VisitorCategory = entity.VisitorCategory

type CreateActivityDTO struct {
	INN                  string          `json:"inn"`
	ActivityTypeID       *int            `json:"activity_type_id,omitempty"`
	CustomActivityID     *int            `json:"custom_activity_id,omitempty"`
	VisitorCategory      VisitorCategory `json:"visitor_category"`
	CostSharePercent     *float64        `json:"cost_share_percent,omitempty"`
	RevenueAmount        *float64        `json:"revenue_amount,omitempty"`
	TotalCount           *int64          `json:"total_count,omitempty"`
	StateTaskCount       *int64          `json:"state_task_count,omitempty"`
	RevenueActivityCount *int64          `json:"revenue_activity_count,omitempty"`
	Year                 int16           `json:"year"`
}

type UpdateActivityDTO struct {
	ID                   *int64          `json:"id"`
	CostSharePercent     *float64        `json:"cost_share_percent,omitempty"`
	RevenueAmount        *float64        `json:"revenue_amount,omitempty"`
	TotalCount           *int64          `json:"total_count,omitempty"`
	StateTaskCount       *int64          `json:"state_task_count,omitempty"`
	RevenueActivityCount *int64          `json:"revenue_activity_count,omitempty"`
	Year                 int16           `json:"year"`
}

type ActivityResponse struct {
	ID                   *int64          `json:"id,omitempty"`
	IDOwner              int              `json:"id_owner"`
	INN                  string           `json:"inn"`
	ActivityTypeID       *int             `json:"activity_type_id,omitempty"`
	ActivityTypeName     *string          `json:"activity_type_name,omitempty"`
	CustomActivityID     *int             `json:"custom_activity_id,omitempty"`
	CustomActivityName   *string          `json:"custom_activity_name,omitempty"`
	VisitorCategory      VisitorCategory  `json:"visitor_category"`
	CostSharePercent     *float64         `json:"cost_share_percent,omitempty"`
	RevenueAmount        *float64         `json:"revenue_amount,omitempty"`
	TotalCount           *int64           `json:"total_count,omitempty"`
	StateTaskCount       *int64           `json:"state_task_count,omitempty"`
	RevenueActivityCount *int64           `json:"revenue_activity_count,omitempty"`
	Year                 int16            `json:"year"`
}

func NewActivityResponse(activity entity.Activity) ActivityResponse {
	return ActivityResponse{
		ID:                   activity.ID,
		IDOwner:              int(activity.IDOwner),
		INN:                  activity.INN,
		ActivityTypeID:       activity.ActivityTypeID,
		ActivityTypeName:     activity.ActivityTypeName,
		CustomActivityID:     activity.CustomActivityID,
		CustomActivityName:   activity.CustomActivityName,
		VisitorCategory:      activity.VisitorCategory,
		CostSharePercent:     activity.CostSharePercent,
		RevenueAmount:        activity.RevenueAmount,
		TotalCount:           activity.TotalCount,
		StateTaskCount:       activity.StateTaskCount,
		RevenueActivityCount: activity.RevenueActivityCount,
		Year:                 activity.Year,
	}
}
