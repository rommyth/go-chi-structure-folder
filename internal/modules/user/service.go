package user

import (
	"context"
	"errors"
)

type Repository interface {
	List(ctx context.Context, page, limit int, search string) ([]User, int, error)
	GetByID(ctx context.Context, id int) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	Create(ctx context.Context, user User) (User, error)
	Update(ctx context.Context, id int, user User) (User, error)
}

type Service struct {
	repo Repository
}

var (
	ErrEmailAlreadyExist = errors.New("email already exist")
)

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUsers(ctx context.Context, page, limit int, search string) ([]User, int, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 || limit > 100 {
		limit = 10
	}

	users, total, err := s.repo.List(ctx, page, limit, search)
	if err != nil {
		return []User{}, 0, err
	}

	return users, total, nil
}

func (s *Service) GetUserByID(ctx context.Context, id int) (User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) Update(ctx context.Context, id int, req UpdateUserRequest) (User, error) {
	_, err := s.repo.GetByEmail(ctx, req.Email)
	if err == nil {
		return User{}, ErrEmailAlreadyExist
	}

	if !errors.Is(err, ErrUserNotFound) {
		return User{}, err
	}

	user := User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Avatar:    req.Avatar,
	}

	return s.repo.Update(ctx, id, user)
}
