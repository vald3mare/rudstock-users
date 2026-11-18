package users

import (
	"net/http"

	"github.com/vald3mare/rudstock-platform/httpx"
)

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	identity, ok := currentIdentity(w, r)
	if !ok {
		return
	}

	profile, err := h.svc.GetProfile(r.Context(), identity.UserID)
	if err != nil {
		writeServiceError(w, r, err)
		return
	}
	httpx.WriteJSON(w, r, http.StatusOK, toProfileResponse(profile, identity))
}
