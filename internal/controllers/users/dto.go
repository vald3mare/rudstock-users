package users

import (
	"github.com/google/uuid"
	"github.com/vald3mare/rudstock-users/internal/infrastructure/auth"
	"github.com/vald3mare/rudstock-users/internal/models"
)

type profileResponse struct {
	ID            uuid.UUID `json:"id"`
	Email         *string   `json:"email"`
	EmailVerified *bool     `json:"email_verified"`
	Name          string    `json:"name"`
	Phone         string    `json:"phone"`
	Role          string    `json:"role"`
}

func toProfileResponse(profile models.Profile, identity auth.Identity) profileResponse {
	return profileResponse{
		ID:            profile.ID,
		Email:         identity.Email,
		EmailVerified: identity.EmailVerified,
		Name:          profile.Name,
		Phone:         profile.Phone,
		Role:          identity.Role,
	}
}

type updateProfileRequest struct {
	Name  *string `json:"name"`
	Phone *string `json:"phone"`
}

type changePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type deleteAccountRequest struct {
	Password string `json:"password"`
}

type addressResponse struct {
	ID      uuid.UUID `json:"id"`
	City    string    `json:"city"`
	Street  string    `json:"street"`
	House   string    `json:"house"`
	Flat    string    `json:"flat"`
	Comment string    `json:"comment"`
}

type createAddressRequest struct {
	City    string `json:"city"`
	Street  string `json:"street"`
	House   string `json:"house"`
	Flat    string `json:"flat"`
	Comment string `json:"comment"`
}

type createAddressResponse struct {
	AddressID uuid.UUID `json:"address_id"`
}

type updateAddressRequest struct {
	City    *string `json:"city"`
	Street  *string `json:"street"`
	House   *string `json:"house"`
	Flat    *string `json:"flat"`
	Comment *string `json:"comment"`
}

type orderSummaryResponse struct {
	ID     uuid.UUID `json:"id"`
	Number string    `json:"number"`
	Status string    `json:"status"`
	Total  int64     `json:"total"`
}

type listOrdersResponse struct {
	Items []orderSummaryResponse `json:"items"`
	Total int64                  `json:"total"`
}

type orderResponse struct {
	orderSummaryResponse
	Items []orderItemResponse `json:"items"`
}

type orderItemResponse struct {
	CardID   uuid.UUID `json:"card_id"`
	Quantity int       `json:"quantity"`
}
