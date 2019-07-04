package main

import "fmt"

var metros float64

const metrosTofeets = 3.337

func main() {
	fmt.Println("Cuántos metros mides")
	fmt.Scan(&metros)
	feets := metros * metrosTofeets
	fmt.Println("Mides:", feets, "feets")
}
