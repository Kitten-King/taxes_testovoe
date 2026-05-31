package repository

import (
	"context"
	"database/sql"
)

type TaxRepository struct {
	db *sql.DB
}

func NewTaxRepository(db *sql.DB) *TaxRepository {
	return &TaxRepository{db: db}
}

func (r *TaxRepository) GetRateByCityID(ctx context.Context, cityID int) (float64, error) {
	var rate float64
	query := `SELECT rate FROM testovoe.tax_rates WHERE city_id = $1`

	err := r.db.QueryRowContext(ctx, query, cityID).Scan(&rate)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0.10, nil
		}
		return 0, err
	}

	return rate, nil
}
