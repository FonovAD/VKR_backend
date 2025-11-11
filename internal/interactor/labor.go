package interactor

import (
	laborRepo "vkr/internal/domain/labor/repository"
	laborService "vkr/internal/domain/labor/service"
	laborStorage "vkr/internal/infrastructure/db/pg/labor"
	laborHandler "vkr/internal/presenter/http/handler/labor"
	laborUseCase "vkr/internal/usecase/labor"
)

func (i *interactor) NewLaborRepository() laborRepo.Repository {
	return laborStorage.NewRepository(i.conn, i.logger)
}

func (i *interactor) NewLaborService() laborService.Service {
	return laborService.NewService()
}

func (i *interactor) NewLaborUseCase() laborUseCase.UseCase {
	return laborUseCase.NewUseCase(
		i.NewLaborRepository(),
		i.NewLaborService(),
	)
}

func (i *interactor) NewLaborHandler() laborHandler.Handler {
	return laborHandler.NewHandler(i.NewLaborUseCase(), i.logger)
}
