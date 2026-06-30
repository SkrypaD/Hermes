package repository

import (
	"context"
	"database/sql"
	"errors"

	"Hermes/internal/domain"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

// Queries user by ID. If found returns User struct without requests.
func (usr_repo *UserRepository) GetByID(ctx context.Context, id domain.ID) (*domain.User, error) {
	query := "SELECT * FROM users WHERE id == ? AND is_active == true;"
	row := usr_repo.Db.QueryRowContext(ctx, query, id)

	user := domain.User{}

	err := row.Scan(&user.ID, &user.Name, &user.CreatedAt, &user.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("No user found.")
		}
		return nil, err
	}

	return &user, nil
}

// Returns all the active users
func (usr_repo *UserRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	query := `SELECT * FROM users WHERE is_active == true;`
	rows, err := usr_repo.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	users := []domain.User{}

	for rows.Next() {
		user := domain.User{}

		err := rows.Scan(&user.ID, &user.Name, &user.CreatedAt, &user.Type)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil

}

func (usr_repo *UserRepository) GetByType(ctx context.Context, user_type string) ([]domain.User, error) {
	query := `SELECT * FROM users WHERE is_active == true AND type == ?`

	rows, err := usr_repo.Db.QueryContext(ctx, query, user_type)
	if err != nil {
		return nil, err
	}

	users := []domain.User{}

	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.CreatedAt, &user.Type)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (usr_repo *UserRepository) Create(ctx context.Context, user domain.User) error {
	query := `INSERT INTO users (id, name, created_at, type, is_active) VALUES (?, ?, ?, ?, ?)`

	_, err := usr_repo.Db.ExecContext(ctx, query, user.ID, user.Name, user.CreatedAt, user.Type, user.IsActive)

	if err != nil {
		return err
	}

	return nil
}

func (usr_repo *UserRepository) GetResponders(ctx context.Context, days_period int) ([]domain.User, error) {
	panic("Not implemented")
}
