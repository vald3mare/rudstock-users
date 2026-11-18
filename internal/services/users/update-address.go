package users

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/vald3mare/rudstock-users/internal/models"
)

func (s *userService) UpdateAddress(ctx context.Context, userID, addressID uuid.UUID, in models.UpdateAddressInput) (models.Address, error) {
	if in.City == nil && in.Street == nil && in.House == nil && in.Flat == nil && in.Comment == nil {
		return models.Address{}, models.ErrEmptyPatch
	}

	patch := models.AddressPatch{
		City:    trimAddressField(in.City),
		Street:  trimAddressField(in.Street),
		House:   trimAddressField(in.House),
		Flat:    in.Flat,
		Comment: in.Comment,
	}
	for _, field := range []*string{patch.City, patch.Street, patch.House} {
		if field != nil && *field == "" {
			return models.Address{}, models.ErrInvalidAddress
		}
	}

	address, err := s.repo.UpdateAddress(ctx, userID, addressID, patch)
	if err != nil {
		return models.Address{}, fmt.Errorf("update address: %w", err)
	}

	return address, nil
}

func trimAddressField(field *string) *string {
	if field == nil {
		return nil
	}
	value := strings.TrimSpace(*field)
	return &value
}
