package interactor

import (
	orgRepoInterface "vkr/internal/domain/organization/repository"
	orgService "vkr/internal/domain/organization/service"
	orgStorage "vkr/internal/infrastructure/db/pg/organization"
	handler "vkr/internal/presenter/http/handler/organisation"
	orgUseCase "vkr/internal/usecase/organisation"
)

func (i *interactor) NewOrganisationRepository() orgRepoInterface.Repository {
	return orgStorage.NewOrganizationRepository(i.conn, i.logger)
}

func (i *interactor) NewOrganizationService() orgService.OrganizationService {
	return orgService.NewStudentService()
}

func (i *interactor) NewOrganizationUseCase() orgUseCase.UseCase {
	return orgUseCase.NewOrganizationUseCase(
		i.NewOrganisationRepository(),
		i.NewOrganizationService(),
	)
}

func (i *interactor) NewOrganizationHandler() handler.OrganizationHandler {
	return handler.NewOrganizationHandler(i.NewOrganizationUseCase(), i.logger)
}
