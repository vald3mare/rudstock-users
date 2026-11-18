package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vald3mare/rudstock-users/internal/models"
)

const deleteAddressQuery = `DELETE FROM user_addresses WHERE user_id = $1 AND id = $2`

func (r *userRepo) DeleteAddress(ctx context.Context, userID, addressID uuid.UUID) error {
	tag, err := r.pool.Exec(ctx, deleteAddressQuery, userID, addressID)
	if err != nil {
		return fmt.Errorf("delete address: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return models.ErrAddressNotFound
	}
	return nil
}
