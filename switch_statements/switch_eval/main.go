package main

import "fmt"

func main() {
	color := ""
	fmt.Println("¿Qué color buscas?")
	fmt.Scan(&color)
	switch {
	case color == "Azul", color == "azul":
		fmt.Printf("%s sí hay \n", color)
	case color == "Rojo", color == "rojo":
		fmt.Printf("%s sí hay \n", color)
	case color == "Amarillo", color == "amarillo":
		fmt.Printf("%s sí hay \n", color)
	case color == "Negro", color == "negro":
		fmt.Printf("%s sí hay \n", color)
	case color == "Verde", color == "verde", color == "Violeta", color == "violeta":
		fmt.Printf("%s no hay. Sí hay Azul, Rojo, Amarillo o Negro \n", color)
	default:
		fmt.Printf("%s no hay \n", color)
	}
}
