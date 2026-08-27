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

func (s *Service) GetList(ctx context.Context, page, limit int, search string) ([]Food, int, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 || limit > 100 {
		limit = 10
	}

	foods, total, err := s.repo.List(ctx, page, limit, search)
	if err != nil {
		return []Food{}, 0, fmt.Errorf("failed get list foods : %w", err)
	}

	return foods, total, nil
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

func (s *Service) Update(ctx context.Context, id int, f UpdateFoodRequest) (Food, error) {
	food, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return Food{}, err
	}

	if _, err = s.menuRepo.GetByID(ctx, int(f.MenuID)); err != nil {
		return Food{}, err
	}

	food.Name = f.Name
	food.Image = f.Image
	food.MenuID = f.MenuID
	food.Price = f.Price

	if _, err := s.repo.Update(ctx, id, food); err != nil {
		return Food{}, err
	}

	return food, nil
}
