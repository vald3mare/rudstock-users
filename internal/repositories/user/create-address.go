package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vald3mare/rudstock-users/internal/models"
)

const createAddressQuery = `
		INSERT INTO user_addresses (user_id, city, street, house, flat, comment)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

func (r *userRepo) CreateAddress(ctx context.Context, userID uuid.UUID, address models.Address) (uuid.UUID, error) {
	var id uuid.UUID
	err := r.pool.QueryRow(ctx, createAddressQuery, userID, address.City, address.Street, address.House, address.Flat, address.Comment).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("insert address: %w", err)
	}
	return id, nil
}
