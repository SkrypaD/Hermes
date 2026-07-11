package sqlite

import (
	"Hermes/internal/domain"
	"context"
	"database/sql"
	"strings"
)

type UserRoleRepository struct {
	DB *sql.DB
}

func (usr_repo *UserRoleRepository) Create(ctx context.Context, role domain.UserRole) (*domain.UserRole, error) {
	query := `INSERT INTO user_roles (name) VALUES (?)
	RETURNING id, created_at`

	err := usr_repo.DB.QueryRowContext(ctx, query, role.Name).Scan(&role.ID, &role.CreatedAt)

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return nil, domain.ErrAlreadyTaken
		}
		return nil, err
	}

	return &role, nil

}
func (usr_repo *UserRoleRepository) GetAll(ctx context.Context) ([]domain.UserRole, error) {
	query := `SELECT id, name, created_at FROM user_roles`

	rows, err := usr_repo.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []domain.UserRole

	for rows.Next() {
		var role domain.UserRole

		if err := rows.Scan(&role.ID, &role.Name, &role.CreatedAt); err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (usr_repo *UserRoleRepository) GetByID(ctx context.Context, roleID int) (*domain.UserRole, error) {
	query := `SELECT id, name, created_at FROM user_roles WHERE id = ?`

	row := usr_repo.DB.QueryRowContext(ctx, query, roleID)
	var role domain.UserRole

	err := row.Scan(&role.ID, &role.Name, &role.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return &role, nil
}
