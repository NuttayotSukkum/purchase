package services

import (
	"github.com/NuttayotSukkum/purchase/internal/constants"
	"github.com/NuttayotSukkum/purchase/internal/models/entities"
	"github.com/NuttayotSukkum/purchase/internal/models/requests"
	"github.com/NuttayotSukkum/purchase/internal/repositories"
	"github.com/google/uuid"
	logger "github.com/labstack/gommon/log"
	"time"
)

type PaymentServiceImpl struct {
	paymentRepository repositories.PaymentRepository
}

func NewPaymentServiceImpl(paymentRepository repositories.PaymentRepository) *PaymentServiceImpl {
	return &PaymentServiceImpl{
		paymentRepository: paymentRepository,
	}
}

func (repo *PaymentServiceImpl) CreatePayment(request requests.PaymentRequest, productPrice float64) (*entities.Payment, error) {
	paymentProduct := entities.Payment{
		ID:            uuid.New().String(),
		ProductID:     request.ProductId,
		NetPrice:      productPrice * float64(request.Amount),
		Amount:        request.Amount,
		PaymentStatus: "pending",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	payment, err := repo.paymentRepository.CreatePayment(paymentProduct)
	logger.Infof("payment:%+v", payment)
	if err != nil {
		logger.Errorf("Error to create Payment: %s", err)
		return nil, err
	}
	return payment, nil
}

func (repo *PaymentServiceImpl) FindPaymentByPaymentId(request requests.PaymentIdRequest) (*entities.Payment, error) {
	payment, err := repo.paymentRepository.FindPaymentById(request.PaymentId)
	if err != nil {
		logger.Errorf("Error to fething database: %s", err)
		return nil, err
	}
	return payment, nil
}

func (repo *PaymentServiceImpl) UpdatePaymentStatus(paymentId string, status string) *error {
	if status == constants.SUCCESS_STATUS {
		if _, err := repo.paymentRepository.UpdatePaymentStatus(paymentId, constants.SUCCESS_STATUS); err != nil {
			logger.Errorf("Error Update data from database: %s", err)
			return &err
		}
	} else if status == constants.FAILED_STATUS {
		if _, err := repo.paymentRepository.UpdatePaymentStatus(paymentId, constants.FAILED_STATUS); err != nil {
			logger.Errorf("Error Update data from database: %s", err)
			return &err
		}
	}
	return nil
}

//func buildProductCreateResponse(product entities.Product) dto2.ProductCreateResponse {
//	return dto2.ProductCreateResponse{
//		Id:     product.Id,
//		Name:   product.Name,
//		Price:  product.Price,
//		Amount: product.Amount,
//	}
//}
//
//func (p *ProductManagementServiceImpl) EditProduct(request dto2.ProductEditRequest) {
//
//}
