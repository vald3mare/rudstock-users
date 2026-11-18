package users

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vald3mare/rudstock-users/internal/models"
)

func (s *userService) GetOrder(ctx context.Context, userID, orderID uuid.UUID) (models.Order, error) {
	order, err := s.orders.GetOrder(ctx, userID, orderID)
	if err != nil {
		return models.Order{}, fmt.Errorf("get order: %w", err)
	}

	return order, nil
}
