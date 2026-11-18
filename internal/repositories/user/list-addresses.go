package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/vald3mare/rudstock-users/internal/models"
)

const listAddressesQuery = `
		SELECT ` + addressColumns + `
		FROM user_addresses
		WHERE user_id = $1
		ORDER BY created_at DESC, id DESC`

func (r *userRepo) ListAddresses(ctx context.Context, userID uuid.UUID) ([]models.Address, error) {
	rows, err := r.pool.Query(ctx, listAddressesQuery, userID)
	if err != nil {
		return nil, fmt.Errorf("select addresses: %w", err)
	}
	addresses, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.Address, error) {
		return scanAddress(row)
	})
	if err != nil {
		return nil, fmt.Errorf("scan addresses: %w", err)
	}
	return addresses, nil
}
