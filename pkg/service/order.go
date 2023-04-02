package service

import (
	"context"
	"time"

	"github.com/Vlad1slavZhuk/grpc-postgresql/domain"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/repository"
)

type OrderInput struct {
	UserID   int32
	ItemID   int32
	StatusID int32
	Count    uint32
	Amount   float64
}

type OrderService struct {
	repo repository.OrderManager
}

func NewOrderService(repo repository.OrderManager) *OrderService {
	return &OrderService{repo: repo}
}

func (o *OrderService) CreateOrder(ctx context.Context, input OrderInput) (domain.Order, error) {
	return o.repo.CreateOrder(ctx, domain.Order{
		UserID:   input.UserID,
		ItemID:   input.ItemID,
		StatusID: input.StatusID,
		Count:    input.Count,
		Amount:   input.Amount,
	})
}

func (o *OrderService) UpdateOrder(ctx context.Context, input OrderInput) (domain.Order, error) {
	return o.repo.UpdateOrder(ctx, domain.Order{
		UserID:   input.UserID,
		ItemID:   input.ItemID,
		StatusID: input.StatusID,
		Count:    input.Count,
		Amount:   input.Amount,
	})
}

func (o *OrderService) DeleteOrder(ctx context.Context, orderID int32) error {
	return o.repo.DeleteOrder(ctx, orderID)
}

func (o *OrderService) GetOrder(ctx context.Context, userID, orderID int32) (domain.Order, error) {
	return o.repo.GetOrder(ctx, userID, orderID)
}
func (o *OrderService) GetOrders(ctx context.Context, userID int32, start, end time.Time) (domain.Orders, error) {
	return o.repo.GetOrders(ctx, userID, start, end)
}
