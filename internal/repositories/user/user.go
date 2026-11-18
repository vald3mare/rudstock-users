package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vald3mare/rudstock-users/internal/models"
)

// UserRepo описывает хранение профиля и адресов пользователя
type UserRepo interface {
	GetProfile(ctx context.Context, userID uuid.UUID) (models.Profile, error)
	UpdateProfile(ctx context.Context, userID uuid.UUID, patch models.ProfilePatch) (models.Profile, error)

	ListAddresses(ctx context.Context, userID uuid.UUID) ([]models.Address, error)
	CreateAddress(ctx context.Context, userID uuid.UUID, address models.Address) (uuid.UUID, error)
	UpdateAddress(ctx context.Context, userID, addressID uuid.UUID, patch models.AddressPatch) (models.Address, error)
	DeleteAddress(ctx context.Context, userID, addressID uuid.UUID) error

	DeleteAccountData(ctx context.Context, userID uuid.UUID) error
}

type userRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) UserRepo {
	return &userRepo{pool: pool}
}

const addressColumns = `id, city, street, house, flat, comment`

func scanAddress(row pgx.Row) (models.Address, error) {
	var address models.Address
	err := row.Scan(&address.ID, &address.City, &address.Street, &address.House, &address.Flat, &address.Comment)
	return address, err
}

func scanProfile(row pgx.Row) (models.Profile, error) {
	var profile models.Profile
	err := row.Scan(&profile.ID, &profile.Name, &profile.Phone)
	return profile, err
}
