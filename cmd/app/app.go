package app

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vald3mare/rudstock-users/internal/controllers"
	userctrl "github.com/vald3mare/rudstock-users/internal/controllers/users"
	"github.com/vald3mare/rudstock-users/internal/infrastructure/auth"
	"github.com/vald3mare/rudstock-users/internal/repositories/order"
	"github.com/vald3mare/rudstock-users/internal/repositories/user"
	"github.com/vald3mare/rudstock-users/internal/services/users"
)

type App struct {
	Router http.Handler
}

func NewApp(pool *pgxpool.Pool, authURL string, accounts users.AccountManager) (*App, error) {
	authClient, err := auth.NewClient(authURL)
	if err != nil {
		return nil, err
	}

	service := users.NewUserService(user.NewUserRepo(pool), order.NewOrderRepo(pool), accounts)
	return &App{
		Router: controllers.NewRouter(userctrl.NewUserHandler(service), authClient),
	}, nil
}
