package users

import (
	"net/http"

	"github.com/vald3mare/rudstock-users/internal/models"
)

func (h *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	identity, ok := currentIdentity(w, r)
	if !ok {
		return
	}

	body, ok := readRequest[changePasswordRequest](w, r)
	if !ok {
		return
	}
	err := h.svc.ChangePassword(r.Context(), identity.UserID, models.ChangePasswordInput{
		OldPassword: body.OldPassword,
		NewPassword: body.NewPassword,
		SessionID:   identity.SessionID,
	})
	if err != nil {
		writeServiceError(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
