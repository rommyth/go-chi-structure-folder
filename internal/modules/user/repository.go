package user

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
	ErrUserNotFound = errors.New("user not found")
)

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) List(ctx context.Context, page, limit int, search string) ([]User, int, error) {
	offset := (page - 1) * limit

	var conditions []string
	var args []interface{}
	argIdx := 1

	if search != "" {
		conditions = append(conditions, fmt.Sprintf("(first_name ILIKE $%d OR last_name ILIKE $%d OR email ILIKE $%d)", argIdx, argIdx+1, argIdx+2))
		args = append(args, "%"+search+"%", "%"+search+"%", "%"+search+"%")
		argIdx += 3
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")

	}

	countQuery := fmt.Sprintf(`
			SELECT COUNT(*)
			FROM users
			%s
		`, whereClause)

	var total int
	err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return []User{}, 0, fmt.Errorf("failed count user: %w", err)
	}

	query := fmt.Sprintf(`
			SELECT id, first_name, last_name, email, avatar, phone, created_at, updated_at
			FROM users
			%s
			LIMIT $%d OFFSET $%d
		`, whereClause, argIdx, argIdx+1)

	args = append(args, limit, offset)

	var users []User
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return []User{}, 0, fmt.Errorf("failed get users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User

		if err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Avatar,
			&user.Phone,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return []User{}, 0, fmt.Errorf("failed scan user: %w", err)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return []User{}, 0, fmt.Errorf("failed to iterate users: %w", err)
	}

	return users, total, nil

}

func (r *repository) GetByID(ctx context.Context, id int) (User, error) {
	query := `
		SELECT id, first_name, last_name, email, avatar, phone, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Avatar,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return User{}, ErrUserNotFound
		}

		return User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (r *repository) GetByEmail(ctx context.Context, email string) (User, error) {
	query := `
		SELECT id, first_name, last_name, email, password, avatar, phone, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user User
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Avatar,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return User{}, ErrUserNotFound
		}

		return User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (r *repository) Create(ctx context.Context, user User) (User, error) {
	query := `
		INSERT INTO users (first_name, last_name, email, password, phone, avatar)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, first_name, last_name, email, phone, avatar, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Phone,
		user.Avatar,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.Avatar,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (r *repository) Update(ctx context.Context, id int, user User) (User, error) {
	query := `
		UPDATE users
		SET first_name=$1, last_name=$2, email=$3, phone=$4, avatar=$5
		WHERE id = $6
		RETURNING id, first_name, last_name, email, phone, avatar, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Avatar,
		id,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.Avatar,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return User{}, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}
