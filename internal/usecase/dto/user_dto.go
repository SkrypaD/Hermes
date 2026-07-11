package dto

import "time"

type LoginUser struct {
	Login    string
	Password string
}

type ReturnUser struct {
	Id       int
	Login    string
	Name     string
	Role     string
	RoleID   int
	CreateAt time.Time
	Requests []ReturnRequest
}

type ReturnUserCredentials struct {
	Id        int
	Login     string
	Name      string
	Role      string
	RoleID    int
	CreatedAt time.Time
	Token     string
}

type CreateUser struct {
	Login    string
	Name     string
	Password string
	RoleID   int
}

// ===========================//
type ReturnRole struct {
	ID   int
	Name string
}

type CreateRole struct {
	Name string
}
