package interfaces

import (
	"github.com/biFebriansyah/goback/src/database/gorm/models"
	help "github.com/biFebriansyah/goback/src/helpers"
)

type UserRepo interface {
	FindAll() (*models.Users, error)
	FindByUsername(username string) (*models.User, error)
	Add(*models.User) (*models.User, error)
}

type UserService interface {
	FindAll() (*help.Response, error)
	FindByUsername(username string) (*help.Response, error)
	Save(*models.User) (*help.Response, error)
}
