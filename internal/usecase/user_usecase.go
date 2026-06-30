package usecase

import (
	"Hermes/internal/domain"
	"context"
	"errors"
)

type UserUseCase struct {
	usr_repo domain.UserRepository
}

func NewUserUseCase(usr_repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{
		usr_repo: usr_repo,
	}
}

func (u *UserUseCase) Register(ctx context.Context, user *domain.User) error {
	// TODO: Add user creation validation
	return u.usr_repo.Create(ctx, *user)
}

func (u *UserUseCase) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := u.usr_repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserUseCase) GetById(ctx context.Context, id domain.ID) (*domain.User, error) {
	if id < 1 {
		return nil, errors.New("Id can not be a negative number.")
	}
	user, err := u.usr_repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
