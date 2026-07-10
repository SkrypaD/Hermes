package domain

import (
	"context"
	"time"
)

type Request struct {
	ID            int
	Title         string
	Description   string
	CreatedAt     time.Time
	UpdatedAt     *time.Time
	ClosedAt      *time.Time
	DispatcherID  int
	ResponderID   int
	RequestTypeID int
}

type RequestRepository interface {
	// Gets all the requests filtered by provided arguments.
	// 'limit' sets max number of request to hand off.
	// 'forDays' sets age of the requests to be taken.
	// 'offset' is provided for pagination.
	// In case a parameter is 0 it is not used to filter the result.
	GetAll(ctx context.Context, forDays int, limit int, offset int) ([]Request, error)
	// Gets all the requests for a specific responder filtered by provided arguments.
	// 'limit' sets max number of request to hand off.
	// 'forDays' sets age of the requests to be taken.
	// 'offset' is provided for pagination.
	// In case a parameter is 0 it is not used to filter the result.
	// If request with such ID does not exist returns ErrNotFound.
	GetForResponder(ctx context.Context, responderID int, forDays int, limit int, offset int) ([]Request, error)
	// Attempts to find a request by its id. If no request found returns ErrNotFound.
	GetByID(ctx context.Context, ID int) (*Request, error)
	// Attempts to create a new request entity in the database.
	Create(ctx context.Context, request Request) (*Request, error)
	// Attempts to close a request  by its id. If request already closed returns 0 otherwise 1.
	// If request with such ID does not exist returns ErrNotFound.
	Close(ctx context.Context, ID int) (int, error)
	// Attempts to update an existing request.
	// If request with such ID does not exist returns ErrNotFound.
	Update(ctx context.Context, request Request) (*Request, error)
}

type RequestType struct {
	ID         int
	Name       string
	IsRelevant bool
	CreatedAt  time.Time
}

type RequestTypeRepository interface {
	// Attempts to create new request type entity. If entity with such name already
	// exists returns ErrAlreadyTaken
	Create(ctx context.Context, requestType RequestType) (*RequestType, error)
	// Returns all the existing request types
	GetAll(ctx context.Context) ([]RequestType, error)
	// Searches for the request type by id. If no type found returns ErrNotFound
	// If isActive true returns only if request is active.
	GetByID(ctx context.Context, typeID int, isActive bool) (*RequestType, error)
	// Attempts to deactivate a request type with provided id. If type already inactive returns 0 otherwise 1.
	// If to type found with provided id returns ErrNotFound
	Deactivate(ctx context.Context, typeID int) (int, error)
	// Attempts to activate a request type with provided id. If type already active returns 0 otherwise 1.
	// If to type found with provided id returns ErrNotFound
	Activate(ctx context.Context, typeID int) (int, error)
}
