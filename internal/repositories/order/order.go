package order

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vald3mare/rudstock-users/internal/models"
)

// OrderRepo читает историю заказов пользователя
type OrderRepo interface {
	// ListOrders возвращает страницу заказов и общее число без учёта пагинации
	ListOrders(ctx context.Context, userID uuid.UUID, filter models.OrderFilter) ([]models.OrderSummary, int64, error)
	// GetOrder возвращает заказ пользователя с его позициями или models.ErrOrderNotFound
	GetOrder(ctx context.Context, userID, orderID uuid.UUID) (models.Order, error)
}

type orderRepo struct {
	pool *pgxpool.Pool
}

func NewOrderRepo(pool *pgxpool.Pool) OrderRepo {
	return &orderRepo{pool: pool}
}
