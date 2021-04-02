package database

import (
	"github.com/stretchr/testify/mock"
)
type Store interface {
	Get(ID int)(int,error)

}

type MockSore struct {
	mock.Mock
}

func (m *MockSore)Get(ID int)(int ,error){
	returnVals := m.Called(ID)

	return returnVals.Get(0).(int),returnVals.Error(1)

}