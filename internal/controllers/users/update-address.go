package users

import (
	"net/http"

	"github.com/vald3mare/rudstock-platform/httpx"
	"github.com/vald3mare/rudstock-users/internal/models"
)

func (h *UserHandler) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	identity, ok := currentIdentity(w, r)
	if !ok {
		return
	}

	id, err := httpx.PathUUID(r, "id")
	if err != nil {
		httpx.WriteError(w, r, err)
		return
	}
	body, ok := readRequest[updateAddressRequest](w, r)
	if !ok {
		return
	}
	address, err := h.svc.UpdateAddress(r.Context(), identity.UserID, id, models.UpdateAddressInput(*body))
	if err != nil {
		writeServiceError(w, r, err)
		return
	}
	httpx.WriteJSON(w, r, http.StatusOK, addressResponse(address))
}
