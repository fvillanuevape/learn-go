package main

import (
	"fmt"
	"github.com/fvillanuevape/calculator"
)

func main() {
	fmt.Println("Hello World!")

	// variables
	var firstName string = "Name"
	var lastName string = "Last Name"

	// variables en una sola linea
	var email, address string = "email@gmail.com", "my address 345"

	// bloquess de variables
	var (
		documentNumer string = "123456789"
		age           int    = 23
	)

	/* otro metodo para declarar variable
	 * puede ser declarado solo dentro de una funcion
	 */
	phoneNumber := "123456789"

	fmt.Println(firstName, lastName, email, address, documentNumer, age, phoneNumber)

	// constantes
	const URLBASE = "https://apicontoso.com.pe/bank/v1"

	// declarar varias constantes
	const (
		StatusOK    = 200
		StatusError = 500
	)
	println(URLBASE, StatusOK, StatusError)

	// referencia a un paquete
	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println("Version: ", calculator.Version)

}
