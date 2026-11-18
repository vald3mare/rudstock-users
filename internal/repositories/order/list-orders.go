package order

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vald3mare/rudstock-users/internal/models"
)

const listOrdersQuery = `
		WITH page AS (
			SELECT id, number, status, total, created_at
			FROM orders
			WHERE user_id = $1
			ORDER BY created_at DESC, id DESC
			LIMIT $2 OFFSET $3
		), tally AS (
			SELECT count(*) AS total FROM orders WHERE user_id = $1
		)
		SELECT page.id, COALESCE(page.number, ''), COALESCE(page.status, ''),
		       COALESCE(page.total, 0), tally.total
		FROM tally LEFT JOIN page ON true
		ORDER BY page.created_at DESC, page.id DESC`

func (r *orderRepo) ListOrders(ctx context.Context, userID uuid.UUID, filter models.OrderFilter) ([]models.OrderSummary, int64, error) {
	// LEFT JOIN сохраняет total даже для страницы за пределами списка.
	// Страница и total читаются из одного снимка БД.
	rows, err := r.pool.Query(ctx, listOrdersQuery, userID, filter.Limit, filter.Offset)
	if err != nil {
		return nil, 0, fmt.Errorf("select orders: %w", err)
	}
	defer rows.Close()

	orders := make([]models.OrderSummary, 0)
	var total int64
	for rows.Next() {
		var id *uuid.UUID
		var order models.OrderSummary
		if err := rows.Scan(&id, &order.Number, &order.Status, &order.Total, &total); err != nil {
			return nil, 0, fmt.Errorf("scan orders: %w", err)
		}
		if id != nil {
			order.ID = *id
			orders = append(orders, order)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("read orders: %w", err)
	}
	return orders, total, nil
}
