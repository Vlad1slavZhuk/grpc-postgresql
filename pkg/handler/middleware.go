package handler

import (
	"context"

	logger "github.com/Vlad1slavZhuk/grpc-postgresql/pkg/log"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserToken struct{}

var ignoreMethod = map[string]bool{
	"/user.v1.UserService/SignUp":       true,
	"/user.v1.UserService/SignIn":       true,
	"/user.v1.UserService/RefreshToken": true,
}

func (h *Handler) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	log := logger.GetLoggerInstance()
	log.Info().
		Str("method-request", fullMethodName).
		Msg("log handle request")

	if v, ok := ignoreMethod[fullMethodName]; ok && v {
		return ctx, nil
	}

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	userID, err := h.tokenManager.Parse(token)
	if err != nil {
		log.Error().
			Err(err).
			Msg("error parse token")
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	return context.WithValue(ctx, UserToken{}, userID), nil
}
