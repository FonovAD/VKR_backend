package interactor

import (
	activityStorageInterface "vkr/internal/domain/activities/repository"
	activService "vkr/internal/domain/activities/service"
	activityStorage "vkr/internal/infrastructure/db/pg/activities"
	activity "vkr/internal/presenter/http/handler/activities"
	activityUseCase "vkr/internal/usecase/activities"
)

func (i *interactor) NewActivityRepository() activityStorageInterface.Repository {
	return activityStorage.NewActivityRepository(i.conn, i.logger)
}

func (i *interactor) NewActivityService() activService.Service {
	return activService.NewService()
}

func (i *interactor) NewActivityUseCase() activityUseCase.ActivityUseCase {
	return activityUseCase.NewActivityUseCase(
		i.NewActivityRepository(),
		i.NewActivityService(),
	)
}

func (i *interactor) NewActivityHandler() activity.ActivityHandler {
	return activity.NewActivityHandler(i.NewActivityUseCase(), i.logger)
}
