package interactor

import (
	reportingFormRepo "vkr/internal/domain/reporting_form/repository"
	reportingFormStorage "vkr/internal/infrastructure/db/pg/reporting_form"
	reportingFormHandler "vkr/internal/presenter/http/handler/reporting_form"
	reportingFormUseCase "vkr/internal/usecase/reporting_form"
)

func (i *interactor) NewReportingFormRepository() reportingFormRepo.Repository {
	return reportingFormStorage.NewRepository(i.conn, i.logger)
}

func (i *interactor) NewLaborFormUseCase() reportingFormUseCase.LaborFormUseCase {
	return reportingFormUseCase.NewLaborFormUseCase(i.NewLaborRepository())
}

func (i *interactor) NewActivitiesFormUseCase() reportingFormUseCase.ActivitiesFormUseCase {
	return reportingFormUseCase.NewActivitiesFormUseCase(i.NewActivityRepository())
}

func (i *interactor) NewFinancialFormUseCase() reportingFormUseCase.FinancialFormUseCase {
	return reportingFormUseCase.NewFinancialFormUseCase(i.NewReportingFormRepository())
}

func (i *interactor) NewReportingFormHandler() reportingFormHandler.ReportingFormHandler {
	return reportingFormHandler.NewReportingFormHandler(
		i.NewLaborFormUseCase(),
		i.NewActivitiesFormUseCase(),
		i.NewFinancialFormUseCase(),
		i.logger,
	)
}

