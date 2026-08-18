package user

import "github.com/jackc/pgx/v5/pgxpool"

type Repository interface{}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}
