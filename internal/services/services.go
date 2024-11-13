package services

import (
	"github.com/NuttayotSukkum/purchase/internal/models/entities"
	"github.com/NuttayotSukkum/purchase/internal/models/requests"
	"github.com/NuttayotSukkum/purchase/internal/models/responses"
)

type ProductService interface {
	CreateProduct(req requests.ProductRequest) (*entities.Product, error)
	FindProductByName(req requests.ProductRequest) (*entities.Product, error)
	EditProduct(name *string, amount int, price float64) (*entities.Product, error)
	FindProductByProductId(productId string) *entities.Product
	SearchProductByName(productName string) (*[]responses.ProductGetIdResp, error)
}

type PaymentService interface {
	CreatePayment(request requests.PaymentRequest, productPrice float64) (*entities.Payment, error)
	FindPaymentByPaymentId(request requests.PaymentIdRequest) (*entities.Payment, error)
	UpdatePaymentStatus(paymentId string, status string) *error
}
