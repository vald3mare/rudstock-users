package users

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/vald3mare/rudstock-users/internal/models"
)

func (s *userService) CreateAddress(ctx context.Context, userID uuid.UUID, in models.CreateAddressInput) (uuid.UUID, error) {
	address := models.Address{
		City:    strings.TrimSpace(in.City),
		Street:  strings.TrimSpace(in.Street),
		House:   strings.TrimSpace(in.House),
		Flat:    in.Flat,
		Comment: in.Comment,
	}
	if address.City == "" || address.Street == "" || address.House == "" {
		return uuid.Nil, models.ErrInvalidAddress
	}

	id, err := s.repo.CreateAddress(ctx, userID, address)
	if err != nil {
		return uuid.Nil, fmt.Errorf("create address: %w", err)
	}

	return id, nil
}
