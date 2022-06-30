package users

import (
	"errors"

	"github.com/biFebriansyah/goback/src/database/gorm/models"
	"gorm.io/gorm"
)

type users_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *users_repo {
	return &users_repo{grm}
}

func (r *users_repo) FindAll() (*models.Users, error) {
	var data models.Users

	result := r.db.Order("created_at desc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("Gagal mengambil data")
	}

	return &data, nil
}

func (r *users_repo) FindByUsername(username string) (*models.User, error) {
	var data models.User

	result := r.db.First(&data, "username = ?", username)

	if result.Error != nil {
		return nil, errors.New("Gagal mengambil data")
	}

	return &data, nil
}

func (r *users_repo) UserExsist(username string) bool {
	var data models.User
	result := r.db.First(&data, "username = ?", username)

	if result.Error != nil {
		return false
	}

	return true
}

func (r *users_repo) Add(data *models.User) (*models.User, error) {

	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("Gagal menyimpan data")
	}

	return data, nil
}
