package handler

import (
	"context"

	pbShop "github.com/Vlad1slavZhuk/grpc-postgresql/api/gen/shop/v1"
	"github.com/Vlad1slavZhuk/grpc-postgresql/domain"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/service"
)

func (h *Handler) CreateItem(ctx context.Context, req *pbShop.CreateItemRequest) (*pbShop.CreateItemResponse, error) {

	category, err := domain.ValidStrCategory(req.Category)
	if err != nil {
		return nil, err
	}

	item, err := h.services.ItemService.CreateItem(ctx, service.ItemInput{
		CategoryID:  category,
		Name:        req.Name,
		Description: req.Description,
		Count:       req.Count,
		Amount:      req.Amount,
		Avaibility:  req.Avaibility,
	})
	if err != nil {
		return nil, err
	}

	return &pbShop.CreateItemResponse{
		Item: &pbShop.Item{
			Id:          item.ID,
			Category:    domain.ValidIntCategory(item.CategoryID),
			Name:        item.Name,
			Description: item.Description,
			Count:       item.Count,
			Amount:      item.Amount,
			Avaibility:  item.Avaibility,
		},
	}, nil
}
func (h *Handler) UpdateItem(ctx context.Context, req *pbShop.UpdateItemRequest) (*pbShop.UpdateItemResponse, error) {

	category, err := domain.ValidStrCategory(req.Category)
	if err != nil {
		return nil, err
	}

	item, err := h.services.ItemService.UpdateItem(ctx, service.ItemInput{
		CategoryID:  category,
		Name:        req.Name,
		Description: req.Description,
		Count:       req.Count,
		Amount:      req.Amount,
		Avaibility:  req.Avaibility,
	})
	if err != nil {
		return nil, err
	}

	return &pbShop.UpdateItemResponse{
		Item: &pbShop.Item{
			Id:          item.ID,
			Category:    domain.ValidIntCategory(item.CategoryID),
			Name:        item.Name,
			Description: item.Description,
			Count:       item.Count,
			Amount:      item.Amount,
			Avaibility:  item.Avaibility,
		},
	}, nil
}
func (h *Handler) DeleteItem(ctx context.Context, req *pbShop.DeleteItemRequest) (*pbShop.DeleteItemResponse, error) {

	err := h.services.ItemService.DeleteItem(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pbShop.DeleteItemResponse{
		Status: "success",
	}, nil
}
func (h *Handler) GetItem(ctx context.Context, req *pbShop.GetItemRequest) (*pbShop.GetItemResponse, error) {

	item, err := h.services.ItemService.GetItem(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pbShop.GetItemResponse{
		Item: &pbShop.Item{
			Id:          item.ID,
			Category:    domain.ValidIntCategory(item.CategoryID),
			Name:        item.Name,
			Description: item.Description,
			Count:       item.Count,
			Amount:      item.Amount,
			Avaibility:  item.Avaibility,
		},
	}, nil
}
