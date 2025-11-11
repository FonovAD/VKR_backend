package organisation

import "vkr/internal/domain/entity"

type BadRequestResponse struct {
	ErrorMsg string `json:"error"`
}

type InternalServerErrorResponse struct {
	ErrorMsg string `json:"error"`
}

type NotFoundResponse struct {
	ErrorMsg string `json:"error"`
}

type OrganizationID = entity.OrganizationID

type CreateOrganizationDTO struct {
	INN         string `json:"inn"`
	Name        string `json:"name"`
	ExistMuseum bool   `json:"exist_museum"`
}

type UpdateOrganizationDTO struct {
	INN         string `json:"inn"`
	Name        string `json:"name"`
	ExistMuseum bool   `json:"exist_museum"`
}
