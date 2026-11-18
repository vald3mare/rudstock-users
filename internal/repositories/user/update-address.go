package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/vald3mare/rudstock-users/internal/models"
)

const updateAddressQuery = `
		UPDATE user_addresses
		SET city    = COALESCE($3, city),
		    street  = COALESCE($4, street),
		    house   = COALESCE($5, house),
		    flat    = COALESCE($6, flat),
		    comment = COALESCE($7, comment)
		WHERE user_id = $1 AND id = $2
		RETURNING ` + addressColumns

func (r *userRepo) UpdateAddress(ctx context.Context, userID, addressID uuid.UUID, patch models.AddressPatch) (models.Address, error) {
	address, err := scanAddress(r.pool.QueryRow(ctx, updateAddressQuery, userID, addressID, patch.City, patch.Street, patch.House, patch.Flat, patch.Comment))
	if errors.Is(err, pgx.ErrNoRows) {
		return models.Address{}, models.ErrAddressNotFound
	}
	if err != nil {
		return models.Address{}, fmt.Errorf("update address: %w", err)
	}
	return address, nil
}
