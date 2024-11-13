package entities

import (
	"time"
)

type Payment struct {
	ID            string    `json:"id" gorm:"type:varchar(50);primaryKey"`
	ProductID     string    `json:"product_id" gorm:"type:uuid;not null"`
	NetPrice      float64   `json:"net_price" gorm:"type:numeric(10,2);not null"`
	Amount        int       `json:"amount" gorm:"not null"`
	PaymentStatus string    `json:"payment_status" gorm:"type:varchar(255)"`
	CreatedAt     time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp;not null"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp;not null"`
}

func (t *Payment) TableName() string {
	return "payment"
}
