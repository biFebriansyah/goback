package models

import "time"

type User struct {
	UserId    uint      `gorm:"primaryKey" json:"id,omitempty"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Password  string    `json:"password,omitempty" validate:"required"`
	Carts     Carts     `gorm:"foreignKey:UsersId;"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type Users []User
