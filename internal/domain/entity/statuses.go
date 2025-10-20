package entity

type FormStatus string

const (
	FormStatusDraft     FormStatus = "draft"
	FormStatusSubmitted FormStatus = "submitted"
	FormStatusApproved  FormStatus = "approved"
)

type LocationType string

const (
	LocationTypeInternal LocationType = "internal"
	LocationTypeExternal LocationType = "external"
	LocationTypeOnline   LocationType = "online"
)

type AudienceType string

const (
	AudienceTypeInternal AudienceType = "internal"
	AudienceTypeExternal AudienceType = "external"
	AudienceTypeAll      AudienceType = "all"
)
