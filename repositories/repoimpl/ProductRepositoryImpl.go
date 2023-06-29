package repoimpl

import (
	"go-learning-restapi/entities"
	"go-learning-restapi/repositories"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) repositories.ProductRepository {
	return &ProductRepositoryImpl{
		DB: db,
	}
}

func (r *ProductRepositoryImpl) View() ([]entities.Product, error) {
	row, err := r.DB.Raw("Select * From products").Rows()
	if err != nil {
		return nil, err
	}
	var products []entities.Product
	for row.Next() {
		var product entities.Product
		err := row.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil

}
