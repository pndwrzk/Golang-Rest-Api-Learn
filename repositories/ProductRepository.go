package repositories

import "go-learning-restapi/entities"

type ProductRepository interface {
	View() ([]entities.Product, error)
}
