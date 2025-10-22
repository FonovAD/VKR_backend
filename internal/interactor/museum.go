package interactor

import (
	museumRepoInterface "vkr/internal/domain/museum/repository"
	museumService "vkr/internal/domain/museum/service"
	museumStorage "vkr/internal/infrastructure/db/pg/museum"
	museumHandler "vkr/internal/presenter/http/handler/museum"
	museumUseCase "vkr/internal/usecase/museum"
)

func (i *interactor) NewMuseumRepository() museumRepoInterface.Repository {
	return museumStorage.NewMuseumRepository(i.conn, i.logger)
}

func (i *interactor) NewMuseumService() museumService.Service {
	// Если у вас пока нет реализации сервиса — можно вернуть заглушку или пустую структуру
	// Например: return &museumService.MuseumService{}
	return *museumService.NewService()
}

func (i *interactor) NewMuseumUseCase() museumUseCase.UseCase {
	return museumUseCase.NewMuseumUseCase(
		i.NewMuseumRepository(),
		i.NewMuseumService(),
	)
}

func (i *interactor) NewMuseumHandler() museumHandler.MuseumHandler {
	return museumHandler.NewMuseumHandler(i.NewMuseumUseCase(), i.logger)
}
