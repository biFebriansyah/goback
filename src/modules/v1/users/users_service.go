package users

import (
	"github.com/biFebriansyah/goback/src/database/gorm/models"
	help "github.com/biFebriansyah/goback/src/helpers"
	"github.com/biFebriansyah/goback/src/interfaces"
)

type users_service struct {
	rep interfaces.UserRepo
}

func NewService(rep interfaces.UserRepo) *users_service {
	return &users_service{rep}
}

func (re *users_service) FindAll() (*help.Response, error) {
	data, err := re.rep.FindAll()
	if err != nil {
		return nil, err
	}

	respone := help.New(data, 200, false)
	return respone, nil
}

func (re *users_service) FindByUsername(username string) (*help.Response, error) {
	data, err := re.rep.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	respone := help.New(data, 200, false)
	return respone, nil
}
func (re *users_service) Save(usr *models.User) (*help.Response, error) {
	if check := re.rep.UserExsist(usr.Username); check {
		return help.New("username already exsist", 400, true), nil
	}

	hassPssword, err := help.HashPasword(usr.Password)
	if err != nil {
		return nil, err
	}

	usr.Password = hassPssword
	data, err := re.rep.Add(usr)
	if err != nil {
		return nil, err
	}

	respone := help.New(data, 200, false)
	return respone, nil
}
