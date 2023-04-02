package domain

import "time"

type User struct {
	ID           int64     `gorm:"id"`
	Username     string    `gorm:"username"`
	Password     string    `gorm:"password"`
	Mobile       string    `gorm:"mobile"`
	Email        string    `gorm:"email"`
	CreatedAt    time.Time `gorm:"created_at"`
	UpdatedAt    time.Time `gorm:"updated_at"`
	RefreshToken string    `gorm:"refresh_token"`
	LastActive   time.Time `gorm:"last_active;autoCreateTime"`
}
