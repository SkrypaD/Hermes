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
	panic("Not implemented")
}

func (req_repo *RequestRepository) GetByID(ctx context.Context, ID int) (*domain.Request, error) {
	panic("Not implemented")
}

func (req_repo *RequestRepository) Create(ctx context.Context, request domain.Request) (*domain.Request, error) {
	panic("Not implemented")
}

func (req_repo *RequestRepository) Close(ctx context.Context, ID int) (int, error) {
	panic("Not implemented")
}

func (req_repo *RequestRepository) Update(ctx context.Context, request domain.Request) (*domain.Request, error) {
	panic("Not implemented")
}
