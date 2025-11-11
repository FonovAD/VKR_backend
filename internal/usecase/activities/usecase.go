package activities

import (
	"context"
	activRepo "vkr/internal/domain/activities/repository"
	activService "vkr/internal/domain/activities/service"
	"vkr/internal/domain/entity"
)

type ActivityUseCase interface {
	Create(ctx context.Context, activity *entity.Activity) (*entity.Activity, error)
	GetByINN(ctx context.Context, inn string) ([]entity.Activity, error)
	GetByMuseumID(ctx context.Context, museumID entity.MuseumID) ([]entity.Activity, error)
	Update(ctx context.Context, activity *entity.Activity) (*entity.Activity, error)
	Delete(ctx context.Context, inn string, activityTypeID int, visitorCategory entity.VisitorCategory, year int16) error
	DeleteByID(ctx context.Context, id int64) error
	List(ctx context.Context) ([]entity.Activity, error)
}

type activityUseCase struct {
	repo    activRepo.Repository
	service activService.Service
}

func NewActivityUseCase(repo activRepo.Repository, service activService.Service) ActivityUseCase {
	return &activityUseCase{
		repo:    repo,
		service: service,
	}
}

func (u *activityUseCase) Create(ctx context.Context, activity *entity.Activity) (*entity.Activity, error) {
	if err := u.repo.Create(ctx, activity); err != nil {
		return nil, err
	}
	return activity, nil
}

func (u *activityUseCase) GetByINN(ctx context.Context, inn string) ([]entity.Activity, error) {
	return u.repo.GetByINN(ctx, inn)
}

func (u *activityUseCase) GetByMuseumID(ctx context.Context, museumID entity.MuseumID) ([]entity.Activity, error) {
	return u.repo.GetByMuseumID(ctx, museumID)
}

func (u *activityUseCase) Update(ctx context.Context, activity *entity.Activity) (*entity.Activity, error) {
	if err := u.repo.Update(ctx, activity); err != nil {
		return nil, err
	}
	return activity, nil
}

func (u *activityUseCase) Delete(ctx context.Context, inn string, activityTypeID int, visitorCategory entity.VisitorCategory, year int16) error {
	return u.repo.Delete(ctx, inn, activityTypeID, visitorCategory, year)
}

func (u *activityUseCase) DeleteByID(ctx context.Context, id int64) error {
	return u.repo.DeleteByID(ctx, id)
}

func (u *activityUseCase) List(ctx context.Context) ([]entity.Activity, error) {
	return u.repo.List(ctx)
}
