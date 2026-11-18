package users

import (
	"net/http"

	"github.com/vald3mare/rudstock-platform/httpx"
)

func (h *UserHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	identity, ok := currentIdentity(w, r)
	if !ok {
		return
	}

	id, err := httpx.PathUUID(r, "id")
	if err != nil {
		httpx.WriteError(w, r, err)
		return
	}
	order, err := h.svc.GetOrder(r.Context(), identity.UserID, id)
	if err != nil {
		writeServiceError(w, r, err)
		return
	}
	response := orderResponse{
		orderSummaryResponse: orderSummaryResponse(order.OrderSummary),
		Items:                make([]orderItemResponse, 0, len(order.Items)),
	}
	for _, item := range order.Items {
		response.Items = append(response.Items, orderItemResponse(item))
	}
	httpx.WriteJSON(w, r, http.StatusOK, response)
}
