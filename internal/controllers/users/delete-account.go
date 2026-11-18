package users

import (
	"net/http"
)

func (h *UserHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	identity, ok := currentIdentity(w, r)
	if !ok {
		return
	}

	body, ok := readRequest[deleteAccountRequest](w, r)
	if !ok {
		return
	}
	err := h.svc.DeleteAccount(r.Context(), identity.UserID, body.Password)
	if err != nil {
		writeServiceError(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
