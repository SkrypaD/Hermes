package domain

import (
	"context"
	"time"
)

type User struct {
	ID        ID
	Name      string
	CreatedAt time.Time
	Type      string

	Requests []Request
}

type UserRepository interface {
	GetByID(ctx context.Context, id ID) (User, error)
	GetAll(ctx context.Context) ([]User, error)
	GetByType(ctx context.Context, user_type string) ([]User, error)
	Create(ctx context.Context, user User) error
	Delete(ctx context.Context, id ID) error
}
