package domain

import (
	"context"
	"time"
)

type User struct {
	ID        int
	Name      string
	Login     string
	Password  string
	IsActive  bool
	RoleID    int
	CreatedAt time.Time
}

type UserRepository interface {
	// Get all the users (accepts argument 'onlyActive' if true returns only active users)
	GetAll(ctx context.Context, onlyActive bool) ([]User, error)
	// Get all the responder type users (accepts argument 'onlyActive' if true returns only active users)
	GetAllResponders(ctx context.Context, onlyActive bool) ([]User, error)
	// Searchers for users by its login returns user entity if exists else returns ErrNotFound
	// (If onlyActive flag is positive would return user only if active)
	GetByLogin(ctx context.Context, login string, onlyActive bool) (*User, error)
	// Searchers for users by its ID returns user entity if exists else returns ErrNotFound
	// (If onlyActive flag is positive would return user only if active)
	GetByID(ctx context.Context, ID int, onlyActive bool) (*User, error)
	// Accepts user struct and attempts to create a new user. If the login or name is already taken
	// returns ErrAlreadyTaken
	Create(ctx context.Context, user User) (*User, error)
	// Attempts to deactive user entity, if entity not found returns ErrNotFound.
	// If it was inactive returns 0 if deactivated returns 1
	Deactivate(ctx context.Context, ID int) (int, error)
	// Attempts to active user entity, if entity not found returns ErrNotFound.
	// If it was active returns 0 if activated returns 1
	Activate(ctx context.Context, ID int) (int, error)
}

type UserRole struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

type UserRoleRepository interface {
	// Attempts to create a new user Role. If role with the same name already exists returns
	// ErrAlreadyTaken Error
	Create(ctx context.Context, role UserRole) (*UserRole, error)
	// Returns all the existing roles
	GetAll(ctx context.Context) ([]UserRole, error)
	// Searches for the role by id. If no role found returns ErrNotFound
	GetByID(ctx context.Context, roleID int) (*UserRole, error)
}
