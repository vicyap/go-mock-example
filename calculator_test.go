package main

import (
	"fmt"
    "testing"

	"github.com/stretchr/testify/mock"
)

type CalculatorMock struct {
	mock.Mock
}

func (m *CalculatorMock) One() int {
    fmt.Println("The Mock One")
	args := m.Called()
	return args.Int(0)
}

func TestOne(t *testing.T) {
    m := CalculatorMock{}
    m.On("One").Return(1)

    Two(&m)
    m.AssertNumberOfCalls(t, "One", 2)
}
