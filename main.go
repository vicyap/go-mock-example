package main

import (
	"fmt"
)

// main

func Two(c CalculatorInterface) int {
    return c.One() + c.One()
}

func main() {
    c := Calculator{}
	fmt.Println(Two(c))
}
