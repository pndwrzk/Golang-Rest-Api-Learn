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

func (r *ProductRepositoryImpl) ReadProduct() ([]entities.Product, error) {
	get, err := r.ProductRepository.Read()

	if err != nil {
		return nil, err
	}

	return get, nil
}

func (r *ProductRepositoryImpl) CreateProduct(product entities.Product) (entities.Product, error) {
	get, err := r.ProductRepository.Create(product)
	return get, err
}
