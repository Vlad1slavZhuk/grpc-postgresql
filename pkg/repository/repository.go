package repository

import (
	"context"
	"time"

	"github.com/Vlad1slavZhuk/grpc-postgresql/domain"
	"gorm.io/gorm"
)

type OrderManager interface {
	CreateOrder(ctx context.Context, o domain.Order) (domain.Order, error)
	UpdateOrder(ctx context.Context, o domain.Order) (domain.Order, error)
	DeleteOrder(ctx context.Context, orderID int32) error
	GetOrder(ctx context.Context, userID, orderID int32) (domain.Order, error)
	GetOrders(ctx context.Context, userID int32, start, end time.Time) (domain.Orders, error)
}

type ItemManager interface {
	CreateItem(ctx context.Context, item domain.Item) (domain.Item, error)
	UpdateItem(ctx context.Context, item domain.Item) (domain.Item, error)
	DeleteItem(ctx context.Context, id int32) error
	GetItem(ctx context.Context, id int32) (domain.Item, error)
}

type UserManager interface {
	CreateUser(ctx context.Context, u domain.User) error
	UpdateUser(ctx context.Context, u domain.User) error
	DeleteUser(ctx context.Context, userID int64) error
	GetUser(ctx context.Context, userID int64) (domain.User, error)
	GetUsers(ctx context.Context) ([]*domain.User, error)

	SetSession(ctx context.Context, userID int64, session domain.Session) error
	GetByRefreshToken(ctx context.Context, refreshToken string) (domain.User, error)
	GetByCredentials(ctx context.Context, email, hash string) (domain.User, error)
}

type Repositories struct {
	OrderRepo OrderManager
	ItemRepo  ItemManager
	UserRepo  UserManager
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		OrderRepo: NewOrderRepository(db),
		ItemRepo:  NewItemRepository(db),
		UserRepo:  NewUsersRepository(db),
	}
}
