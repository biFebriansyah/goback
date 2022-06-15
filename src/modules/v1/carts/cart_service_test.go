package carts

import (
	"testing"

	"github.com/biFebriansyah/goback/src/database/gorm/models"
	"github.com/biFebriansyah/goback/src/modules/v1/carts/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var modelMock = models.Cart{
	ID: 1,
}

func TestGetByUserId(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = cart_service{&repo}

	// prepare function
	repo.Mock.On("FindByUserId", 1).Return(&modelMock, nil)
	data, err := service.GetByUserId("1")
	var expectId uint = 1

	carts := data.Data.(*models.Cart)
	assert.Equal(t, expectId, carts.ID, "Expect id = 1")
	assert.Nil(t, err)

}
