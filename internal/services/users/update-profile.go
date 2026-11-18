package users

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vald3mare/rudstock-users/internal/models"
)

func (s *userService) UpdateProfile(ctx context.Context, userID uuid.UUID, in models.UpdateProfileInput) (models.Profile, error) {
	if in.Name == nil && in.Phone == nil {
		return models.Profile{}, models.ErrEmptyPatch
	}

	patch := models.ProfilePatch{
		Name:  in.Name,
		Phone: in.Phone,
	}
	profile, err := s.repo.UpdateProfile(ctx, userID, patch)
	if err != nil {
		return models.Profile{}, fmt.Errorf("update profile: %w", err)
	}

	return profile, nil
}
