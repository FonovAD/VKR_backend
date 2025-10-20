package entity

import "vkr/internal/domain/statuses"

type Metric struct {
	ID           int
	Name         string
	AudienceType AudienceType // enum: internal, external, all
}
