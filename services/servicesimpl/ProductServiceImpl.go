package servicesimpl

import (
	"go-learning-restapi/entities"
	"go-learning-restapi/repositories"
	"go-learning-restapi/services"
)

type ProductRepositoryImpl struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) services.ProductService {
	return &ProductRepositoryImpl{
		productRepo: productRepo,
	}
}

func (r *ProductRepositoryImpl) ViewProduct() ([]entities.Product, error) {
	get, err := r.productRepo.View()

	if err != nil {
		return nil, err
	}

	return get, nil
}
