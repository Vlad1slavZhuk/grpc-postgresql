package handler

import (
	pbShop "github.com/Vlad1slavZhuk/grpc-postgresql/api/gen/shop/v1"
	pbUser "github.com/Vlad1slavZhuk/grpc-postgresql/api/gen/user/v1"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/auth"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/service"
)

type Handler struct {
	services     *service.Services
	tokenManager auth.TokenManager

	pbUser.UnimplementedUserServiceServer
	pbShop.UnimplementedShopServiceServer
}

func NewHandler(services *service.Services, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}
