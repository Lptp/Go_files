package main

import (
	"fmt"
	"math"
)

var a = "variable"

func mas(x float64) float64 {
	return math.Pow(2, x)
}

func main() {
	b := 17.0

	mas := mas(b) //EJEMPLO DE VARIABLE SHADOWING
	fmt.Printf("%v \t %v \t %v \n", a, b, mas)
	foo()
}

func foo() {
	b := 54.0

	elev := mas(b) //EJEMPLO DE  NO VARIABLE SHADOWING
	fmt.Printf("%v \t %v \t %v \n", a, b, elev)
}
