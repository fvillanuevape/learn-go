/*
Creación de un paquete reutilizado por main.go
*/
package calculator

var logMessage = "[LOG]"

// Version of the calculator
var Version = "1.0"

func internalSum(number int) int {
	return number - 1
}

// Sum two integer numbers
func Sum(numberOne int, numberTwo int) int {
	return numberOne + numberTwo
}
