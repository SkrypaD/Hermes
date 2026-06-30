package domain

import "context"

type UserUseCase interface {
	Register(ctx context.Context, user *User) error
	GetAll(ctx context.Context) ([]User, error)
	GetById(ctx context.Context, id ID) (*User, error)
}
