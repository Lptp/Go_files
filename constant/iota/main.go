package main

import "fmt"

const (
	A = iota //0
	B        //1
	C        //2
)

const (
	D = iota
	E
	F
)

const (
	_ = iota //Dropping the zero
	H = iota * 10
	I = iota * 10
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(H)
	fmt.Println(I)
}
