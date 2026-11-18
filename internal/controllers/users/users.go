package users

import (
	"net/http"

	"github.com/vald3mare/rudstock-platform/errs"
	"github.com/vald3mare/rudstock-platform/httpx"
	"github.com/vald3mare/rudstock-users/internal/infrastructure/auth"
	"github.com/vald3mare/rudstock-users/internal/services/users"
)

type UserHandler struct {
	svc users.UserService
}

func NewUserHandler(svc users.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func currentIdentity(w http.ResponseWriter, r *http.Request) (auth.Identity, bool) {
	identity, ok := auth.IdentityFrom(r.Context())
	if !ok {
		httpx.WriteError(w, r, errs.Internal("identity-missing", "Internal server error", nil))
	}
	return identity, ok
}

func readRequest[T any](w http.ResponseWriter, r *http.Request) (*T, bool) {
	var input *T
	if err := httpx.DecodeJSON(w, r, &input); err != nil {
		httpx.WriteError(w, r, err)
		return nil, false
	}
	if input == nil {
		httpx.WriteError(w, r, errs.InvalidInput("invalid-json-body", "Request body must be a JSON object", nil))
		return nil, false
	}
	return input, true
}
