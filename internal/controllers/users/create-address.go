package users

import (
	"net/http"

	"github.com/vald3mare/rudstock-platform/httpx"
	"github.com/vald3mare/rudstock-users/internal/models"
)

func (h *UserHandler) CreateAddress(w http.ResponseWriter, r *http.Request) {
	identity, ok := currentIdentity(w, r)
	if !ok {
		return
	}

	body, ok := readRequest[createAddressRequest](w, r)
	if !ok {
		return
	}
	id, err := h.svc.CreateAddress(r.Context(), identity.UserID, models.CreateAddressInput(*body))
	if err != nil {
		writeServiceError(w, r, err)
		return
	}
	httpx.WriteJSON(w, r, http.StatusCreated, createAddressResponse{AddressID: id})
}
