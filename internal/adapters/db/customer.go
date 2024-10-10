package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Odvin/go-live-reloading/internal/application/domain"
)

func (a DbAdapter) CreateCustomer(customer *domain.Customer) error {
	query := `
		INSERT INTO staff.customer (title, active)
		VALUES ($1, $2)
		RETURNING id, created, updated, version;`

	args := []any{customer.Title, customer.Active}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return a.DB.QueryRowContext(ctx, query, args...).Scan(
		&customer.ID,
		&customer.Created,
		&customer.Updated,
		&customer.Version,
	)
}

func (a DbAdapter) GetCustomer(id int64) (*domain.Customer, error) {
	if id < 1 {
		return nil, domain.ErrRecordNotFound
	}

	query := `
		SELECT id, title, created, updated, active, version
		FROM staff.customer
		WHERE id = $1;`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var customer domain.Customer

	err := a.DB.QueryRowContext(ctx, query, id).Scan(
		&customer.ID,
		&customer.Title,
		&customer.Created,
		&customer.Updated,
		&customer.Active,
		&customer.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, domain.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &customer, nil
}
