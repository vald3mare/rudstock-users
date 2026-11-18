package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/vald3mare/rudstock-users/internal/models"
)

const getProfileQuery = `SELECT user_id, name, phone FROM user_profiles WHERE user_id = $1`

func (r *userRepo) GetProfile(ctx context.Context, userID uuid.UUID) (models.Profile, error) {
	profile, err := scanProfile(r.pool.QueryRow(ctx, getProfileQuery, userID))
	if errors.Is(err, pgx.ErrNoRows) {
		return models.Profile{ID: userID}, nil
	}
	if err != nil {
		return models.Profile{}, fmt.Errorf("select profile: %w", err)
	}
	return profile, nil
}
