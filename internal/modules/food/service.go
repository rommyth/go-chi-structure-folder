package food

import (
	"context"
	"errors"
	"fmt"
	"restaurant-management/internal/modules/menu"
)

type Repository interface {
	GetByID(ctx context.Context, id int) (Food, error)
	List(ctx context.Context, page, limit int, search string) ([]Food, int, error)
	Create(ctx context.Context, f Food) (Food, error)
	Update(ctx context.Context, id int, f Food) (Food, error)
}

type Service struct {
	repo     Repository
	menuRepo menu.Repository
}

func NewService(repo Repository, menuRepo menu.Repository) *Service {
	return &Service{repo: repo, menuRepo: menuRepo}
}

func (s *Service) GetByID(ctx context.Context, id int) (Food, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) Create(ctx context.Context, req CreateFoodRequest) (Food, error) {
	_, err := s.menuRepo.GetByID(ctx, int(req.MenuID))
	if err != nil {
		if errors.Is(err, menu.ErrMenuNotFound) {
			return Food{}, menu.ErrMenuNotFound
		}
		return Food{}, fmt.Errorf("validate menu id: %w", err)
	}

	food := Food{
		MenuID: req.MenuID,
		Name:   req.Name,
		Price:  req.Price,
		Image:  req.Image,
	}

	res, err := s.repo.Create(ctx, food)
	if err != nil {
		return Food{}, fmt.Errorf("failed create food: %w", err)
	}

	return res, nil
}
