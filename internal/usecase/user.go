package usecase

import (
	"Hermes/internal/usecase/dto"
	"context"
)

type UserUsecase interface {
	// Attempts to login users to the server.
	// Checks if provided login exists if not returns ErrNotFound.
	// Checks if invalid password provided returns ErrNotFound.   <- Same error here in case someone would want to look for logins
	// Returns user entity with JWT token.
	Login(ctx context.Context, loginCreds dto.LoginUser) (*dto.ReturnUserCredentials, error)
	// Attempts to create a new user.
	// If provided login already taken returns ErrAlreadyTaken.
	// If role does not exist returns ErrNotFound.
	Create(ctx context.Context, newUserEntity dto.CreateUser) (*dto.ReturnUser, error)
	// Attempts to find user with provided ID.
	// If no user found returns ErrNotFound.
	// (If onlyActive flag is positive would return user only if active)
	GetByID(ctx context.Context, ID int, onlyActive bool) (*dto.ReturnUser, error)
	// Attempts to find user with provided Login.
	// If no user found returns ErrNotFound.
	// (If onlyActive flag is positive would return user only if active)
	GetByLogin(ctx context.Context, login string, onlyActive bool) (*dto.ReturnUser, error)
	// Attempts to fetch all the existing users.
	// If 'onlyActive' flag is true returns only currently active users.
	GetAll(ctx context.Context, onlyActive bool) ([]dto.ReturnUser, error)
	// Attempts to active user with provided ID.
	// If no user found for provided ID returns ErrNotFound.
	// If user already active returns 0 otherwise 1.
	Activate(ctx context.Context, ID int) (int, error)
	// Attempts to deactive user with provided ID.
	// If no user found for provided ID returns ErrNotFound.
	// If user already inactive returns 0 otherwise 1.
	Deactivate(ctx context.Context, ID int) (int, error)
	// Fetches all the users of a reponder type.
	// If 'onlyActive' flag is true returns only currently active users.
	GetAllResponders(ctx context.Context, onlyActive bool) ([]dto.ReturnUser, error)

	// Attempts to find user role by provided ID.
	// If no role found returns ErrNotFound.
	GetRoleByID(ctx context.Context, ID int) (*dto.ReturnRole, error)
	// Returns all the existing roles.
	GetAllRoles(ctx context.Context) ([]dto.ReturnRole, error)
	// Attempts to create a new role.
	// If role with provided name exists returns ErrAlreadyTaken.
	CreateRole(ctx context.Context, createEntity dto.CreateRole) (*dto.ReturnRole, error)
}
