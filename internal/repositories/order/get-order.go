package order

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vald3mare/rudstock-users/internal/models"
)

const gerOrderQuery = `
		SELECT o.id, o.number, o.status, o.total, i.card_id, COALESCE(i.quantity, 0)
		FROM orders o
		LEFT JOIN order_items i ON i.order_id = o.id
		WHERE o.user_id = $1 AND o.id = $2
		ORDER BY i.position`

func (r *orderRepo) GetOrder(ctx context.Context, userID, orderID uuid.UUID) (models.Order, error) {
	rows, err := r.pool.Query(ctx, gerOrderQuery, userID, orderID)
	if err != nil {
		return models.Order{}, fmt.Errorf("select order: %w", err)
	}
	defer rows.Close()

	order := models.Order{Items: make([]models.OrderItem, 0)}
	found := false
	for rows.Next() {
		var cardID *uuid.UUID
		var quantity int
		if err := rows.Scan(&order.ID, &order.Number, &order.Status, &order.Total, &cardID, &quantity); err != nil {
			return models.Order{}, fmt.Errorf("scan order: %w", err)
		}
		found = true
		if cardID != nil {
			order.Items = append(order.Items, models.OrderItem{CardID: *cardID, Quantity: quantity})
		}
	}
	if err := rows.Err(); err != nil {
		return models.Order{}, fmt.Errorf("read order: %w", err)
	}
	if !found {
		return models.Order{}, models.ErrOrderNotFound
	}
	return order, nil
}
