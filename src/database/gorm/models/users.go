package models

import "time"

type User struct {
	UserId    uint      `gorm:"primaryKey" json:"id,omitempty"`
	Username  string    `json:"Username" validate:"required"`
	Email     string    `json:"Email"`
	Password  string    `json:"Password" validate:"required"`
	Carts     Carts     `gorm:"foreignKey:UsersId;"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type Users []User
