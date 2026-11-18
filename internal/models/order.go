package models

import "github.com/google/uuid"

// OrderSummary данные заказа для списка в личном кабинете
type OrderSummary struct {
	ID     uuid.UUID
	Number string
	Status string
	Total  int64 // в копейках
}

// Order данные заказа для просмотра с его позициями
type Order struct {
	OrderSummary
	Items []OrderItem
}

// OrderItem товар и его количество в заказе
type OrderItem struct {
	CardID   uuid.UUID
	Quantity int
}

// OrderFilter задаёт размер страницы и смещение от начала списка заказов.
type OrderFilter struct {
	Limit  int64
	Offset int64
}

// ListOrdersInput задаёт страницу истории заказов
// Нумерация страниц начинается с 1. Ноль выбирает значение по умолчанию
type ListOrdersInput struct {
	Page  int64
	Limit int64
}
