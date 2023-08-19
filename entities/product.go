package entities

type Product struct {
	ID          uint   `json:"product_id" gorm:"column:product_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
}
