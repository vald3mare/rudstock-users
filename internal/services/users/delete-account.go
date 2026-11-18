package users

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vald3mare/rudstock-users/internal/models"
)

func (s *userService) DeleteAccount(ctx context.Context, userID uuid.UUID, password string) error {
	if s.accounts == nil {
		return models.ErrAccountManagementUnavailable
	}
	if err := s.accounts.DeleteAccount(ctx, userID, password); err != nil {
		return fmt.Errorf("delete auth account: %w", err)
	}
	return s.DeleteAccountData(ctx, userID)
}
