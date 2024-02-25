package main

import "fmt"

var globalScope int = 30

func main() {

	var name string = "This is name"
	fmt.Println(name)

}

func getGlobalVariable() int{

	return globalScope
}

func getLocalVariable() int{
	var localScope int = 60
	return localScope
}
