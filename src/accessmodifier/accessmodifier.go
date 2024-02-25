package main

import (
	"fmt"
	"learn-go/src/accessmodifier/mypackage"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	mypackage.PrintMessage()
}
