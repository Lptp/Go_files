package main

import "fmt"

func main() {
	fmt.Println("Memory Location:")
	fmt.Println("Decimal \t Hexadecimal \t Binary")
	for i := 0; i < 257; i++ {
		a := i
		fmt.Println("Memory Location:")
		fmt.Println("Decimal \t Hexadecimal \t Binary")
		fmt.Printf("%d \t %#x \t %b \n", &a, &a, &a)
	}
}
