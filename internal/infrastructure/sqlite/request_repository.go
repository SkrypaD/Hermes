package sqlite

import (
	"Hermes/internal/domain"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type RequestRepository struct {
	DB *sql.DB
}

// type Request struct {
// 	ID            int
// 	Title         string
// 	Description   string
// 	CreatedAt     time.Time
// 	UpdatedAt     *time.Time
// 	ClosedAt      *time.Time
// 	DispatcherID  int
// 	ResponderID   int
// 	RequestTypeID int
// }

func (req_repo *RequestRepository) GetAll(ctx context.Context, forDays int, limit int, offset int) ([]domain.Request, error) {
	query := `SELECT id, title, description, created_at, updated_at, closed_at, dispatcher_id, responder_id, request_type_id FROM requests`

	var conditions []string
	var args []any

	if forDays > 0 {
		conditions = append(conditions, `created_at >= datetime('now', ?)`)
		args = append(args, fmt.Sprintf("-%d days", forDays))
	}

	if len(conditions) > 0 {
		query += ` WHERE ` + strings.Join(conditions, " AND ")
	}

	query += ` ORDER BY created_at DESC`

	if limit > 0 {
		query += fmt.Sprintf(` LIMIT %d`, limit)

		if offset > 0 {
			query += fmt.Sprintf(` OFFSET %d`, offset)
		}
	}

	rows, err := req_repo.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []domain.Request
	if limit > 0 {
		requests = make([]domain.Request, 0, limit)
	}

	for rows.Next() {
		var request domain.Request
		var updated_at, closed_at sql.NullTime

		err = rows.Scan(&request.ID, &request.Title, &request.Description, &request.CreatedAt, &updated_at, &closed_at, &request.DispatcherID, &request.ResponderID, &request.RequestTypeID)
		if err != nil {
			return nil, err
		}

		if updated_at.Valid {
			request.UpdatedAt = &updated_at.Time
		}
		if closed_at.Valid {
			request.ClosedAt = &closed_at.Time
		}

		requests = append(requests, request)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return requests, nil
}

func (req_repo *RequestRepository) GetForResponder(ctx context.Context, responderID int, forDays int, limit int, offset int) ([]domain.Request, error) {
	query := `SELECT id, title, description, created_at, updated_at, closed_at, dispatcher_id, responder_id, request_type_id FROM requests`

	var conditions []string
	var args []any

	if responderID < 0 {
		return nil, domain.ErrInvalidOperation
	}

	conditions = append(conditions, ` responder_id = ? `)
	args = append(args, responderID)

	if forDays > 0 {
		conditions = append(conditions, ` created_at >= datetime('now', ?) `)
		args = append(args, fmt.Sprintf("-%d days", forDays))
	}

	if len(conditions) > 0 {
		query += ` WHERE ` + strings.Join(conditions, " AND ")
	}

	query += ` ORDER BY created_at DESC`

	if limit > 0 {
		query += fmt.Sprintf(` LIMIT %d`, limit)

		if offset > 0 {
			query += fmt.Sprintf(` OFFSET %d`, offset)
		}
	}

	rows, err := req_repo.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []domain.Request
	if limit > 0 {
		requests = make([]domain.Request, 0, limit)
	}

	for rows.Next() {
		var request domain.Request
		var updated_at, closed_at sql.NullTime

		err = rows.Scan(&request.ID, &request.Title, &request.Description, &request.CreatedAt, &updated_at, &closed_at, &request.DispatcherID, &request.ResponderID, &request.RequestTypeID)
		if err != nil {
			return nil, err
		}

		if updated_at.Valid {
			request.UpdatedAt = &updated_at.Time
		}
		if closed_at.Valid {
			request.ClosedAt = &closed_at.Time
		}

		requests = append(requests, request)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return requests, nil
}

func (req_repo *RequestRepository) GetByID(ctx context.Context, ID int) (*domain.Request, error) {
	query := `SELECT id, title, description, created_at, updated_at, closed_at, dispatcher_id, responder_id, request_type_id FROM requests
	WHERE id = ?`

	row := req_repo.DB.QueryRowContext(ctx, query, ID)

	var request domain.Request
	var updated_at, closed_at sql.NullTime

	err := row.Scan(&request.ID, &request.Title, &request.Description, &request.CreatedAt, &updated_at, &closed_at, &request.DispatcherID, &request.ResponderID, &request.RequestTypeID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	if updated_at.Valid {
		request.UpdatedAt = &updated_at.Time
	}
	if closed_at.Valid {
		request.ClosedAt = &closed_at.Time
	}

	return &request, nil

}

func (req_repo *RequestRepository) Create(ctx context.Context, request domain.Request) (*domain.Request, error) {
	query := `INSERT INTO requests (title, description, dispatcher_id, responder_id, request_type_id) VALUES (?, ?, ?, ?, ?)
	RETURNING id, created_at`

	err := req_repo.DB.QueryRowContext(ctx, query, request.Title, request.Description, request.DispatcherID, request.ResponderID, request.RequestTypeID).Scan(&request.ID, &request.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &request, nil
}

func (req_repo *RequestRepository) Close(ctx context.Context, ID int) (int, error) {
	query := `UPDATE requests SET closed_at = CURRENT_TIMESTAMP WHERE id = ? AND closed_at IS NULL`

	res, err := req_repo.DB.ExecContext(ctx, query, ID)
	if err != nil {
		return -1, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		var exists bool
		checkQuery := `SELECT EXISTS(SELECT 1 FROM requests WHERE id = ?)`
		err := req_repo.DB.QueryRowContext(ctx, checkQuery, ID).Scan(&exists)

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

// TODO: still think that this approach is bad. Need to rewrite it 100%.
func (req_repo *RequestRepository) Update(ctx context.Context, request domain.Request) (*domain.Request, error) {
	existing, err := req_repo.GetByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	if request.Title != "" {
		existing.Title = request.Title
	}
	if request.Description != "" {
		existing.Description = request.Description
	}
	if request.DispatcherID > 0 {
		existing.DispatcherID = request.DispatcherID
	}
	if request.ResponderID > 0 {
		existing.ResponderID = request.ResponderID
	}
	if request.RequestTypeID > 0 {
		existing.RequestTypeID = request.RequestTypeID
	}

	query := `UPDATE requests
	SET title = ?, description = ?, dispatcher_id = ?, responder_id = ?, request_type_id = ?, updated_at = CURRENT_TIMESTAMP
	WHERE id = ?
	RETURNING updated_at`

	var updatedAt sql.NullTime
	err = req_repo.DB.QueryRowContext(ctx, query,
		existing.Title,
		existing.Description,
		existing.DispatcherID,
		existing.ResponderID,
		existing.RequestTypeID,
		existing.ID,
	).Scan(&updatedAt)

	if err != nil {
		return nil, err
	}

	if updatedAt.Valid {
		existing.UpdatedAt = &updatedAt.Time
	}

	return existing, nil
}
