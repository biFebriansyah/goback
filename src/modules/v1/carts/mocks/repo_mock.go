package mocks

import (
	"github.com/biFebriansyah/goback/src/database/gorm/models"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	Mock mock.Mock
}

func (pr *RepoMock) FindByUserId(id int) (*models.Cart, error) {
	args := pr.Mock.Called(id)
	return args.Get(0).(*models.Cart), args.Error(1)
}

func (pr *RepoMock) Save(data *models.Cart) (*models.Cart, error) {
	args := pr.Mock.Called(data)
	return args.Get(0).(*models.Cart), args.Error(1)
}
