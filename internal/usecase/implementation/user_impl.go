package implementation

import (
	"Hermes/internal/domain"
	"Hermes/internal/usecase/dto"
	"Hermes/internal/usecase/services"
	"context"
)

// TODO: probably need to add some kind of config struct that would hold JWT config stuff
type UserUsecase struct {
	UsrRepo      domain.UserRepository
	UsrRoleRepo  domain.UserRoleRepository
	ReqRepo      domain.RequestRepository
	HashService  services.HasherService
	TokenService services.TokenGeneratorService
}

func NewUserUsecase(usrRepo domain.UserRepository, usrRoleRepo domain.UserRoleRepository, reqRepo domain.RequestRepository,
	tokenServ services.TokenGeneratorService, hashServ services.HasherService) *UserUsecase {
	return &UserUsecase{
		UsrRepo:      usrRepo,
		UsrRoleRepo:  usrRoleRepo,
		ReqRepo:      reqRepo,
		HashService:  hashServ,
		TokenService: tokenServ,
	}
}

// TOTHINK: Do I actually need to query while loggin in. I mean whouldnt it be better to split it into get user by login and then pass values to Login then all login has to do is just check password and create token
// Attempts to login users to the server.
// Checks if provided login exists if not returns ErrNotFound.
// Checks if invalid password provided returns ErrNotFound.   <- Same error here in case someone would want to look for logins
// Returns user entity with JWT token.
func (usr_case *UserUsecase) Login(ctx context.Context, loginCreds dto.LoginUser) (*dto.ReturnUserCredentials, error) {
	user, err := usr_case.UsrRepo.GetByLogin(ctx, loginCreds.Login, true)
	if err != nil {
		return nil, err
	}

	role, err := usr_case.UsrRoleRepo.GetByID(ctx, user.RoleID)
	if err != nil {
		return nil, err
	}

	// TODO: Has to be some kind of arror to mark invalid password
	err = usr_case.HashService.VerifyPassword(user.Password, loginCreds.Password)
	if err != nil {
		return nil, err
	}

	// TODO: create a struct to hold JWT config fields
	token, err := usr_case.TokenService.GenerateToken(services.CustomClaims{
		UserID: user.ID,
		RoleID: user.RoleID,
		Role:   role.Name,
		Login:  user.Login,
	})
	if err != nil {
		return nil, err
	}

	return &dto.ReturnUserCredentials{
		ID:        user.ID,
		Name:      user.Name,
		Login:     user.Login,
		Role:      role.Name,
		RoleID:    user.RoleID,
		CreatedAt: user.CreatedAt,
		Token:     token,
	}, nil
}

// Attempts to create a new user.
// If provided login already taken returns ErrAlreadyTaken.
// If role does not exist returns ErrNotFound.
func (usr_case *UserUsecase) Create(ctx context.Context, newUserEntity dto.CreateUser) (*dto.ReturnUser, error) {
	role, err := usr_case.UsrRoleRepo.GetByID(ctx, newUserEntity.RoleID)
	if err != nil {
		return nil, err
	}

	createEnt := domain.User{
		Name:     newUserEntity.Name,
		Login:    newUserEntity.Login,
		Password: newUserEntity.Password,
		RoleID:   newUserEntity.RoleID,
	}

	user, err := usr_case.UsrRepo.Create(ctx, createEnt)
	if err != nil {
		return nil, err
	}

	return &dto.ReturnUser{
		ID:       user.ID,
		Login:    user.Login,
		Name:     user.Name,
		Role:     role.Name,
		RoleID:   user.RoleID,
		CreateAt: user.CreatedAt,
		Requests: nil,
	}, nil
}

// Attempts to find user with provided ID.
// If no user found returns ErrNotFound.
// (If onlyActive flag is positive would return user only if active)
func (usr_case *UserUsecase) GetByID(ctx context.Context, ID int, onlyActive bool) (*dto.ReturnUser, error) {
	user, err := usr_case.UsrRepo.GetByID(ctx, ID, onlyActive)

	if err != nil {
		return nil, err
	}

	role, err := usr_case.UsrRoleRepo.GetByID(ctx, user.RoleID)
	if err != nil {
		return nil, err
	}

	rtUsr := dto.ReturnUser{
		ID:       user.ID,
		Login:    user.Login,
		Name:     user.Name,
		Role:     role.Name,
		RoleID:   user.RoleID,
		CreateAt: user.CreatedAt,
		Requests: nil,
	}

	return &rtUsr, nil
}

