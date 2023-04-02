package service

import (
	"context"

	"github.com/Vlad1slavZhuk/grpc-postgresql/domain"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/repository"
)

type ItemInput struct {
	CategoryID  int32
	Name        string
	Description string
	Count       uint32
	Amount      float64
	Avaibility  bool
}

type ItemsService struct {
	repo repository.ItemManager
}

func NewItemService(repo repository.ItemManager) *ItemsService {
	return &ItemsService{repo: repo}
}

func (i *ItemsService) CreateItem(ctx context.Context, in ItemInput) (domain.Item, error) {
	return i.repo.CreateItem(ctx, domain.Item{
		CategoryID:  in.CategoryID,
		Name:        in.Name,
		Description: in.Name,
		Count:       in.Count,
		Amount:      in.Amount,
		Avaibility:  in.Avaibility,
	})
}
func (i *ItemsService) UpdateItem(ctx context.Context, in ItemInput) (domain.Item, error) {
	return i.repo.UpdateItem(ctx, domain.Item{
		CategoryID:  in.CategoryID,
		Name:        in.Name,
		Description: in.Name,
		Count:       in.Count,
		Amount:      in.Amount,
		Avaibility:  in.Avaibility,
	})
}
func (i *ItemsService) DeleteItem(ctx context.Context, itemID int32) error {
	return i.repo.DeleteItem(ctx, itemID)
}
func (i *ItemsService) GetItem(ctx context.Context, itemID int32) (domain.Item, error) {
	return i.repo.GetItem(ctx, itemID)
}
