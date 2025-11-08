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
	ActivityTypeID       int             `json:"activity_type_id"`
	VisitorCategory      VisitorCategory `json:"visitor_category"`
	CostSharePercent     *float64        `json:"cost_share_percent,omitempty"`
	RevenueAmount        *float64        `json:"revenue_amount,omitempty"`
	TotalCount           *int64          `json:"total_count,omitempty"`
	StateTaskCount       *int64          `json:"state_task_count,omitempty"`
	RevenueActivityCount *int64          `json:"revenue_activity_count,omitempty"`
	Year                 int16           `json:"year"`
}

type UpdateActivityDTO struct {
	INN                  string          `json:"inn"`
	ActivityTypeID       int             `json:"activity_type_id"`
	VisitorCategory      VisitorCategory `json:"visitor_category"`
	CostSharePercent     *float64        `json:"cost_share_percent,omitempty"`
	RevenueAmount        *float64        `json:"revenue_amount,omitempty"`
	TotalCount           *int64          `json:"total_count,omitempty"`
	StateTaskCount       *int64          `json:"state_task_count,omitempty"`
	RevenueActivityCount *int64          `json:"revenue_activity_count,omitempty"`
	Year                 int16           `json:"year"`
}
