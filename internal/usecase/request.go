package usecase

import (
	"Hermes/internal/usecase/dto"
	"context"
)

type RequestUsecase interface {
	// Gets all the requests filtered by provided arguments.
	// 'limit' sets max number of request to hand off.
	// 'forDays' sets age of the requests to be taken.
	// 'offset' is provided for pagination.
	// In case a parameter is 0 it is not used to filter the result.
	GetAll(ctx context.Context, params dto.GetAllRequests) ([]dto.ReturnRequest, error)
	// Gets all the requests for a specific responder filtered by provided arguments.
	// 'limit' sets max number of request to hand off.
	// 'forDays' sets age of the requests to be taken.
	// 'offset' is provided for pagination.
	// In case a parameter is 0 it is not used to filter the result.
	// If request with such ID does not exist returns ErrNotFound.
	GetForResponder(ctx context.Context, params dto.GetForResponder) ([]dto.ReturnRequest, error)
	// Attemtps to return request with provided ID.
	// If no ID found returns ErrNotFound.
	GetByID(ctx context.Context, ID int) (*dto.ReturnRequest, error)
	// Attempts to create a new request row.
	Create(ctx context.Context, createEntity dto.CreateRequest) (*dto.ReturnRequest, error)
	// Attempts to close existing request.
	// If request with provided ID not found returns ErrNotFound
	// If request already closed returns 0 otherwise 1.
	Close(ctx context.Context, ID int) (int, error)
	// Attempts to update an existing request.
	// If request with such ID does not exist returns ErrNotFound.
	// If request already closed returns ErrInvalidOperation.
	Update(ctx context.Context, request dto.UpdateRequest) (*dto.ReturnRequest, error)

	// Attempts to create new request type.
	// If type with provided name already exists returns ErrAlreadyTaken.
	CreateType(ctx context.Context, requestType dto.CreateRequestType) (*dto.ReturnRequestType, error)
	// Fetches all the request types.
	// If 'onlyActive' flag set to true returns only active types.
	GetTypes(ctx context.Context, onlyActive bool) ([]dto.ReturnRequestType, error)
	// Attemts to find request type by provided ID.
	// If 'onlyActive' flag set to true returns only active type.
	// If request type with such ID not found returns ErrNotFound.
	GetTypeWithID(ctx context.Context, ID int, onlyActive bool) (*dto.ReturnRequestType, error)
	// Attemts to active a request type with provided ID.
	// If ID not found returns ErrNotFound. If type already active returns 0 otherwise 1.
	ActivateType(ctx context.Context, ID int) (int, error)
	// Attemts to deactive a request type with provided ID.
	// If ID not found returns ErrNotFound. If type already inactive returns 0 otherwise 1.
	DeactivateType(ctx context.Context, ID int) (int, error)
}
