package mypackage

import "fmt"

// CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}
type carPrivate struct {
	brand string
	year  int
}

func PrintMessage() {
	fmt.Println("Hola")
}

func privateFunction() {
	fmt.Println("Private Message")
}
