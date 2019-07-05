package main

import "fmt"

func main() {
	nombre := ""
	fmt.Println("¿Qué color buscas?")
	fmt.Scan(&nombre)
	switch nombre {
	case "azul":
		fmt.Printf("%s sí hay. \n", nombre)
		fallthrough
	case "restoAz":
		fmt.Println("También hay rojo, amarillo o negro.")
	case "rojo":
		fmt.Printf("%s sí hay. \n", nombre)
		fallthrough
	case "restoRo":
		fmt.Println("También hay azul, amarillo o negro.")
	case "amarillo":
		fmt.Printf("%s sí hay. \n", nombre)
		fallthrough
	case "restoAm":
		fmt.Println("También hay azul, rojo o negro.")
	case "negro":
		fmt.Printf("%s sí hay. \n", nombre)
		fallthrough
	case "restoNe":
		fmt.Println("También hay azul, rojo o amarillo.")
	default:
		fmt.Printf("%s no hay. Sí hay azul, rojo, amarillo o negro. \n", nombre)
	}
}
