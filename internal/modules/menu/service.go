package menu

import "context"

type Repository interface {
	GetByID(ctx context.Context, id int) (Menu, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}
