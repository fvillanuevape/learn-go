package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	x := 28

	//instruccion condicional
	if x%2 == 0 {
		fmt.Println(x, "is even")
	}

	/* instrucciones condicional compuestos
	num solo existe en los bloques condicional else
	*/
	if num := givemeanumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has only one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// Convertir texto a numero
	value, err := strconv.Atoi("45")
	if err != nil {
		log.Fatal(err,value)
	}
	if day:= getWeekDay(); day=="Thursday"{
		fmt.Println("It is Thursday")
	}

}
func givemeanumber() int {
	return -1
}

func getWeekDay() string{
	today:= time.Now().Weekday()
	return today.String()
}
