package usecase

import (
	"errors"
	"testing"

	"github.com/rafaelcorrea26/Desafio-Clean-Architecture-Listagem-Orders/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type OrderRepositoryMock struct {
	mock.Mock
}

func (m *OrderRepositoryMock) Save(order *entity.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *OrderRepositoryMock) FindAll() ([]entity.Order, error) {
	args := m.Called()
	return args.Get(0).([]entity.Order), args.Error(1)
}

func TestGivenSavedOrders_WhenExecute_ThenShouldReturnOrdersDTO(t *testing.T) {
	repository := &OrderRepositoryMock{}
	repository.On("FindAll").Return([]entity.Order{
		{ID: "123", Price: 10.0, Tax: 2.0, FinalPrice: 12.0},
	}, nil)

	output, err := NewListOrdersUseCase(repository).Execute()

	assert.Nil(t, err)
	assert.Equal(t, []OrderOutputDTO{{ID: "123", Price: 10.0, Tax: 2.0, FinalPrice: 12.0}}, output)
	repository.AssertExpectations(t)
}

func TestGivenARepositoryError_WhenExecute_ThenShouldReturnTheError(t *testing.T) {
	repository := &OrderRepositoryMock{}
	repository.On("FindAll").Return([]entity.Order{}, errors.New("database error"))

	output, err := NewListOrdersUseCase(repository).Execute()

	assert.EqualError(t, err, "database error")
	assert.Nil(t, output)
}
