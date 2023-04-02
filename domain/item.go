package domain

import "time"

type Item struct {
	ID          int32     `gorm:"id"`
	CategoryID  int32     `gorm:"category_id"`
	Name        string    `gorm:"name"`
	Description string    `gorm:"description"`
	Count       uint32    `gorm:"count"`
	Amount      float64   `gorm:"amount"`
	Avaibility  bool      `gorm:"avaibility"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
}
