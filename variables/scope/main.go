package main

import "fmt"

var a = "variable"

func mas(x int) int {
	return 2 ^ x
}

func main() {
	b := 17

	mas := mas(8)
	fmt.Printf("%v \t %v \t %v \n", a, b, mas)
	foo()
}

func foo() {
	b := 54

	mas := mas(16)
	fmt.Printf("%v \t %v \t %v \n", a, b, mas)
}
