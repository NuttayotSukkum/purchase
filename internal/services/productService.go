package services

import (
	"errors"
	"github.com/NuttayotSukkum/purchase/internal/models/entities"
	"github.com/NuttayotSukkum/purchase/internal/models/requests"
	"github.com/NuttayotSukkum/purchase/internal/models/responses"
	"github.com/NuttayotSukkum/purchase/internal/repositories"
	logger "github.com/labstack/gommon/log"
)

type ProductServiceImpl struct {
	productRepository repositories.ProductRepository
}

func NewProductServiceImpl(productRepository repositories.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{
		productRepository: productRepository,
	}
}

func (repo *ProductServiceImpl) CreateProduct(req requests.ProductRequest) (*entities.Product, error) {
	product, err := repo.productRepository.CreateProduct(req.BuildProduct())
	if err != nil {
		logger.Errorf("unable to create product: %s", err)
		return nil, errors.New("unable to create product:" + err.Error())
	}
	return &product, nil
}

func (repo *ProductServiceImpl) FindProductByName(req requests.ProductRequest) (*entities.Product, error) {
	product, err := repo.productRepository.FindProductByName(req.Name)
	if err != nil {
		return nil, err
	}
	logger.Errorf("Product:%v", product)
	return product, err
}

func (repo *ProductServiceImpl) EditProduct(name *string, amount int, price float64) (*entities.Product, error) {
	product, err := repo.productRepository.EditProduct(name, amount, price)
	if err != nil {
		logger.Errorf("Error Fething Data from Database %v", err)
		return nil, err
	}
	return product, nil
}

func (repo *ProductServiceImpl) FindProductByProductId(productId string) *entities.Product {
	product, err := repo.productRepository.FindProductById(productId)
	if err != nil {
		return &entities.Product{}
	}
	return product
}

func (repo *ProductServiceImpl) SearchProductByName(productName string) (*[]responses.ProductGetIdResp, error) {
	listOfProduct, err := repo.productRepository.PartialSearchProduct(productName)
	if err != nil {
		logger.Errorf("Error fetching Data from database: %s", err)
		return nil, err
	}
	var productResponse []responses.ProductGetIdResp
	for _, product := range *listOfProduct {
		if product.Amount > 0 {
			productResponse = append(productResponse, responses.ProductGetIdResp{
				Id:     product.ID,
				Name:   product.Name,
				Amount: product.Amount,
				Price:  product.Price,
				Status: "available",
			})
		} else {
			productResponse = append(productResponse, responses.ProductGetIdResp{
				Id:     product.ID,
				Name:   product.Name,
				Amount: product.Amount,
				Price:  product.Price,
				Status: "sold_out",
			})
		}

	}
	return &productResponse, nil
}
