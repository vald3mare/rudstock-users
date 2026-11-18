package controllers

import (
	"net/http"

	"github.com/vald3mare/rudstock-platform/httpx"
	"github.com/vald3mare/rudstock-users/internal/controllers/users"
	"github.com/vald3mare/rudstock-users/internal/infrastructure/auth"
)

func NewRouter(h *users.UserHandler, authClient *auth.Client) http.Handler {
	mux := http.NewServeMux()
	me := http.NewServeMux()

	mux.HandleFunc("GET /health", health)

	me.HandleFunc("GET /me", h.GetProfile)
	me.HandleFunc("PATCH /me", h.UpdateProfile)
	me.HandleFunc("POST /me/password", h.ChangePassword)
	me.HandleFunc("DELETE /me", h.DeleteAccount)
	me.HandleFunc("GET /me/addresses", h.ListAddresses)
	me.HandleFunc("POST /me/addresses", h.CreateAddress)
	me.HandleFunc("PATCH /me/addresses/{id}", h.UpdateAddress)
	me.HandleFunc("DELETE /me/addresses/{id}", h.DeleteAddress)
	me.HandleFunc("GET /me/orders", h.ListOrders)
	me.HandleFunc("GET /me/orders/{id}", h.GetOrder)

	protected := authClient.Middleware(me)
	mux.Handle("/me", protected)
	mux.Handle("/me/", protected)

	return mux
}

func health(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(w, r, http.StatusOK, map[string]string{"status": "ok"})
}
