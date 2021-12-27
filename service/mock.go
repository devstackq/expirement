package service

import (
	"inert/model"
	"log"

	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) AddUser(in model.Inert) (int, error) {
	log.Print("call mock add user", in)

	args := m.Called(in)
	return args.Int(0), args.Error(1)
}
