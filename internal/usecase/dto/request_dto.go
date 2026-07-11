package dto

import "time"

type ReturnRequest struct {
	ID            int
	Title         string
	Description   string
	RequestType   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ClosedAt      time.Time
	DispatcherID  int
	ResponderID   int
	RequestTypeID int
}

type CreateRequest struct {
	Title         string
	Description   string
	DispatcherID  int
	ResponderID   int
	RequestTypeID int
}

type UpdateRequest struct {
	ID            int
	Title         string
	Description   string
	ResponderID   int
	RequestTypeID int
}

type ReturnRequestType struct {
	ID         int
	Name       string
	IsRelevant bool
	CreatedAt  time.Time
}

type CreateRequestType struct {
	Name string
}

//====================================//

type ActivateRequestType struct {
	ID int
}

type GetAllRequests struct {
	ForDays int
	Offset  int
	Limit   int
}

type GetForResponder struct {
	ResponderID int
	ForDays     int
	Offset      int
	Limit       int
}
