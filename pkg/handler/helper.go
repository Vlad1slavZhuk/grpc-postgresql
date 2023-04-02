package handler

import (
	"context"
	"strconv"

	pbShop "github.com/Vlad1slavZhuk/grpc-postgresql/api/gen/shop/v1"
	"github.com/Vlad1slavZhuk/grpc-postgresql/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserIDFromCtx(ctx context.Context) int32 {
	userID, _ := strconv.Atoi(ctx.Value(UserToken{}).(string))
	return int32(userID)
}

func convertOrdersToPB(orders []*domain.Order) []*pbShop.Order {
	result := make([]*pbShop.Order, 0, len(orders))

	for _, o := range orders {
		result = append(result, &pbShop.Order{
			Id:        o.ID,
			UserId:    o.UserID,
			ItemId:    o.ItemID,
			Status:    domain.ValidIntStatus(o.StatusID),
			Count:     o.Count,
			Amount:    o.Amount,
			CreatedAt: timestamppb.New(o.CreatedAt),
			UpdatedAt: timestamppb.New(o.UpdatedAt),
		})
	}

	return result
}
