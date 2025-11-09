package interactor

import (
	reportingRepoInterface "vkr/internal/domain/reporting_form/repository"
	reportingService "vkr/internal/domain/reporting_form/service"
	reportingStorage "vkr/internal/infrastructure/db/pg/reporting_form"
	reportingHandler "vkr/internal/presenter/http/handler/reporting_form"
	reportingUseCase "vkr/internal/usecase/reporting_form"
)

func (i *interactor) NewReportingFormRepository() reportingRepoInterface.Repository {
	return reportingStorage.NewReportingFormRepository(i.conn, i.logger)
}

func (i *interactor) NewReportingFormService() *reportingService.Service {
	return reportingService.NewService()
}

func (i *interactor) NewReportingFormUseCase() reportingUseCase.UseCase {
	return reportingUseCase.NewUseCase(
		i.NewReportingFormRepository(),
		i.NewReportingFormService(),
	)
}

func (i *interactor) NewReportingFormHandler() reportingHandler.ReportingFormHandler {
	return reportingHandler.NewReportingFormHandler(i.NewReportingFormUseCase(), i.logger)
}
