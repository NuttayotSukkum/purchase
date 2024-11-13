package requests

import (
	"github.com/NuttayotSukkum/purchase/internal/models/entities"
	"github.com/google/uuid"
	"time"
)

type ProductRequest struct {
	Name   string  `json:"name"`
	Amount int     `json:"amount"`
	Price  float64 `json:"price"`
}

func (req *ProductRequest) BuildProduct() entities.Product {
	return entities.Product{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Amount:    req.Amount,
		Price:     req.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type IdRequest struct {
	Id string `json:"id"`
}

type ProductNameRequest struct {
	ProductName string `json:"product_name"`
}
