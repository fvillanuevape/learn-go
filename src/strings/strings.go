package main

import (
	"fmt"
	"strings"
)

func main() {

	// Using strings package
	name := "Fid#l"
	newName := strings.ReplaceAll(name, "#", "e")
	toLower:= strings.ToLower(newName)
	fmt.Println(newName)
	fmt.Println(toLower)
	

}
