package products

import (
	"errors"

	"github.com/biFebriansyah/goback/src/database/gorm/models"
	"gorm.io/gorm"
)

type product_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *product_repo {
	return &product_repo{grm}
}

func (r *product_repo) FindAll() (*models.Products, error) {
	var products models.Products

	result := r.db.Find(&products)

	if result.Error != nil {
		return nil, errors.New("Gagal mengambil data")
	}

	return &products, nil
}

func (r *product_repo) FindById(id int) (*models.Products, error) {
	var products models.Products

	result := r.db.First(&products, "product_id = ?", id)

	if result.Error != nil {
		return nil, errors.New("Gagal mengambil data")
	}

	return &products, nil
}

func (r *product_repo) Add(data *models.Product) (*models.Product, error) {

	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("Gagal menyimpan data")
	}

	return data, nil
}
