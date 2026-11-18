package users

import (
	"net/http"

	"github.com/vald3mare/rudstock-platform/httpx"
)

func (h *UserHandler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	identity, ok := currentIdentity(w, r)
	if !ok {
		return
	}

	id, err := httpx.PathUUID(r, "id")
	if err != nil {
		httpx.WriteError(w, r, err)
		return
	}
	err = h.svc.DeleteAddress(r.Context(), identity.UserID, id)
	if err != nil {
		writeServiceError(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
