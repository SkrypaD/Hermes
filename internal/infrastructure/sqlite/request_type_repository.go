package sqlite

import (
	"Hermes/internal/domain"
	"context"
	"database/sql"
	"strings"
)

type RequestTypeRepository struct {
	DB *sql.DB
}

func (req_repo *RequestTypeRepository) Create(ctx context.Context, requestType domain.RequestType) (*domain.RequestType, error) {
	query := `INSERT INTO request_types (name) VALUES (?) RETURNING id, is_relevant, created_at`

	err := req_repo.DB.QueryRowContext(ctx, query, requestType.Name).Scan(&requestType.ID, &requestType.IsRelevant, &requestType.CreatedAt)

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return nil, domain.ErrAlreadyTaken
		}
		return nil, err
	}

	return &requestType, nil
}

func (req_repo *RequestTypeRepository) GetAll(ctx context.Context, onlyActive bool) ([]domain.RequestType, error) {
	query := `SELECT id, name, is_relevant, created_at FROM request_types`
	if onlyActive {
		query += ` WHERE is_relevant = 1`
	}

	rows, err := req_repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requestTypes []domain.RequestType
	for rows.Next() {
		var requestType domain.RequestType

		err = rows.Scan(&requestType.ID, &requestType.Name, &requestType.IsRelevant, &requestType.CreatedAt)
		if err != nil {
			return nil, err
		}

		requestTypes = append(requestTypes, requestType)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return requestTypes, nil
}

func (req_repo *RequestTypeRepository) GetByID(ctx context.Context, typeID int, isActive bool) (*domain.RequestType, error) {
	query := `SELECT id, name, is_relevant, created_at FROM request_types WHERE `
	if isActive {
		query += `is_relevant = 1 AND `
	}
	query += `id = ?`

	row := req_repo.DB.QueryRowContext(ctx, query, typeID)

	var requestType domain.RequestType
	err := row.Scan(&requestType.ID, &requestType.Name, &requestType.IsRelevant, &requestType.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return &requestType, nil
}

func (req_repo *RequestTypeRepository) Deactivate(ctx context.Context, typeID int) (int, error) {
	query := `UPDATE request_types SET is_relevant = 0 WHERE id = ? AND is_relevant = 1`

	res, err := req_repo.DB.ExecContext(ctx, query, typeID)
	if err != nil {
		return -1, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		var exists bool
		checkQuery := `SELECT EXISTS(SELECT 1 FROM request_types WHERE id = ?)`
		err := req_repo.DB.QueryRowContext(ctx, checkQuery, typeID).Scan(&exists)

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

func (req_repo *RequestTypeRepository) Activate(ctx context.Context, typeID int) (int, error) {
	query := `UPDATE request_types SET is_relevant = 1 WHERE id = ? AND is_relevant = 0`

	res, err := req_repo.DB.ExecContext(ctx, query, typeID)
	if err != nil {
		return -1, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		var exists bool
		checkQuery := `SELECT EXISTS(SELECT 1 FROM request_types WHERE id = ?)`
		err := req_repo.DB.QueryRowContext(ctx, checkQuery, typeID).Scan(&exists)

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
