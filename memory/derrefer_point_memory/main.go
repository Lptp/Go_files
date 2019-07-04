package main

import "fmt"

func zero(z int) {
	z = 0
}

func zeroMem(y *int) {
	*y = 0
}

func main() {
	x := 105
	zero(x)
	fmt.Println("zero(x)")
	fmt.Println("Valor de x:", x)
	fmt.Println("Address space x:", &x)
	zeroMem(&x)
	fmt.Println("zeroMem(&x)")
	fmt.Println("Valor de x:", x)
	fmt.Println("Address space x:", &x)
}
