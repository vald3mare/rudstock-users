package users

import (
	"math"
	"net/http"

	"github.com/vald3mare/rudstock-platform/httpx"
	"github.com/vald3mare/rudstock-users/internal/models"
)

func (h *UserHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	identity, ok := currentIdentity(w, r)
	if !ok {
		return
	}

	page, err := httpx.QueryInt64(r, "page")
	if err != nil {
		httpx.WriteError(w, r, err)
		return
	}
	limit, err := httpx.QueryInt64(r, "limit")
	if err != nil {
		httpx.WriteError(w, r, err)
		return
	}
	if page < 0 || limit < 0 || page > math.MaxInt || limit > math.MaxInt {
		writeServiceError(w, r, models.ErrInvalidPagination)
		return
	}
	orders, total, err := h.svc.ListOrders(r.Context(), identity.UserID, models.ListOrdersInput{
		Page:  int(page),
		Limit: int(limit),
	})
	if err != nil {
		writeServiceError(w, r, err)
		return
	}
	response := listOrdersResponse{Items: make([]orderSummaryResponse, 0, len(orders)), Total: total}
	for _, order := range orders {
		response.Items = append(response.Items, orderSummaryResponse(order))
	}
	httpx.WriteJSON(w, r, http.StatusOK, response)
}
