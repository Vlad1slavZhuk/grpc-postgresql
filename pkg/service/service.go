package service

import (
	"context"
	"time"

	"github.com/Vlad1slavZhuk/grpc-postgresql/domain"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/auth"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/hash"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/repository"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type UserManager interface {
	SignUp(ctx context.Context, input UserSignUpInput) error
	SignIn(ctx context.Context, input UserSignInInput) (Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (Tokens, error)
}

type ItemManager interface {
	CreateItem(ctx context.Context, input ItemInput) (domain.Item, error)
	UpdateItem(ctx context.Context, input ItemInput) (domain.Item, error)
	DeleteItem(ctx context.Context, itemID int32) error
	GetItem(ctx context.Context, itemID int32) (domain.Item, error)
}

type OrderManager interface {
	CreateOrder(ctx context.Context, input OrderInput) (domain.Order, error)
	UpdateOrder(ctx context.Context, input OrderInput) (domain.Order, error)
	DeleteOrder(ctx context.Context, orderID int32) error
	GetOrder(ctx context.Context, userID, orderID int32) (domain.Order, error)
	GetOrders(ctx context.Context, userID int32, start, end time.Time) (domain.Orders, error)
}

type Services struct {
	UserService  UserManager
	ItemService  ItemManager
	OrderService OrderManager
}

type Deps struct {
	Repos        *repository.Repositories
	Hasher       hash.PasswordHasher
	TokenManager auth.TokenManager
	// AccessTokenTTL  time.Duration
	// RefreshTokenTTL time.Duration
}

func NewServices(deps Deps) *Services {
	return &Services{
		UserService:  NewUserService(deps.Repos.UserRepo, deps.Hasher, deps.TokenManager),
		ItemService:  NewItemService(deps.Repos.ItemRepo),
		OrderService: NewOrderService(deps.Repos.OrderRepo),
	}
}
