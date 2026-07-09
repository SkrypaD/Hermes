# User entity

```JSON
{
  "id"          : "int, primary key",
  "name"        : "string",
  "login"       : "string",
  "created_at"  : "time",
  "is_active"   : "bool",
  "role_id"     : "int, foreign key",
  "password"    : "string"
}
```

```Go
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
```

# User role entity

```JSON
{
  "id"          : "int, primary key",
  "name"        : "string",
  "created_at"  : "time"
}
```

```Go
type UserRoleRepository interface {
	// Attempts to create a new user Role. If role with the same name already exists returns
	// ErrAlreadyTaken Error
	Create(ctx context.Context, role UserRole) (*UserRole, error)
	// Returns all the existing roles
	GetAll(ctx context.Context) ([]UserRole, error)
	// Searches for the role by id. If no role found returns ErrNotFound
	GetByID(ctx context.Context, roleID int) (*UserRole, error)
}
```

# Request entity

```JSON
{
  "id"          : "int, primary key",
  "title"       : "string",
  "description" : "string",
  "created_at"  : "time",
  "updated_at"  : "time",
  "closed_at"   : "time",
  "dispatcher_id"   : "int, foreign key",
  "responder_id"    : "int, foreign key",
  "request_type_id" : "int, foreign key",
}
```


```Go
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
```

# Request type entity

```JSON
{
  "id"          : "int, primary key",
  "name"        : "string",
  "created_at"  : "time",
  "is_relevant" : "bool",
}
```

```Go
type RequestTypeRepository interface {
	// Attempts to create new request type entity. If entity with such name already
	// exists returns ErrAlreadyTaken
	Create(ctx context.Context, requestType RequestType) (*RequestTyp, error)
	// Returns all the existing request types
	GetAll(ctx context.Context) ([]RequestType, error)
	// Searches for the request type by id. If no type found returns ErrNotFound
	GetByID(ctx context.Context, typeID int) (*RequestType, error)
	// Attempts to deactivate a request type with provided id. If type already inactive returns 0 otherwise 1.
	// If to type found with provided id returns ErrNotFound
	Deactivate(ctx context.Context, typeID int) (int, error)
	// Attempts to activate a request type with provided id. If type already active returns 0 otherwise 1.
	// If to type found with provided id returns ErrNotFound
	Activate(ctx context.Context, typeID int) (int, error)
}
```



# Errors
**ErrAlreadyTaken** - invoked if name is already taken
**ErrNotFound** - invoked if entity does not exist
