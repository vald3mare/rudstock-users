package models

import "errors"

var (
	ErrEmptyPatch                   = errors.New("nothing to update")
	ErrInvalidAddress               = errors.New("city, street and house must not be empty")
	ErrUserNotFound                 = errors.New("user not found")
	ErrAddressNotFound              = errors.New("address not found")
	ErrOrderNotFound                = errors.New("order not found")
	ErrInvalidPagination            = errors.New("invalid pagination")
	ErrInvalidPassword              = errors.New("invalid current password")
	ErrWeakPassword                 = errors.New("password must be at least 8 bytes and contain a letter, a digit and a special character")
	ErrInvalidSession               = errors.New("current session is required")
	ErrAccountManagementUnavailable = errors.New("account management is unavailable")
)
