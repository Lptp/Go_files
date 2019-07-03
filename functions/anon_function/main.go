package main

import (
	"fmt"
	"math"
)

var x = 4.0

func expon() float64 {
	return math.Pow(2, x)

}

func main() {
	x := 2.0

	expon2 := func() float64 {
		return math.Pow(2, x)
	}

	fmt.Println(expon())
	fmt.Println(expon2())
}
