package main

import (
	"fmt"
	"strings"
)

func main() {

	nameSlice := []string{"Fidel", "Manuel", "Carlos"}

	// recorrer Arrays

	for index, value := range nameSlice {
		fmt.Println(index, value)
	}

	// Get only value

	for _, value := range nameSlice {
		fmt.Println(value)
	}


	// Gte only index
	for index := range nameSlice {
		fmt.Println(index)
	}

	// Palindromo

	response := palindromo("Ama")
	fmt.Println(response)
}

func palindromo(text string) bool {
	var textReverse string
	var esPalindromo bool
	newText := strings.ToLower(text)

	for i := len(newText) - 1; i >= 0; i-- {
		textReverse += string(newText[i])
	}
	if textReverse == newText {
		esPalindromo = true
	}
	return esPalindromo
}
