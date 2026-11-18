package users

import (
	"context"
	"fmt"
	"math"

	"github.com/google/uuid"
	"github.com/vald3mare/rudstock-users/internal/models"
)

const (
	defaultOrdersLimit = 20
	maxOrdersLimit     = 100
)

func (s *userService) ListOrders(ctx context.Context, userID uuid.UUID, in models.ListOrdersInput) ([]models.OrderSummary, int64, error) {
	if in.Page < 0 || in.Limit < 0 {
		return nil, 0, models.ErrInvalidPagination
	}

	page := max(in.Page, 1)
	limit := in.Limit
	if limit == 0 {
		limit = defaultOrdersLimit
	}
	limit = min(limit, maxOrdersLimit)
	if page-1 > math.MaxInt/limit {
		return nil, 0, models.ErrInvalidPagination
	}

	filter := models.OrderFilter{
		Limit:  limit,
		Offset: (page - 1) * limit,
	}
	orders, total, err := s.orders.ListOrders(ctx, userID, filter)
	if err != nil {
		return nil, 0, fmt.Errorf("list orders: %w", err)
	}

	return orders, total, nil
}
