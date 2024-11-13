package db

import (
	"github.com/NuttayotSukkum/purchase/internal/models/entities"
	"gorm.io/gorm"
	"time"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (p *ProductRepositoryImpl) CreateProduct(product entities.Product) (entities.Product, error) {
	err := p.db.Create(&product).Error
	return product, err
}
func (repo *ProductRepositoryImpl) FindProductByName(name string) (*entities.Product, error) {
	var productResponse entities.Product
	if err := repo.db.Where("name = ?", name).First(&productResponse).Error; err != nil {
		return &entities.Product{}, err
	}
	return &productResponse, nil
}

func (repo *ProductRepositoryImpl) EditProduct(name *string, amount int, price float64) (*entities.Product, error) {
	var productResponse entities.Product
	if err := repo.db.Where("name = ?", name).First(&productResponse).Error; err != nil {
		return nil, err
	}

	if err := repo.db.Model(&productResponse).Where("id = ?", productResponse.ID).Updates(map[string]interface{}{
		"name":       name,
		"price":      price,
		"amount":     amount,
		"updated_at": time.Now(),
	}).Error; err != nil {
		return nil, err
	}

	return &productResponse, nil
}

func (repo *ProductRepositoryImpl) FindProductById(productId string) (*entities.Product, error) {
	var productResponse entities.Product
	if err := repo.db.Where("id = ?", productId).First(&productResponse).Error; err != nil {
		return nil, err
	}
	return &productResponse, nil
}

func (repo *ProductRepositoryImpl) PartialSearchProduct(productName string) (*[]entities.Product, error) {
	var listProduct []entities.Product
	if err := repo.db.Where("name LIKE ?", "%"+productName+"%").Limit(10).Find(&listProduct).Error; err != nil {
		return nil, err
	}
	return &listProduct, nil
}
