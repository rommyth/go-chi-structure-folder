package food

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

var (
	ErrFoodNotFound = errors.New("food not found")
)

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) GetByID(ctx context.Context, id int) (Food, error) {
	query := `
		SELECT id, menu_id, name, price, image, created_at, updated_at
		FROM foods
		WHERE id = $1
	`

	var food Food
	err := r.db.QueryRow(ctx, query, id).Scan(
		&food.ID,
		&food.MenuID,
		&food.Name,
		&food.Price,
		&food.Image,
		&food.CreatedAt,
		&food.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Food{}, ErrFoodNotFound
		}
		return Food{}, fmt.Errorf("get food: %w", err)
	}

	return food, nil
}

func (r *repository) List(ctx context.Context, page, limit int, search string) ([]Food, int, error) {
	offset := (page - 1) * limit

	var conditions []string
	var args []interface{}
	argIdx := 1

	if search != "" {
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", argIdx))
		args = append(args, "%"+search+"%")
		argIdx++
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	// Count all Foods
	countQuery := fmt.Sprintf(`
			SELECT COUNT(*)
			FROM foods
			%s
		`, whereClause)

	var total int
	err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return []Food{}, 0, fmt.Errorf("count foods: %w", err)
	}

	// main query
	query := fmt.Sprintf(`
			SELECT id, menu_id, name, price, image, created_at, updated_at
			FROM foods
			%s
			ORDER BY name ASC
			LIMIT $%d OFFSET $%d
		`, whereClause, argIdx, argIdx+1)

	args = append(args, limit, offset)

	var foods []Food
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return []Food{}, 0, fmt.Errorf("get foods: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var food Food
		err := rows.Scan(
			&food.ID,
			&food.MenuID,
			&food.Name,
			&food.Price,
			&food.Image,
			&food.CreatedAt,
			&food.UpdatedAt,
		)
		if err != nil {
			return []Food{}, 0, fmt.Errorf("scan food: %w", err)
		}

		foods = append(foods, food)
	}

	return foods, total, nil
}

func (r *repository) Create(ctx context.Context, f Food) (Food, error) {
	query := `
		INSERT INTO foods (menu_id, name, price, image)
		VALUES ($1, $2, $3, $4)
		RETURNING id, menu_id, name, price, image, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query, f.MenuID, f.Name, f.Price, f.Image).Scan(
		&f.ID,
		&f.MenuID,
		&f.Name,
		&f.Price,
		&f.Image,
		&f.CreatedAt,
		&f.UpdatedAt,
	)
	if err != nil {
		return Food{}, fmt.Errorf("create food: %w", err)
	}

	return f, nil
}

func (r *repository) Update(ctx context.Context, id int, f Food) (Food, error) {
	query := `
		UPDATE foods
		SET name=$1, menu_id=$2, price=$3, image=COALESCE($4, image)
		WHERE id=$5
		RETURNING id, menu_id, name, price, image, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query,
		f.Name,
		f.MenuID,
		f.Price,
		f.Image,
		id,
	).Scan(
		&f.ID,
		&f.MenuID,
		&f.Name,
		&f.Price,
		&f.Image,
		&f.CreatedAt,
		&f.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Food{}, ErrFoodNotFound
		}

		return Food{}, fmt.Errorf("failed to update food: %w", err)
	}

	return f, nil

}
