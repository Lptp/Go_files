package main

import "fmt"

func main() {
	nombre := ""
	fmt.Println("¿Qué color buscas?")
	fmt.Scan(&nombre)
	switch nombre {
	case "Azul":
		fmt.Printf("%s sí hay \n", nombre)
	case "Rojo":
		fmt.Printf("%s sí hay \n", nombre)
	case "Amarillo":
		fmt.Printf("%s sí hay \n", nombre)
	case "Negro":
		fmt.Printf("%s sí hay \n", nombre)
	default:
		fmt.Printf("%s no hay \n", nombre)
	}
}
