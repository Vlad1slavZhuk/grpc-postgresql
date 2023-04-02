package repository

import (
	"context"
	"time"

	"github.com/Vlad1slavZhuk/grpc-postgresql/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, u domain.User) error {

	err := r.db.WithContext(ctx).Table("users").Create(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UpdateUser(ctx context.Context, u domain.User) error {

	err := r.db.WithContext(ctx).Table("users").Updates(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, userID int64) error {
	var u domain.User

	err := r.db.WithContext(ctx).Table("users").Delete(&u, "id = ?", userID).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUser(ctx context.Context, userID int64) (domain.User, error) {
	var u domain.User

	err := r.db.WithContext(ctx).Table("users").First(&u, "id = ?", userID).Error
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (r *userRepository) GetUsers(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User

	if err := r.db.WithContext(ctx).Table("users").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetByCredentials(ctx context.Context, email, hash string) (domain.User, error) {
	var u domain.User

	err := r.db.WithContext(ctx).Table("users").
		Where("email = ?", email).
		Where("password = ?", hash).Take(&u).Error
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (r *userRepository) GetByRefreshToken(ctx context.Context, refreshToken string) (domain.User, error) {
	var u domain.User

	err := r.db.WithContext(ctx).Table("users").First(&u, "refresh_token = ?", refreshToken).Error
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (r *userRepository) SetSession(ctx context.Context, userID int64, session domain.Session) error {
	err := r.db.WithContext(ctx).Table("users").Where("id = ?", userID).Updates(map[string]interface{}{
		"last_active":   time.Now(),
		"refresh_token": session.RefreshToken,
	}).Error
	if err != nil {
		return err
	}

	return nil
}
