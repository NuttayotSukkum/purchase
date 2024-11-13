package db

import (
	"github.com/NuttayotSukkum/purchase/internal/models/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentRepositoryImpl struct {
	db *gorm.DB
}

func NewPaymentRepositoryImpl(db *gorm.DB) *PaymentRepositoryImpl {
	return &PaymentRepositoryImpl{
		db: db,
	}
}

func (repo *PaymentRepositoryImpl) CreatePayment(payment entities.Payment) (*entities.Payment, error) {
	if err := repo.db.Create(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (repo *PaymentRepositoryImpl) UpdatePaymentStatus(productId string, status string) (*entities.Payment, error) {
	var updatedPayment entities.Payment
	result := repo.db.Model(&updatedPayment).
		Where("id = ?", productId).
		Update("payment_status", status)

	if result.Error != nil {
		return nil, result.Error
	}

	if err := repo.db.Where("id = ?", productId).First(&updatedPayment).Error; err != nil {
		return nil, err
	}

	return &updatedPayment, nil
}

func (repo *PaymentRepositoryImpl) FindPaymentById(Id string) (*entities.Payment, error) {
	var payment entities.Payment
	if err := repo.db.Where("id = ?", Id).First(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (repo *PaymentRepositoryImpl) UpdatePaymentAmount(productId uuid.UUID, totalAmount float64) (*entities.Payment, error) {
	result := repo.db.Model(&entities.Payment{}).
		Where("product_id = ?", productId).
		Update("net_price", totalAmount)

	if result.Error != nil {
		return nil, result.Error
	}

	var updatedPayment entities.Payment
	if err := repo.db.Where("product_id = ?", productId).First(&updatedPayment).Error; err != nil {
		return nil, err
	}

	return &updatedPayment, nil
}
