package main

import (
	"fmt"
)

func main() {
	argumentos("Fidel", "Villanueva")

	returnValue := retornaValor("Mi mensaje")
	fmt.Println("Retorna Valor: ", returnValue)

	//call function when return two value
	valueOne, valueTwo := returnTwoValue("Fidel", "Villanueva")
	fmt.Println("Value One: ", valueOne)
	fmt.Println("Value Two: ", valueTwo)

	//call function when return two value but use Only Value
	value, _ := returnTwoValue("Fidel", "Villanueva")
	fmt.Println("Only value: ", value)
}

// Funcion con argumentos
func argumentos(name string, lastName string) {
	fmt.Println(name, lastName)
}

// funcion que retorna un valor
func retornaValor(message string) string {
	return message
}

// funcion que retorna mas de un valor 912226035

func returnTwoValue(name, lastName string) (valueName, valueLastName string) {
	return name, lastName
}
