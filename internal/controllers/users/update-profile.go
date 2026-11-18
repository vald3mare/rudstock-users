package users

import (
	"net/http"

	"github.com/vald3mare/rudstock-platform/httpx"
	"github.com/vald3mare/rudstock-users/internal/models"
)

func (h *UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	identity, ok := currentIdentity(w, r)
	if !ok {
		return
	}

	body, ok := readRequest[updateProfileRequest](w, r)
	if !ok {
		return
	}
	profile, err := h.svc.UpdateProfile(r.Context(), identity.UserID, models.UpdateProfileInput(*body))
	if err != nil {
		writeServiceError(w, r, err)
		return
	}
	httpx.WriteJSON(w, r, http.StatusOK, toProfileResponse(profile, identity))
}
