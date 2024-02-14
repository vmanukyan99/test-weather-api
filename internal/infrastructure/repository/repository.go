package repository

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/vmanukyan99/test-weather-api/config"
)

type Repository struct {
	pool *pgx.ConnPool
}

func New(c config.DataBase) (*Repository, error) {
	conf := pgx.ConnConfig{
		Host:     c.Host,
		Port:     5432,
		User:     c.User,
		Password: c.Password,
		Database: c.Database,
	}

	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{ConnConfig: conf})
	if err != nil {
		return nil, fmt.Errorf("pgx.NewConnPool: %w", err)
	}

	repo := &Repository{
		pool: pool,
	}

	return repo, nil
}
