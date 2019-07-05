package main

import "fmt"

type contacto struct {
	saludos string
	nombre  string
}

func switchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("número")
	case string:
		fmt.Println("letras")
	case contacto:
		fmt.Println("contacto")
	default:
		fmt.Println("unknown")

	}
}

func main() {
	switchOnType(42)
	switchOnType("L0L")
	var t = contacto{"Hola qué tal,", "Paco"}
	switchOnType(t)
	switchOnType(t.saludos)
	switchOnType(t.nombre)
}
