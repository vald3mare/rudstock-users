package users

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vald3mare/rudstock-users/internal/models"
)

func (s *userService) GetProfile(ctx context.Context, userID uuid.UUID) (models.Profile, error) {
	profile, err := s.repo.GetProfile(ctx, userID)
	if err != nil {
		return models.Profile{}, fmt.Errorf("get profile: %w", err)
	}

	return profile, nil
}
