package main

import "fmt"

func main() {
	for words := 97; words < 123; words++ {
		letras := words
		var memLetras *int = &letras
		fmt.Printf("Letra:%q \tPoints to: %v Con valor: %v \n", letras, memLetras, *memLetras)
	}
}
