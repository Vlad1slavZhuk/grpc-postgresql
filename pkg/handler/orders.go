package handler

import (
	"context"

	pbShop "github.com/Vlad1slavZhuk/grpc-postgresql/api/gen/shop/v1"
	"github.com/Vlad1slavZhuk/grpc-postgresql/domain"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) CreateOrder(ctx context.Context, req *pbShop.CreateOrderRequest) (*pbShop.CreateOrderResponse, error) {
	status, err := domain.ValidStrStatus(req.Status)
	if err != nil {
		return nil, err
	}

	order, err := h.services.OrderService.CreateOrder(ctx, service.OrderInput{
		UserID:   UserIDFromCtx(ctx),
		ItemID:   req.ItemId,
		StatusID: status,
		Count:    req.Count,
		Amount:   req.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &pbShop.CreateOrderResponse{
		Order: &pbShop.Order{
			Id:        order.ID,
			UserId:    order.UserID,
			ItemId:    order.ItemID,
			Status:    domain.ValidIntStatus(order.StatusID),
			Count:     order.Count,
			Amount:    order.Amount,
			CreatedAt: timestamppb.New(order.CreatedAt),
			UpdatedAt: timestamppb.New(order.UpdatedAt),
		},
	}, nil
}

func (h *Handler) UpdateOrder(ctx context.Context, req *pbShop.UpdateOrderRequest) (*pbShop.UpdateOrderResponse, error) {
	status, err := domain.ValidStrStatus(req.Status)
	if err != nil {
		return nil, err
	}

	order, err := h.services.OrderService.UpdateOrder(ctx, service.OrderInput{
		UserID:   UserIDFromCtx(ctx),
		ItemID:   req.ItemId,
		StatusID: status,
		Count:    req.Count,
		Amount:   req.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &pbShop.UpdateOrderResponse{
		Order: &pbShop.Order{
			Id:        order.ID,
			UserId:    order.UserID,
			ItemId:    order.ItemID,
			Status:    domain.ValidIntStatus(order.StatusID),
			Count:     order.Count,
			Amount:    order.Amount,
			CreatedAt: timestamppb.New(order.CreatedAt),
			UpdatedAt: timestamppb.New(order.UpdatedAt),
		},
	}, nil
}

func (h *Handler) DeleteOrder(ctx context.Context, req *pbShop.DeleteOrderRequest) (*pbShop.DeleteOrderResponse, error) {

	err := h.services.OrderService.DeleteOrder(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}

	return &pbShop.DeleteOrderResponse{
		Status: "success",
	}, nil
}

func (h *Handler) GetOrder(ctx context.Context, req *pbShop.GetOrderRequest) (*pbShop.GetOrderResponse, error) {

	order, err := h.services.OrderService.GetOrder(ctx, UserIDFromCtx(ctx), req.OrderId)
	if err != nil {
		return nil, err
	}

	return &pbShop.GetOrderResponse{
		Order: &pbShop.Order{
			Id:        order.ID,
			UserId:    order.UserID,
			ItemId:    order.ItemID,
			Status:    domain.ValidIntStatus(order.StatusID),
			Count:     order.Count,
			Amount:    order.Amount,
			CreatedAt: timestamppb.New(order.CreatedAt),
			UpdatedAt: timestamppb.New(order.UpdatedAt),
		},
	}, nil
}

func (h *Handler) GetOrders(ctx context.Context, req *pbShop.GetOrdersRequest) (*pbShop.GetOrdersResponse, error) {

	orders, err := h.services.OrderService.GetOrders(ctx, UserIDFromCtx(ctx), req.Start.AsTime(), req.End.AsTime())
	if err != nil {
		return nil, err
	}

	return &pbShop.GetOrdersResponse{
		Orders:      convertOrdersToPB(orders.Orders),
		TotalAmount: orders.TotalAmount,
	}, nil
}
