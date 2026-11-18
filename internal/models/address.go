package models

import "github.com/google/uuid"

type Address struct {
	ID      uuid.UUID
	City    string
	Street  string
	House   string
	Flat    string
	Comment string
}

// AddressPatch содержит изменения адреса. nil оставляет поле без изменений.
// Пустая строка очищает необязательное поле.
type AddressPatch struct {
	City    *string
	Street  *string
	House   *string
	Flat    *string
	Comment *string
}

// CreateAddressInput содержит адрес доставки. Flat и Comment необязательны
type CreateAddressInput struct {
	City    string
	Street  string
	House   string
	Flat    string
	Comment string
}

// UpdateAddressInput содержит поля адреса для частичного обновления
// nil оставляет поле без изменений. Пустая строка очищает необязательное поле
type UpdateAddressInput struct {
	City    *string
	Street  *string
	House   *string
	Flat    *string
	Comment *string
}
