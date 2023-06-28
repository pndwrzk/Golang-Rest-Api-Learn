package services

import "go-learning-restapi/entities"

type ProductService interface {
	ViewProduct() ([]entities.Product, error)
}