package main

import "fmt"

func main() {
	x := 28

	//instruccion if
	if x%2 == 0 {
		fmt.Println(x, "is even")
	}

	/* instrucciones if compuestos
	num solo existe en los bloques if else
	*/
	if num := givemeanumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has only one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
func givemeanumber() int {
	return -1
}
