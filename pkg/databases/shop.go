package databases

import (
	"context"

	"github.com/Vlad1slavZhuk/grpc-postgresql/domain"
	"gorm.io/gorm"
)

type ShopRepository interface {
	CreateOrder(ctx context.Context, u domain.User) error
	UpdateOrder(ctx context.Context, u domain.User) error
	GetOrder(ctx context.Context, userID, orderID int64) (domain.User, error)
	GetOrders(ctx context.Context) ([]*domain.User, error)

	AddItem(ctx context.Context) error
	UpdateItem(ctx context.Context) error
	DeleteItem(ctx context.Context) error
	PriceList(ctx context.Context) error
}

type shopRepository struct {
	DB *gorm.DB
}

func NewShopRepository(db *gorm.DB) ShopRepository {
	return &shopRepository{DB: db}
}

// AddItem implements ShopRepository
func (*shopRepository) AddItem(ctx context.Context) error {
	panic("unimplemented")
}

// CreateOrder implements ShopRepository
func (*shopRepository) CreateOrder(ctx context.Context, u domain.User) error {
	panic("unimplemented")
}

// DeleteItem implements ShopRepository
func (*shopRepository) DeleteItem(ctx context.Context) error {
	panic("unimplemented")
}

// GetOrder implements ShopRepository
func (*shopRepository) GetOrder(ctx context.Context, userID, orderID int64) (domain.User, error) {
	panic("unimplemented")
}

// GetOrders implements ShopRepository
func (*shopRepository) GetOrders(ctx context.Context) ([]*domain.User, error) {
	panic("unimplemented")
}

// PriceList implements ShopRepository
func (*shopRepository) PriceList(ctx context.Context) error {
	panic("unimplemented")
}

// UpdateItem implements ShopRepository
func (*shopRepository) UpdateItem(ctx context.Context) error {
	panic("unimplemented")
}

// UpdateOrder implements ShopRepository
func (*shopRepository) UpdateOrder(ctx context.Context, u domain.User) error {
	panic("unimplemented")
}
