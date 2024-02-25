package main

import "fmt"

func main() {

	// For condicional
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum of 1..100 is", sum)

	//For while
	// Aplica hasta que la condición se cumpla
	counter := 0
	for counter < 10 {
		fmt.Println(counter)

		// se incrementa el contador
		counter++
	}

	// For forever

	counterForEver := 0
	for {
		fmt.Println(counterForEver)
		counterForEver++
	}

}

// example loops in function

