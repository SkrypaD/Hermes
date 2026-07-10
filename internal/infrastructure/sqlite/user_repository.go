package sqlite

import (
	"Hermes/internal/domain"
	"context"
	"database/sql"
	"strings"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (usr_repo *UserRepository) GetAll(ctx context.Context, onlyActive bool) ([]domain.User, error) {
	query := `SELECT id, name, login, role_id, created_at, is_active, password FROM users `
	if onlyActive == true {
		query += `WHERE is_active = 1`
	}

	query += ``

	rows, err := usr_repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var tempUsr domain.User

		err = rows.Scan(&tempUsr.ID, &tempUsr.Name, &tempUsr.Login, &tempUsr.RoleID, &tempUsr.CreatedAt, &tempUsr.IsActive, &tempUsr.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, tempUsr)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (usr_repo *UserRepository) GetAllResponders(ctx context.Context, onlyActive bool) ([]domain.User, error) {
	query := `SELECT id, name, login, role_id, created_at, is_active, password FROM users WHERE `
	if onlyActive == true {
		query += `is_active = 1 AND `
	}

	query += `role_id = 1`

	rows, err := usr_repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var tempUsr domain.User

		err = rows.Scan(&tempUsr.ID, &tempUsr.Name, &tempUsr.Login, &tempUsr.RoleID, &tempUsr.CreatedAt, &tempUsr.IsActive, &tempUsr.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, tempUsr)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (usr_repo *UserRepository) GetByLogin(ctx context.Context, login string, onlyActive bool) (*domain.User, error) {
	query := `SELECT id, name, login, role_id, created_at, is_active, password FROM users WHERE `
	if onlyActive {
		query += `is_active = 1 AND `
	}
	query += `login = ?`

	row := usr_repo.DB.QueryRowContext(ctx, query, login)
	var user domain.User

	err := row.Scan(&user.ID, &user.Name, &user.Login, &user.RoleID, &user.CreatedAt, &user.IsActive, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (usr_repo *UserRepository) GetByID(ctx context.Context, ID int, onlyActive bool) (*domain.User, error) {
	query := `SELECT id, name, login, role_id, created_at, is_active, password FROM users WHERE `
	if onlyActive {
		query += `is_active = 1 AND `
	}
	query += `id = ?`

	row := usr_repo.DB.QueryRowContext(ctx, query, ID)
	var user domain.User

	err := row.Scan(&user.ID, &user.Name, &user.Login, &user.RoleID, &user.CreatedAt, &user.IsActive, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (usr_repo *UserRepository) Create(ctx context.Context, user domain.User) (*domain.User, error) {
	query := `INSERT INTO users (name, login, password, role_id) VALUES (?, ?, ?, ?)
	RETURNING id, is_active, created_at`

	err := usr_repo.DB.QueryRowContext(ctx, query, user.Name, user.Login, user.Password, user.RoleID).Scan(&user.ID, &user.IsActive, &user.CreatedAt)

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return nil, domain.ErrAlreadyTaken
		}
		return nil, err
	}

	return &user, nil
}

func (usr_repo *UserRepository) Deactivate(ctx context.Context, ID int) (int, error) {
	query := `UPDATE users SET is_active = 0 WHERE id = ? AND is_active = 1`

	res, err := usr_repo.DB.ExecContext(ctx, query, ID)

	if err != nil {
		return -1, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		var exists bool
		checkQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)`
		err := usr_repo.DB.QueryRowContext(ctx, checkQuery, ID).Scan(&exists)

		if err != nil {
			return 0, err
		}

		if !exists {
			return -1, domain.ErrNotFound
		}

		return 0, nil
	}

	return 1, nil
}

// TODO: In such cases with more then one db request at a time use transaction.
func (usr_repo *UserRepository) Activate(ctx context.Context, ID int) (int, error) {
	query := `UPDATE users SET is_active = 1 WHERE id == ? AND is_active == 0`

	res, err := usr_repo.DB.ExecContext(ctx, query, ID)

	if err != nil {
		return -1, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		var exists bool
		checkQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)`
		err := usr_repo.DB.QueryRowContext(ctx, checkQuery, ID).Scan(&exists)

		if err != nil {
			return 0, err
		}

		if !exists {
			return -1, domain.ErrNotFound
		}

		return 0, nil
	}

	return 1, nil
}
