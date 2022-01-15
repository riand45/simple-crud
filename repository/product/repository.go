package product

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetByOutletID(outletID int) ([]Product, error)
	FindById(id int) (Product, error)
	Save(product Product) (Product, error)
	CreateImage(productImage ProductImage) (ProductImage, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByOutletID(outletID int) ([]Product, error) {
	var products []Product
	
	err := r.db.Where("outlet_id = ?", outletID).Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) FindById(id int) (Product, error) {
	var product Product

	err := r.db.Where("id = ?", id).Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Save(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) CreateImage(productImage ProductImage) (ProductImage, error) {
	err := r.db.Create(&productImage).Error
	if err != nil {
		return productImage, err
	}

	return productImage, nil
}