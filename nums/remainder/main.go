package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%d es par \n", i)
		} else {
			fmt.Printf("%d es impar \n", i)
		}
	}
}
