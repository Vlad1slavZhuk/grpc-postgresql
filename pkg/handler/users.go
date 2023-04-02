package handler

import (
	"context"

	pbUser "github.com/Vlad1slavZhuk/grpc-postgresql/api/gen/user/v1"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/service"
)

func (h *Handler) SignUp(ctx context.Context, req *pbUser.SignUpRequest) (*pbUser.SignUpResponse, error) {
	err := h.services.UserService.SignUp(ctx, service.UserSignUpInput{
		Name:     req.User.Username,
		Email:    req.User.Email,
		Phone:    req.User.Mobile,
		Password: req.User.Password,
	})
	if err != nil {
		return nil, err
	}

	return &pbUser.SignUpResponse{
		User: req.User,
	}, nil
}

func (h *Handler) SignIn(ctx context.Context, req *pbUser.SignInRequest) (*pbUser.SignInResponse, error) {
	tokens, err := h.services.UserService.SignIn(ctx, service.UserSignInInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &pbUser.SignInResponse{
		Jwt:          tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (h *Handler) RefreshToken(ctx context.Context, req *pbUser.RefreshTokenRequest) (*pbUser.RefreshTokenResponse, error) {
	res, err := h.services.UserService.RefreshTokens(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &pbUser.RefreshTokenResponse{
		Jwt:          res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}
