package requests

import (
	"github.com/NuttayotSukkum/purchase/internal/models/entities"
	"github.com/google/uuid"
	"time"
)

type PaymentRequest struct {
	ProductId string `json:"product_id"`
	Amount    int    `json:"amount"`
}

func (req *PaymentRequest) BuildPayment() entities.Payment {
	return entities.Payment{
		ID:            uuid.New().String(),
		ProductID:     req.ProductId,
		Amount:        req.Amount,
		PaymentStatus: "pending",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

type PaymentIdRequest struct {
	PaymentId string `json:"paymentId"`
}
