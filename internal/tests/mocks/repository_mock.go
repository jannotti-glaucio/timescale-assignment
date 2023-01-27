package mocks

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type mockedRepository struct {
	mock.Mock
}

func NewMockedRepository() mockedRepository {
	return mockedRepository{}
}

func (r mockedRepository) OpenConnection() error {
	ret := r.Mock.Called()
	return ret.Error(0)
}

func (r mockedRepository) RunQuery(hostname string, startDate time.Time, endDate time.Time) (*float32, *float32, error) {
	ret := r.Mock.Called(hostname, startDate, endDate)
	return ret.Get(0).(*float32), ret.Get(1).(*float32), ret.Error(2)
}

func (r mockedRepository) CloseConnection() error {
	ret := r.Mock.Called()
	return ret.Error(0)
}
