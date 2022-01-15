package product

import (
	"errors"
	"majoo/repository/outlet"
)

type service struct {
	repository 			Repository
	outletRepository 	outlet.Repository
}

type Service interface {
	FindProducts(userID int, input GetOutletProductInput) ([] Product, error)
	CreateProduct(userID int, inputID GetOutletProductInput, input CreateProductInput) (Product, error)
	SaveProductImage(input CreateProductImageInput, fileLocation string) (ProductImage, error)
}

func NewService(repository Repository, outletRepository outlet.Repository) *service {
	return &service{repository, outletRepository}
}

func (s *service) FindProducts(userID int, input GetOutletProductInput) ([]Product, error) {
	outlet, err := s.outletRepository.FindByID(userID, input.ID)
	if err != nil {
		return []Product{}, err
	}

	if outlet.UserID != input.User.ID {
		return []Product{}, errors.New("Not an owner of the outlet")
	}

	products, err := s.repository.GetByOutletID(input.ID)
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) CreateProduct(userID int, inputID GetOutletProductInput, input CreateProductInput) (Product, error) {
	outlet, err := s.outletRepository.FindByID(userID, input.OutletID)
	if err != nil {
		return Product{}, err
	}

	if outlet.UserID != inputID.User.ID {
		return Product{}, errors.New("Not an owner of the outlet")
	}
	
	product := Product{}
	product.Name = input.Name
	product.OutletID = input.OutletID
	product.Description = input.Description
	product.Stock = input.Stock
	product.Price = input.Price

	newProduct, err := s.repository.Save(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) SaveProductImage(input CreateProductImageInput, fileLocation string) (ProductImage, error) {
	// _, err := s.repository.FindById(input.ProductID)
	// if err != nil {
	// 	return ProductImage{}, err
	// }

	productImage := ProductImage{}
	productImage.ProductID = input.ProductID
	productImage.ProductFileName = fileLocation

	newProductImage, err := s.repository.CreateImage(productImage)
	if err != nil {
		return newProductImage, err
	}

	return newProductImage, nil
}