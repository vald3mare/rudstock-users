package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

const deleteAccountDataQuery = `
	WITH deleted_profile AS (
		DELETE FROM user_profiles WHERE user_id = $1
	), deleted_addresses AS (
		DELETE FROM user_addresses WHERE user_id = $1
	)
	UPDATE orders SET user_id = NULL WHERE user_id = $1`

func (r *userRepo) DeleteAccountData(ctx context.Context, userID uuid.UUID) error {
	if _, err := r.pool.Exec(ctx, deleteAccountDataQuery, userID); err != nil {
		return fmt.Errorf("delete account data: %w", err)
	}
	return nil
}
