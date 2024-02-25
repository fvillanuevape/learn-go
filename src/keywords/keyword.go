package main

import "fmt"

func main() {

	// Defer, primero Imprime Second y después First
	defer fmt.Println("Hello First")
	fmt.Println("Hello Second")

	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("It´s two")
			continue
		}
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
