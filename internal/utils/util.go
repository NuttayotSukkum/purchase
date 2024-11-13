package utils

import (
	"github.com/NuttayotSukkum/purchase/internal/constants"
	"github.com/NuttayotSukkum/purchase/internal/models/entities"
	"github.com/NuttayotSukkum/purchase/internal/models/requests"
)

func CalculationAmount(paymentAmount, productAmount float64) float64 {
	totalAmount := paymentAmount - productAmount
	return totalAmount
}

func CheckProduct(productReq *requests.ProductRequest, productDb *entities.Product, checkStatus string) bool {
	switch checkStatus {
	case constants.CHECK_PRODUCT:
		return productReq.Name == "" || productReq.Amount < 0 || productReq.Price <= 0

	case constants.CHECK_PRODUCT_EXIST:
		if productDb == nil {
			return false
		}
		return productReq.Name == productDb.Name && productReq.Price == productDb.Price

	default:
		return false
	}
}
