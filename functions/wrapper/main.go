package main

import (
	"fmt"
	"math"
)

func wrapper() func() float64 {
	x := 8.0

	return func() float64 {
		x++
		return math.Pow(2, x)
	}

}

func main() {
	elevar := wrapper()
	fmt.Println(elevar())
	fmt.Println(elevar())
}