// Attempts to find user with provided Login.
// If no user found returns ErrNotFound.
// (If onlyActive flag is positive would return user only if active)
func (usr_case *UserUsecase) GetByLogin(ctx context.Context, login string, onlyActive bool) (*dto.ReturnUser, error) {
	user, err := usr_case.UsrRepo.GetByLogin(ctx, login, onlyActive)

	if err != nil {
		return nil, err
	}

	role, err := usr_case.UsrRoleRepo.GetByID(ctx, user.RoleID)
	if err != nil {
		return nil, err
	}

	rtUsr := dto.ReturnUser{
		ID:       user.ID,
		Login:    user.Login,
		Name:     user.Name,
		Role:     role.Name,
		RoleID:   user.RoleID,
		CreateAt: user.CreatedAt,
		Requests: nil,
	}

	return &rtUsr, nil
}

// Attempts to fetch all the existing users.
// If 'onlyActive' flag is true returns only currently active users.
func (usr_case *UserUsecase) GetAll(ctx context.Context, onlyActive bool) ([]dto.ReturnUser, error) {
	users, err := usr_case.UsrRepo.GetAll(ctx, onlyActive)
	if err != nil {
		return nil, err
	}

	roles, err := usr_case.UsrRoleRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	rolesMap := make(map[int]string)
	for _, r := range roles {
		rolesMap[r.ID] = r.Name
	}

	rtUsrs := make([]dto.ReturnUser, 0, len(users))

	for _, usr := range users {
		rtUsr := dto.ReturnUser{
			ID:       usr.ID,
			Login:    usr.Login,
			Name:     usr.Name,
			Role:     rolesMap[usr.ID],
			RoleID:   usr.RoleID,
			CreateAt: usr.CreatedAt,
			Requests: nil,
		}
		rtUsrs = append(rtUsrs, rtUsr)
	}

	return rtUsrs, nil
}

// Attempts to active user with provided ID.
// If no user found for provided ID returns ErrNotFound.
// If user already active returns 0 otherwise 1.
func (usr_case *UserUsecase) Activate(ctx context.Context, ID int) (int, error) {
	return usr_case.UsrRepo.Activate(ctx, ID)
}

// Attempts to deactive user with provided ID.
// If no user found for provided ID returns ErrNotFound.
// If user already inactive returns 0 otherwise 1.
func (usr_case *UserUsecase) Deactivate(ctx context.Context, ID int) (int, error) {
	return usr_case.UsrRepo.Deactivate(ctx, ID)
}

// Fetches all the users of a reponder type.
// If 'onlyActive' flag is true returns only currently active users.
func (usr_case *UserUsecase) GetAllResponders(ctx context.Context, onlyActive bool) ([]dto.ReturnUser, error) {
	users, err := usr_case.UsrRepo.GetAllResponders(ctx, onlyActive)
	if err != nil {
		return nil, err
	}

	roles, err := usr_case.UsrRoleRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	rolesMap := make(map[int]string)
	for _, r := range roles {
		rolesMap[r.ID] = r.Name
	}

	rtUsrs := make([]dto.ReturnUser, 0, len(users))

	for _, usr := range users {
		rtUsr := dto.ReturnUser{
			ID:       usr.ID,
			Login:    usr.Login,
			Name:     usr.Name,
			Role:     rolesMap[usr.ID],
			RoleID:   usr.RoleID,
			CreateAt: usr.CreatedAt,
			Requests: nil,
		}
		rtUsrs = append(rtUsrs, rtUsr)
	}

	return rtUsrs, nil

}

// Attempts to find user role by provided ID.
// If no role found returns ErrNotFound.
func (usr_case *UserUsecase) GetRoleByID(ctx context.Context, ID int) (*dto.ReturnRole, error) {
	role, err := usr_case.UsrRoleRepo.GetByID(ctx, ID)

	if err != nil {
		return nil, err
	}

	return &dto.ReturnRole{
		ID:   role.ID,
		Name: role.Name,
	}, err
}

// Returns all the existing roles.
func (usr_case *UserUsecase) GetAllRoles(ctx context.Context) ([]dto.ReturnRole, error) {
	roles, err := usr_case.UsrRoleRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	rtRoles := make([]dto.ReturnRole, 0, len(roles))

	for _, role := range roles {
		rtRole := dto.ReturnRole{
			ID:   role.ID,
			Name: role.Name,
		}

		rtRoles = append(rtRoles, rtRole)
	}

	return rtRoles, nil
}

// Attempts to create a new role.
// If role with provided name exists returns ErrAlreadyTaken.
func (usr_case *UserUsecase) CreateRole(ctx context.Context, createEntity dto.CreateRole) (*dto.ReturnRole, error) {
	creationEnt := domain.UserRole{
		Name: createEntity.Name,
	}

	newRole, err := usr_case.UsrRoleRepo.Create(ctx, creationEnt)
	if err != nil {
		return nil, err
	}

	return &dto.ReturnRole{
		ID:   newRole.ID,
		Name: newRole.Name,
	}, nil
}
