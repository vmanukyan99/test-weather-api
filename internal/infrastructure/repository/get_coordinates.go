package repository

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/pgtype"
)

type get struct {
	latitude  pgtype.Float8
	longitude pgtype.Float8
}

func (r *Repository) GetCoordinates(country, city string) (float64, float64, error) {
	query := goqu.Select("latitude", "longitude").
		From("coordinates").
		Where(
			goqu.And(
				goqu.C("country").Eq(country),
				goqu.C("city").Eq(city),
			),
		)

	sql, params, err := query.ToSQL()
	if err != nil {
		return 0, 0, fmt.Errorf("query.ToSQL: %w", err)
	}

	g := &get{}

	dest := []any{
		&g.latitude,
		&g.longitude,
	}

	row := r.pool.QueryRow(sql, params...)

	err = row.Scan(dest...)
	if err != nil {
		return 0, 0, fmt.Errorf("row.Scan: %w", err)
	}

	return g.latitude.Float, g.longitude.Float, nil
}
