package users

import (
	"context"

	"github.com/google/uuid"
	"github.com/vald3mare/rudstock-users/internal/models"
)

// UserService описывает операции личного кабинета
// userID берётся из проверенного токена пользователя
type UserService interface {
	// GetProfile возвращает профиль текущего пользователя
	GetProfile(ctx context.Context, userID uuid.UUID) (models.Profile, error)

	// UpdateProfile обновляет переданные поля и возвращает профиль целиком
	UpdateProfile(ctx context.Context, userID uuid.UUID, in models.UpdateProfileInput) (models.Profile, error)

	// ChangePassword проверяет текущий пароль, меняет его и отзывает остальные сессии
	ChangePassword(ctx context.Context, userID uuid.UUID, in models.ChangePasswordInput) error

	// DeleteAccount подтверждает удаление паролем, обезличивает заказы и отзывает все сессии
	DeleteAccount(ctx context.Context, userID uuid.UUID, password string) error
	DeleteAccountData(ctx context.Context, userID uuid.UUID) error

	// ListAddresses возвращает сохранённые адреса пользователя
	ListAddresses(ctx context.Context, userID uuid.UUID) ([]models.Address, error)

	// CreateAddress сохраняет адрес и возвращает его ID
	CreateAddress(ctx context.Context, userID uuid.UUID, in models.CreateAddressInput) (uuid.UUID, error)

	// UpdateAddress обновляет переданные поля адреса и возвращает адрес целиком
	// Для чужого адреса возвращается та же ошибка, что и для отсутствующего
	UpdateAddress(ctx context.Context, userID, addressID uuid.UUID, in models.UpdateAddressInput) (models.Address, error)

	// DeleteAddress удаляет адрес пользователя
	// Для чужого адреса возвращается та же ошибка, что и для отсутствующего
	DeleteAddress(ctx context.Context, userID, addressID uuid.UUID) error

	// ListOrders возвращает страницу заказов пользователя и их общее число
	ListOrders(ctx context.Context, userID uuid.UUID, in models.ListOrdersInput) ([]models.OrderSummary, int64, error)

	// GetOrder возвращает заказ пользователя с его позициями
	// Для чужого заказа возвращается та же ошибка, что и для отсутствующего
	GetOrder(ctx context.Context, userID, orderID uuid.UUID) (models.Order, error)
}

// UserRepo описывает операции, которые нужны сервису от хранилища
type UserRepo interface {
	GetProfile(ctx context.Context, userID uuid.UUID) (models.Profile, error)
	UpdateProfile(ctx context.Context, userID uuid.UUID, patch models.ProfilePatch) (models.Profile, error)
	ListAddresses(ctx context.Context, userID uuid.UUID) ([]models.Address, error)
	CreateAddress(ctx context.Context, userID uuid.UUID, address models.Address) (uuid.UUID, error)
	UpdateAddress(ctx context.Context, userID, addressID uuid.UUID, patch models.AddressPatch) (models.Address, error)
	DeleteAddress(ctx context.Context, userID, addressID uuid.UUID) error
	DeleteAccountData(ctx context.Context, userID uuid.UUID) error
}

// OrderRepo описывает операции с заказами, которые нужны личному кабинету
type OrderRepo interface {
	ListOrders(ctx context.Context, userID uuid.UUID, filter models.OrderFilter) ([]models.OrderSummary, int64, error)
	GetOrder(ctx context.Context, userID, orderID uuid.UUID) (models.Order, error)
}

type AccountManager interface {
	ChangePassword(ctx context.Context, userID, sessionID uuid.UUID, oldPassword, newPassword string) error
	DeleteAccount(ctx context.Context, userID uuid.UUID, password string) error
}

type userService struct {
	repo     UserRepo
	orders   OrderRepo
	accounts AccountManager
}

func NewUserService(repo UserRepo, orders OrderRepo, accounts AccountManager) UserService {
	return &userService{repo: repo, orders: orders, accounts: accounts}
}
