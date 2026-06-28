package domain

import (
	"context"
	"time"
)

type Request struct {
	Id          ID
	Title       string
	Description string
	CreatedAt   time.Time
	ClosedAt    time.Time

	RegistratorID ID
	ResponderID   ID
	RequestTypeID ID
}

type RequestType struct {
	Id   ID
	Name string
}

type RequestRepository interface {
	GetByID(ctx context.Context, id ID) (Request, error)
	GetAll(ctx context.Context) ([]Request, error)
	GetByType(ctx context.Context, req_type RequestType) ([]Request, error)
	Create(ctx context.Context, request Request) error
	Delete(ctx context.Context, id ID) error
}
