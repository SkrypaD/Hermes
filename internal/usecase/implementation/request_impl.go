package implementation

import (
	"Hermes/internal/domain"
	"Hermes/internal/usecase/dto"
	"context"
)

type RequestUsecase struct {
	usrRepo     domain.UserRepository
	reqRepo     domain.RequestRepository
	reqTypeRepo domain.RequestTypeRepository
}

func NewRequestUsecase(usrR domain.UserRepository, reqR domain.RequestRepository, reqTR domain.RequestTypeRepository) *RequestUsecase {
	return &RequestUsecase{
		usrRepo:     usrR,
		reqRepo:     reqR,
		reqTypeRepo: reqTR,
	}
}

// Gets all the requests filtered by provided arguments.
// 'limit' sets max number of request to hand off.
// 'forDays' sets age of the requests to be taken.
// 'offset' is provided for pagination.
// In case a parameter is 0 it is not used to filter the result.
func (reqUsc *RequestUsecase) GetAll(ctx context.Context, params dto.GetAllRequests) ([]dto.ReturnRequest, error) {
	res, err := reqUsc.reqRepo.GetAll(ctx, params.ForDays, params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}

	reqstType, err := reqUsc.reqTypeRepo.GetAll(ctx, false)
	if err != nil {
		return nil, err
	}

	requestTypes := make(map[int]string)
	for _, t := range reqstType {
		requestTypes[t.ID] = t.Name
	}

	returnReq := []dto.ReturnRequest{}
	for _, request := range res {
		returnReq = append(returnReq, dto.ReturnRequest{
			ID:            request.ID,
			Title:         request.Title,
			Description:   request.Description,
			RequestType:   requestTypes[request.RequestTypeID],
			CreatedAt:     request.CreatedAt,
			UpdatedAt:     request.UpdatedAt,
			ClosedAt:      request.ClosedAt,
			DispatcherID:  request.DispatcherID,
			ResponderID:   request.ResponderID,
			RequestTypeID: request.RequestTypeID,
		})
	}
	return returnReq, nil
}

// Gets all the requests for a specific responder filtered by provided arguments.
// 'limit' sets max number of request to hand off.
// 'forDays' sets age of the requests to be taken.
// 'offset' is provided for pagination.
// In case a parameter is 0 it is not used to filter the result.
// If request with such ID does not exist returns ErrNotFound.
func (reqUsc *RequestUsecase) GetForResponder(ctx context.Context, params dto.GetForResponder) ([]dto.ReturnRequest, error) {

	res, err := reqUsc.reqRepo.GetForResponder(ctx, params.ResponderID, params.ForDays, params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}

	reqstType, err := reqUsc.reqTypeRepo.GetAll(ctx, false)
	if err != nil {
		return nil, err
	}

	requestTypes := make(map[int]string)
	for _, t := range reqstType {
		requestTypes[t.ID] = t.Name
	}

	returnReq := []dto.ReturnRequest{}
	for _, request := range res {
		returnReq = append(returnReq, dto.ReturnRequest{
			ID:            request.ID,
			Title:         request.Title,
			Description:   request.Description,
			RequestType:   requestTypes[request.RequestTypeID],
			CreatedAt:     request.CreatedAt,
			UpdatedAt:     request.UpdatedAt,
			ClosedAt:      request.ClosedAt,
			DispatcherID:  request.DispatcherID,
			ResponderID:   request.ResponderID,
			RequestTypeID: request.RequestTypeID,
		})
	}
	return returnReq, nil
}

// Attemtps to return request with provided ID.
// If no ID found returns ErrNotFound.
func (reqUsc *RequestUsecase) GetByID(ctx context.Context, ID int) (*dto.ReturnRequest, error) {
	request, err := reqUsc.reqRepo.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	reqstType, err := reqUsc.reqTypeRepo.GetByID(ctx, request.RequestTypeID, false)
	if err != nil {
		return nil, err
	}

	return &dto.ReturnRequest{
		ID:            request.ID,
		Title:         request.Title,
		Description:   request.Description,
		RequestType:   reqstType.Name,
		CreatedAt:     request.CreatedAt,
		UpdatedAt:     request.UpdatedAt,
		ClosedAt:      request.ClosedAt,
		DispatcherID:  request.DispatcherID,
		ResponderID:   request.ResponderID,
		RequestTypeID: request.RequestTypeID,
	}, nil
}

