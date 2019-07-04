package main

import "fmt"

func main() {
	for valor := 0; valor < 30; valor++ {
		if valor%2 == 0 {
			fmt.Println("Foo")
		} else {
			fmt.Println("Bar")
		}
	}
}
