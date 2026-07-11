package dto

import "time"

type ReturnRequest struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	RequestType   string    `json:"request_type"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	ClosedAt      time.Time `json:"closed_at"`
	DispatcherID  int       `json:"dispatcher_id"`
	ResponderID   int       `json:"responder_id"`
	RequestTypeID int       `json:"request_type_id"`
}

type CreateRequest struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	DispatcherID  int    `json:"dispatcher_id"`
	ResponderID   int    `json:"responder_id"`
	RequestTypeID int    `json:"request_type_id"`
}

type UpdateRequest struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	ResponderID   int    `json:"responder_id"`
	RequestTypeID int    `json:"request_type_id"`
}

type ReturnRequestType struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	IsRelevant bool      `json:"is_relevant"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateRequestType struct {
	Name string `json:"name"`
}

//====================================//

type ActivateRequestType struct {
	ID int `json:"id"`
}

type GetAllRequests struct {
	ForDays int `json:"for_days"`
	Offset  int `json:"offset"`
	Limit   int `json:"limit"`
}

type GetForResponder struct {
	ResponderID int `json:"responder_id"`
	ForDays     int `json:"for_days"`
	Offset      int `json:"offset"`
	Limit       int `json:"limit"`
}
