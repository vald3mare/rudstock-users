package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vald3mare/rudstock-users/internal/models"
)

const updateProfileQuery = `
	INSERT INTO user_profiles (user_id, name, phone)
	VALUES ($1, COALESCE($2::text, ''), COALESCE($3::text, ''))
	ON CONFLICT (user_id) DO UPDATE
	SET name = COALESCE($2, user_profiles.name),
	    phone = COALESCE($3, user_profiles.phone)
	RETURNING user_id, name, phone`

func (r *userRepo) UpdateProfile(ctx context.Context, userID uuid.UUID, patch models.ProfilePatch) (models.Profile, error) {
	profile, err := scanProfile(r.pool.QueryRow(ctx, updateProfileQuery, userID, patch.Name, patch.Phone))
	if err != nil {
		return models.Profile{}, fmt.Errorf("update profile: %w", err)
	}
	return profile, nil
}