// Attempts to create a new request row.
func (reqUsc *RequestUsecase) Create(ctx context.Context, createEntity dto.CreateRequest) (*dto.ReturnRequest, error) {
	createReq := domain.Request{
		Title:         createEntity.Title,
		Description:   createEntity.Description,
		DispatcherID:  createEntity.DispatcherID,
		ResponderID:   createEntity.ResponderID,
		RequestTypeID: createEntity.RequestTypeID,
	}

	res, err := reqUsc.reqRepo.Create(ctx, createReq)
	if err != nil {
		return nil, err
	}

	typ, err := reqUsc.reqTypeRepo.GetByID(ctx, res.RequestTypeID, false)
	if err != nil {
		return nil, err
	}

	return &dto.ReturnRequest{
		ID:            res.ID,
		Title:         res.Title,
		Description:   res.Description,
		RequestType:   typ.Name,
		CreatedAt:     res.CreatedAt,
		UpdatedAt:     res.UpdatedAt,
		ClosedAt:      res.ClosedAt,
		DispatcherID:  res.DispatcherID,
		ResponderID:   res.ResponderID,
		RequestTypeID: res.RequestTypeID,
	}, nil
}

// Attempts to close existing request.
// If request with provided ID not found returns ErrNotFound
// If request already closed returns 0 otherwise 1.
func (reqUsc *RequestUsecase) Close(ctx context.Context, ID int) (int, error) {
	_, err := reqUsc.reqRepo.GetByID(ctx, ID)
	if err != nil {
		return -1, err
	}

	flag, err := reqUsc.Close(ctx, ID)
	return flag, err
}

// Attempts to update an existing request.
// If request with such ID does not exist returns ErrNotFound.
// If request already closed returns ErrInvalidOperation.
func (reqUsc *RequestUsecase) Update(ctx context.Context, request dto.UpdateRequest) (*dto.ReturnRequest, error)

// Attempts to create new request type.
// If type with provided name already exists returns ErrAlreadyTaken.
func (reqUsc *RequestUsecase) CreateType(ctx context.Context, requestType dto.CreateRequestType) (*dto.ReturnRequestType, error) {
	reqType := domain.RequestType{
		Name: requestType.Name,
	}

	res, err := reqUsc.reqTypeRepo.Create(ctx, reqType)
	if err != nil {
		return nil, err
	}

	return &dto.ReturnRequestType{
		ID:         res.ID,
		Name:       res.Name,
		IsRelevant: res.IsRelevant,
		CreatedAt:  res.CreatedAt,
	}, nil
}

// Fetches all the request types.
// If 'onlyActive' flag set to true returns only active types.
func (reqUsc *RequestUsecase) GetTypes(ctx context.Context, onlyActive bool) ([]dto.ReturnRequestType, error) {
	types, err := reqUsc.reqTypeRepo.GetAll(ctx, onlyActive)
	if err != nil {
		return nil, err
	}

	returnTypes := make([]dto.ReturnRequestType, len(types))

	for _, t := range types {
		returnTypes = append(returnTypes, dto.ReturnRequestType{
			ID:         t.ID,
			Name:       t.Name,
			IsRelevant: t.IsRelevant,
			CreatedAt:  t.CreatedAt,
		})
	}
	return returnTypes, nil
}

// Attemts to find request type by provided ID.
// If 'onlyActive' flag set to true returns only active type.
// If request type with such ID not found returns ErrNotFound.
func (reqUsc *RequestUsecase) GetTypeWithID(ctx context.Context, ID int, onlyActive bool) (*dto.ReturnRequestType, error) {
	res, err := reqUsc.reqTypeRepo.GetByID(ctx, ID, onlyActive)
	if err != nil {
		return nil, err
	}

	return &dto.ReturnRequestType{
		ID:         res.ID,
		Name:       res.Name,
		IsRelevant: res.IsRelevant,
		CreatedAt:  res.CreatedAt,
	}, nil
}

// Attemts to active a request type with provided ID.
// If ID not found returns ErrNotFound. If type already active returns 0 otherwise 1.
func (reqUsc *RequestUsecase) ActivateType(ctx context.Context, ID int) (int, error) {
	res, err := reqUsc.reqTypeRepo.Activate(ctx, ID)
	return res, err
}

// Attemts to deactive a request type with provided ID.
// If ID not found returns ErrNotFound. If type already inactive returns 0 otherwise 1.
func (reqUsc *RequestUsecase) DeactivateType(ctx context.Context, ID int) (int, error) {
	res, err := reqUsc.reqTypeRepo.Deactivate(ctx, ID)
	return res, err
}
