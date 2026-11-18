package users

import (
	"net/http"

	"github.com/vald3mare/rudstock-platform/httpx"
)

func (h *UserHandler) ListAddresses(w http.ResponseWriter, r *http.Request) {
	identity, ok := currentIdentity(w, r)
	if !ok {
		return
	}

	addresses, err := h.svc.ListAddresses(r.Context(), identity.UserID)
	if err != nil {
		writeServiceError(w, r, err)
		return
	}
	response := make([]addressResponse, 0, len(addresses))
	for _, address := range addresses {
		response = append(response, addressResponse(address))
	}
	httpx.WriteJSON(w, r, http.StatusOK, response)
}
