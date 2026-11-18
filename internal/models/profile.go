package models

import "github.com/google/uuid"

type Profile struct {
	ID    uuid.UUID
	Name  string
	Phone string
}

// ProfilePatch содержит изменения профиля. nil оставляет поле без изменений.
type ProfilePatch struct {
	Name  *string
	Phone *string
}

// UpdateProfileInput содержит поля профиля, доступные для изменения
// nil оставляет поле без изменений.
type UpdateProfileInput struct {
	Name  *string
	Phone *string
}

// ChangePasswordInput содержит текущий и новый пароли
type ChangePasswordInput struct {
	OldPassword string
	NewPassword string
	SessionID   uuid.UUID // ID текущей сессии из проверенного токена
}
