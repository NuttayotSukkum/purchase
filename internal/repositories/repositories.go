package repositories

import (
	"github.com/NuttayotSukkum/purchase/internal/models/entities"
	"github.com/google/uuid"
)

type ProductRepository interface {
	CreateProduct(product entities.Product) (entities.Product, error)
	FindProductByName(name string) (*entities.Product, error)
	EditProduct(name *string, amount int, price float64) (*entities.Product, error)
	FindProductById(productId string) (*entities.Product, error)
	PartialSearchProduct(productName string) (*[]entities.Product, error)
	//EditProduct(product entities.Product)
	//FindById(id uuid.UUID) (*entities.Product, error)
	//GetProductById(id string) (*entities.Product, error)
}

type PaymentRepository interface {
	CreatePayment(payment entities.Payment) (*entities.Payment, error)
	UpdatePaymentStatus(productId string, status string) (*entities.Payment, error)
	FindPaymentById(Id string) (*entities.Payment, error)
	UpdatePaymentAmount(productId uuid.UUID, totalAmount float64) (*entities.Payment, error)
}
