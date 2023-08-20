package entities

import "time"

type User struct {
	ID       uint   `json:"user_id" gorm:"column:user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestRefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}