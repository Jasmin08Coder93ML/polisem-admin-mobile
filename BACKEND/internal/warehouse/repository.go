package warehouse

import (
	"context"
	"database/sql"
)

type Repository interface {
	UpdateStock(ctx context.Context, itemID string, qty int) error
}

type warehouseRepo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &warehouseRepo{db: db}
}

func (r *warehouseRepo) UpdateStock(ctx context.Context, itemID string, qty int) error {
	query := `UPDATE stock SET quantity = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, qty, itemID)
	return err
}
