package carts

import (
	"strconv"

	"github.com/biFebriansyah/goback/src/database/gorm/models"
	help "github.com/biFebriansyah/goback/src/helpers"
	"github.com/biFebriansyah/goback/src/interfaces"
)

type cart_service struct {
	re interfaces.CartRepo
}

func NewService(rep interfaces.CartRepo) *cart_service {
	return &cart_service{rep}
}

func (r *cart_service) GetByUserId(id string) (*help.Response, error) {

	userId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	data, err := r.re.FindByUserId(userId)
	if err != nil {
		return nil, err
	}

	result := help.New(data, 200, false)
	return result, nil
}

func (r *cart_service) Save(data *models.Cart) (*help.Response, error) {

	data, err := r.re.Save(data)
	if err != nil {
		return nil, err
	}

	result := help.New(data, 200, false)
	return result, nil
}
