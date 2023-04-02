package domain

import "time"

type Order struct {
	ID        int32     `gorm:"id,omitempty"`
	UserID    int32     `gorm:"user_id,omitempty"`
	ItemID    int32     `gorm:"item_id,omitempty"`
	StatusID  int32     `gorm:"status_id,omitempty"`
	Count     uint32    `gorm:"count,omitempty"`
	Amount    float64   `gorm:"amount,omitempty"`
	CreatedAt time.Time `gorm:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"updated_at,omitempty"`
}

type Orders struct {
	Orders      []*Order
	TotalAmount float64
}

func (o *Orders) CalculateTotalAmount() {
	var result float64

	for _, o := range o.Orders {
		result += o.Amount * float64(o.Count)
	}

	o.TotalAmount = result
}
