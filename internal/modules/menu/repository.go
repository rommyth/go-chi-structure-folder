package menu

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

var (
	ErrMenuNotFound = errors.New("menu not found")
)

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) GetByID(ctx context.Context, id int) (Menu, error) {
	menu := Menu{
		ID:        int64(id),
		Name:      "Menu 1",
		Category:  "Breakfast",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	return menu, nil
}
