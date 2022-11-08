package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// llamar a la funcion personalizada
	sum := sum(os.Args[1], os.Args[2])
	println(sum)

	// llamar a una funcion que retorna dos valores
	suma, multiplicacion := calc(os.Args[1], os.Args[2])
	println(suma, multiplicacion)

	// omitir el retorno dos variables
	resultadoUno, _ := calc(os.Args[1], os.Args[2])
	fmt.Println(resultadoUno)
}

/*
 * funcion personalizada
 * retorna un unico valor de tipo int
 */
func sum(numberOne string, numberTwo string) int {
	numeroUno, _ := strconv.Atoi(numberOne)
	numeroDos, _ := strconv.Atoi(numberTwo)
	return numeroUno + numeroDos
}

/*
 * Funcion personalizada
 * Retorna dos valores (sum, mult)
 */
func calc(numberOne string, numberTwo string) (sum int, mult int) {
	numeroUno, _ := strconv.Atoi(numberOne)
	numeroDos, _ := strconv.Atoi(numberTwo)
	sum = numeroUno + numeroDos
	mult = numeroUno * numeroDos
	return sum, mult
}
