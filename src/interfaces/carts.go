package interfaces

import (
	"github.com/biFebriansyah/goback/src/database/gorm/models"
	help "github.com/biFebriansyah/goback/src/helpers"
)

type CartRepo interface {
	FindByUserId(id int) (*models.Cart, error)
	Save(data *models.Cart) (*models.Cart, error)
}

type CartService interface {
	GetByUserId(id string) (*help.Response, error)
	Save(data *models.Cart) (*help.Response, error)
}
