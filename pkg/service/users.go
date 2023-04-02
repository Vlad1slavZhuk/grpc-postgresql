package service

import (
	"context"
	"errors"
	"time"

	"github.com/Vlad1slavZhuk/grpc-postgresql/domain"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/auth"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/hash"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserSignUpInput struct {
	Name     string
	Email    string
	Phone    string
	Password string
}

type UserSignInInput struct {
	Email    string
	Password string
}

type UsersService struct {
	repo         repository.UserManager
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager

	accessTokenTTL time.Duration
	// refreshTokenTTL time.Duration
}

func NewUserService(
	repo repository.UserManager,
	hasher hash.PasswordHasher,
	tokenManager auth.TokenManager,
) *UsersService {
	return &UsersService{
		repo:         repo,
		hasher:       hasher,
		tokenManager: tokenManager,
	}
}

func (us *UsersService) SignUp(ctx context.Context, input UserSignUpInput) error {
	passwordHash, err := us.hasher.Hash(input.Password)
	if err != nil {
		return err
	}

	user := domain.User{
		Username: input.Name,
		Password: passwordHash,
		Mobile:   input.Phone,
		Email:    input.Email,
	}

	return us.repo.CreateUser(ctx, user)
}

func (us *UsersService) SignIn(ctx context.Context, input UserSignInInput) (Tokens, error) {
	passwordHash, err := us.hasher.Hash(input.Password)
	if err != nil {
		return Tokens{}, err
	}

	user, err := us.repo.GetByCredentials(ctx, input.Email, passwordHash)
	if err != nil {
		return Tokens{}, err
	}

	return us.createSession(ctx, user.ID)
}

func (us *UsersService) RefreshTokens(ctx context.Context, refreshToken string) (Tokens, error) {
	user, err := us.repo.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Tokens{}, status.Error(codes.NotFound, err.Error())
		}
		return Tokens{}, err
	}

	return us.createSession(ctx, user.ID)
}

func (us *UsersService) createSession(ctx context.Context, userID int64) (Tokens, error) {
	var (
		tokens Tokens
		err    error
	)

	tokens.AccessToken, err = us.tokenManager.NewJWT(userID, us.accessTokenTTL)
	if err != nil {
		return tokens, err
	}

	tokens.RefreshToken, err = us.tokenManager.NewRefreshToken()
	if err != nil {
		return tokens, err
	}

	return tokens, us.repo.SetSession(ctx, userID, domain.Session{
		RefreshToken: tokens.RefreshToken,
	})
}
