package users

import (
	"errors"
	"net/http"

	"github.com/vald3mare/rudstock-platform/errs"
	"github.com/vald3mare/rudstock-platform/httpx"
	"github.com/vald3mare/rudstock-users/internal/models"
)

func writeServiceError(w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, models.ErrAccountManagementUnavailable) {
		httpx.WriteJSON(w, r, http.StatusServiceUnavailable, httpx.ErrorResponse{
			Slug: "account-management-unavailable",
			Msg:  "Account management is unavailable",
		})
		return
	}
	httpx.WriteError(w, r, mapServiceError(err))
}

func mapServiceError(err error) error {
	switch {
	case errors.Is(err, models.ErrEmptyPatch):
		return errs.InvalidInput("empty-patch", "At least one field must be provided", err)
	case errors.Is(err, models.ErrInvalidAddress):
		return errs.InvalidInput("invalid-address", "City, street and house must not be empty", err)
	case errors.Is(err, models.ErrInvalidPagination):
		return errs.InvalidInput("invalid-pagination", "Invalid page or limit", err)
	case errors.Is(err, models.ErrWeakPassword):
		return errs.InvalidInput("weak-password", models.ErrWeakPassword.Error(), err)
	case errors.Is(err, models.ErrInvalidPassword):
		return errs.Unauthorized("invalid-password", "Invalid current password", err)
	case errors.Is(err, models.ErrInvalidSession):
		return errs.Unauthorized("invalid-session", "Current session is required", err)
	case errors.Is(err, models.ErrUserNotFound):
		return errs.NotFound("user-not-found", "User not found", err)
	case errors.Is(err, models.ErrAddressNotFound):
		return errs.NotFound("address-not-found", "Address not found", err)
	case errors.Is(err, models.ErrOrderNotFound):
		return errs.NotFound("order-not-found", "Order not found", err)
	default:
		return err
	}
}
