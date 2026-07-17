package warehouse

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AdjustStock(ctx context.Context, itemID string, newQty int) error {
	// Здесь можно добавить проверку: например, нельзя ставить отрицательное количество
	if newQty < 0 {
		return ErrInvalidQuantity
	}
	return s.repo.UpdateStock(ctx, itemID, newQty)
}
