package users

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vald3mare/rudstock-users/internal/models"
)

func (s *userService) ListAddresses(ctx context.Context, userID uuid.UUID) ([]models.Address, error) {
	addresses, err := s.repo.ListAddresses(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("list addresses: %w", err)
	}

	return addresses, nil
}
