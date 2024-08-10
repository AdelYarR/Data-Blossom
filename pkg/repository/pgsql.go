package repository

import (
	"context"

	"github.com/AdelYarR/Data-Blossom/pkg/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PGRepo struct {
	pool *pgxpool.Pool
}

func New(cfg *config.Config) (*PGRepo, error) {
	pool, err := pgxpool.Connect(context.Background(), cfg.Store.DBurl)
	if err != nil {
		return nil, err
	}

	return &PGRepo{
		pool: pool,
	}, nil
}
