package users

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vald3mare/rudstock-users/internal/models"
)

func (s *userService) ChangePassword(ctx context.Context, userID uuid.UUID, in models.ChangePasswordInput) error {
	if s.accounts == nil {
		return models.ErrAccountManagementUnavailable
	}
	if err := s.accounts.ChangePassword(ctx, userID, in.SessionID, in.OldPassword, in.NewPassword); err != nil {
		return fmt.Errorf("change account password: %w", err)
	}
	return nil
}
