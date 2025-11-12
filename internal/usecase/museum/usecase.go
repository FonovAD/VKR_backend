package museum

import (
	"context"

	"vkr/internal/domain/entity"
	museumRepo "vkr/internal/domain/museum/repository"
	museumService "vkr/internal/domain/museum/service"
)

type UseCase interface {
	Create(ctx context.Context, museum *entity.Museum) (*entity.Museum, error)
	GetByID(ctx context.Context, id entity.MuseumID) (*entity.Museum, error)
	Update(ctx context.Context, museum *entity.Museum) (*entity.Museum, error)
	Delete(ctx context.Context, id entity.MuseumID) error
	FindByINN(ctx context.Context, inn string) (*entity.Museum, error)
	FindByOwner(ctx context.Context, ownerID entity.OrganizationID) ([]*entity.Museum, error)
	List(ctx context.Context, name, museumType *string, limit, offset int) (*museumRepo.ListResult, error)
}

type museumUseCase struct {
	repo    museumRepo.Repository
	service museumService.Service
}

func NewMuseumUseCase(repo museumRepo.Repository, service museumService.Service) UseCase {
	return &museumUseCase{
		repo:    repo,
		service: service,
	}
}

func (u *museumUseCase) Create(ctx context.Context, museum *entity.Museum) (*entity.Museum, error) {
	if err := u.repo.Create(ctx, museum); err != nil {
		return nil, err
	}
	return museum, nil
}

func (u *museumUseCase) GetByID(ctx context.Context, id entity.MuseumID) (*entity.Museum, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *museumUseCase) Update(ctx context.Context, museum *entity.Museum) (*entity.Museum, error) {
	if err := u.repo.Update(ctx, museum); err != nil {
		return nil, err
	}
	return museum, nil
}

func (u *museumUseCase) Delete(ctx context.Context, id entity.MuseumID) error {
	return u.repo.Delete(ctx, id)
}

func (u *museumUseCase) FindByINN(ctx context.Context, inn string) (*entity.Museum, error) {
	return u.repo.FindByINN(ctx, inn)
}

func (u *museumUseCase) FindByOwner(ctx context.Context, ownerID entity.OrganizationID) ([]*entity.Museum, error) {
	return u.repo.FindByOwner(ctx, ownerID)
}

func (u *museumUseCase) List(ctx context.Context, name, museumType *string, limit, offset int) (*museumRepo.ListResult, error) {
	return u.repo.List(ctx, museumRepo.ListParams{
		Name:       name,
		MuseumType: museumType,
		Limit:      limit,
		Offset:     offset,
	})
}
