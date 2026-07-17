package warehouse

import "database/sql"

type Repository interface {
	UpdateStock(itemID string, qty int) error
}

type warehouseRepo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &warehouseRepo{db: db}
}

func (r *warehouseRepo) UpdateStock(itemID string, qty int) error {
	query := `UPDATE stock SET quantity = $1 WHERE id = $2`
	_, err := r.db.Exec(query, qty, itemID)
	return err
}
