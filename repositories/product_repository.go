package repositories

import "go-learning-restapi/entities"

type ProductRepository interface {
	Read() ([]entities.Product, error)
	Create(product entities.Product) (entities.Product, error)
}
