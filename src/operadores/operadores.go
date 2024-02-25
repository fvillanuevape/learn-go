package main

import "fmt"

func main() {

	// operadores
	x := 10
	y := 5

	//multiplicacion
	resultado := x * y
	fmt.Println("Multiplicacion: ", resultado)

	// division

	resultado = x / y
	fmt.Println("Division: ", resultado)

	// resta
	resultado = x - y
	fmt.Println("Resta: ", resultado)

	// modulo
	resultado = x % y
	fmt.Println("Modulo ", resultado)

	// Incremental
	x++
	fmt.Println("Incremental: ", x)

	// Decremental
	x--
	fmt.Println("Decremental: ", x)
	
}
