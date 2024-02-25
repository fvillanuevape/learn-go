package main

import "fmt"

func main() {

	// Define maps
	maps := make(map[string]int)
	maps["Jose"] = 23
	maps["Fidel"] = 26
	maps["Carlos"] = 29

	fmt.Print(maps, "\n")

	// Recorrer

	for key, value := range maps {
		fmt.Println(key, value)
	}

	// Find value

	value, error := maps["Fidel"]
	fmt.Println(value, error)

}
