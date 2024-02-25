package main

import "fmt"

func main() {

	// constantes
	const pi float32 = 3.14
	const piTwo = 3.15

	fmt.Println("Pi One: ", pi)
	fmt.Println("Pi Two: ", piTwo)

	//Variables

	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	name := "Fidel"
	fmt.Println(name)

	// multiples variables
	var (
		email       string = "email@learn.com"
		phoneNumber string = "943856924"
	)
	fmt.Println(email)
	fmt.Println(phoneNumber)

	// default values for variables

	var intValue int
	var floatValue float32
	var stringValue string
	var boolValue bool

	fmt.Println(intValue, floatValue, stringValue, boolValue)

	// declare variables using the type alias
	var temp Celsius
	var pid IDnum

	fmt.Println(temp)


}
