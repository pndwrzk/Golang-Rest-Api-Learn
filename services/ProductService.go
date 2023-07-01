package services

import "go-learning-restapi/entities"

type ProductService interface {
	ReadProduct() ([]entities.Product, error)
	CreateProduct(product entities.Product) (entities.Product, error)
}
