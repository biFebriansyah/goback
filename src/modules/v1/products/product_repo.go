package products

import (
	"errors"

	"gorm.io/gorm"
)

type product_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *product_repo {
	return &product_repo{grm}
}

func (r *product_repo) FindAll() (*Products, error) {
	var products Products

	result := r.db.Find(&products)

	if result.Error != nil {
		return nil, errors.New("Gagal mengambil data")
	}

	return &products, nil
}

func (r *product_repo) Add(data *Product) (*Product, error) {

	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("Gagal menyimpan data")
	}

	return data, nil
}
