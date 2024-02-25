package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// llamar a la funcion personalizada
	sum := sum(os.Args[1], os.Args[2])
	fmt.Println(sum)

	// llamar a una funcion que retorna dos valores
	suma, multiplicacion := calc(os.Args[1], os.Args[2])
	fmt.Println(suma, multiplicacion)

	// omitir el retorno dos variables usando "_"
	resultadoUno, _ := calc(os.Args[1], os.Args[2])
	fmt.Println(resultadoUno)

	// call normal function
	normalFunction("This is a new message")


	//Naked
	fmt.Println(split(9))
}

// Funcion normal
func normalFunction(message string) {
	fmt.Println(message)
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

// funciones con 3 parametros
func parametersFunction(name, phoneNumber, email string) (userInformation string) {
	userInfo := name + phoneNumber + email
	return userInfo
}


func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
