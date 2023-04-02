package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Vlad1slavZhuk/grpc-postgresql/domain"
	"gorm.io/gorm"
)

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *itemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) CreateItem(ctx context.Context, item domain.Item) (domain.Item, error) {

	err := r.db.WithContext(ctx).Table("items").Create(&item).Error
	if err != nil {
		return domain.Item{}, err
	}

	return item, nil
}

func (r *itemRepository) UpdateItem(ctx context.Context, item domain.Item) (domain.Item, error) {
	var i domain.Item

	tx := r.db.WithContext(ctx).Table("items").Take(&i, "id = ?", item.ID)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return domain.Item{}, fmt.Errorf("not found item_id: %v", item.ID)
	}

	i = item

	return i, tx.Save(&i).Error
}

func (r *itemRepository) DeleteItem(ctx context.Context, id int32) error {
	var i domain.Item

	err := r.db.WithContext(ctx).Table("items").Delete(&i, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *itemRepository) GetItem(ctx context.Context, id int32) (domain.Item, error) {
	var i domain.Item

	err := r.db.WithContext(ctx).Table("items").First(&i, "id = ?", id).Error
	if err != nil {
		return domain.Item{}, err
	}

	return i, nil
}

func (r *itemRepository) GetItems(ctx context.Context) ([]*domain.Item, error) {
	var items []*domain.Item

	err := r.db.WithContext(ctx).Table("items").Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}
