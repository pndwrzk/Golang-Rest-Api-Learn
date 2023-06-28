package entities

type Customer struct {
	ID          uint   `json:"user_id" gorm:"column:user_id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"Address"`
	Email       string `json:"email"`
}