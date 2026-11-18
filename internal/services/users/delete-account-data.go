package users

import (
	"context"

	"github.com/google/uuid"
)

func (s *userService) DeleteAccountData(ctx context.Context, userID uuid.UUID) error {
	return s.repo.DeleteAccountData(ctx, userID)
}
