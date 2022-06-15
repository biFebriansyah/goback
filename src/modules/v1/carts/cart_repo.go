package carts

import (
	"errors"

	"github.com/biFebriansyah/goback/src/database/gorm/models"
	"gorm.io/gorm"
)

type cart_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *cart_repo {
	return &cart_repo{grm}
}

func (r *cart_repo) FindByUserId(id int) (*models.Cart, error) {
	var cart models.Cart

	data := r.db.Preload("Products").Where("users_id = ?", id).Find(&cart)

	if data.Error != nil {
		return nil, errors.New("Gagal mengambil data")
	}

	return &cart, nil
}

func (r *cart_repo) Save(data *models.Cart) (*models.Cart, error) {

	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("Gagal mengambil data")
	}

	return data, nil
}
