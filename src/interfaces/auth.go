package interfaces

import (
	"github.com/biFebriansyah/goback/src/database/gorm/models"
	help "github.com/biFebriansyah/goback/src/helpers"
)

type AuthService interface {
	Login(body models.User) (*help.Response, error)
}
