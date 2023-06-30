package entities

type Customer struct {
	ID          uint   `json:"customer_id" gorm:"column:customer_id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"Address"`
	Email       string `json:"email"`
}
