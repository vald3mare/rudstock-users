package users

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *userService) DeleteAddress(ctx context.Context, userID, addressID uuid.UUID) error {
	if err := s.repo.DeleteAddress(ctx, userID, addressID); err != nil {
		return fmt.Errorf("delete address: %w", err)
	}

	return nil
}
