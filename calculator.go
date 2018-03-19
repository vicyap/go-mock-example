package main

import (
	"fmt"
)

type CalculatorInterface interface {
    One() int
}

type Calculator struct {}

func (c Calculator) One() int {
    fmt.Println("The Real Add")
    return 1
}
