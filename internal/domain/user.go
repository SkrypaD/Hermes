package domain

import (
	"context"
	"time"
)

type User struct {
	ID        ID
	Name      string
	Type      string
	IsActive  bool
	CreatedAt time.Time

	Requests []Request
}

type UserRepository interface {
	GetByID(ctx context.Context, id ID) (*User, error)
	GetAll(ctx context.Context) ([]User, error)
	GetByType(ctx context.Context, user_type string) ([]User, error)
	Create(ctx context.Context, user User) error

	// Returns all the users of a responder type with requests created
	// in provided number of days.
	GetResponders(ctx context.Context, days_period int) ([]User, error)
}
