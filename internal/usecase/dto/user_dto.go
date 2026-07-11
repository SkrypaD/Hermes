package dto

import "time"

type LoginUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ReturnUser struct {
	ID       int             `json:"id"`
	Login    string          `json:"login"`
	Name     string          `json:"name"`
	Role     string          `json:"role"`
	RoleID   int             `json:"role_id"`
	CreateAt time.Time       `json:"created_at"`
	Requests []ReturnRequest `json:"requests"`
}

type ReturnUserCredentials struct {
	ID        int       `json:"id"`
	Login     string    `json:"login"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	RoleID    int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `json:"token"`
}

type CreateUser struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
}

// ===========================//
type ReturnRole struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateRole struct {
	Name string `json:"name"`
}
