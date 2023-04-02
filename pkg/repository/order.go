package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Vlad1slavZhuk/grpc-postgresql/domain"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(ctx context.Context, o domain.Order) (domain.Order, error) {

	err := r.db.WithContext(ctx).Table("orders").Create(&o).Error
	if err != nil {
		return domain.Order{}, err
	}

	return o, nil
}

func (r *orderRepository) UpdateOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	var o domain.Order

	tx := r.db.WithContext(ctx).Take(&o, "id = ?", order.ID)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return domain.Order{}, fmt.Errorf("not found order_id: %v", order.ID)
	}

	o = order

	return o, tx.Save(&o).Error
}

func (r *orderRepository) DeleteOrder(ctx context.Context, orderID int32) error {
	var o domain.Item

	err := r.db.WithContext(ctx).Table("orders").Delete(&o, "id = ?", orderID).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) GetOrder(ctx context.Context, userID, orderID int32) (domain.Order, error) {
	var o domain.Order

	err := r.db.WithContext(ctx).Table("orders").First(&o, "id = ? AND user_id = ?", orderID, userID).Error
	if err != nil {
		return domain.Order{}, err
	}

	return o, nil
}

func (r *orderRepository) GetOrders(ctx context.Context, userID int32, start, end time.Time) (domain.Orders, error) {
	var orders []*domain.Order

	err := r.db.WithContext(ctx).Table("orders").Where("dt_inserted BETWEEN ? AND ?").Find(&orders, "user_id = ?", userID).Error
	if err != nil {
		return domain.Orders{}, err
	}

	return domain.Orders{Orders: orders}, nil
}
