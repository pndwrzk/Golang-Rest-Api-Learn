package servicesimpl

import (
	"go-learning-restapi/entities"
	"go-learning-restapi/repositories"
	"go-learning-restapi/services"
)

type ProductRepositoryImpl struct {
	ProductRepository repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) services.ProductService {
	return &ProductRepositoryImpl{
		ProductRepository: productRepo,
	}
}

func (r *ProductRepositoryImpl) ViewProduct() ([]entities.Product, error) {
	get, err := r.ProductRepository.View()

	if err != nil {
		return nil, err
	}

	return get, nil
}
