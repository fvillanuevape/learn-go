package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

// Funcion Receptora

func (myPc pc) ping() {
	fmt.Println(myPc.brand)
}
func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

func main() {

	// & operator returns the address of a variable/function
	// * operator returns data at an address

	a := 50
	b := &a
	*b = 200

	fmt.Println("a: ", a)
	fmt.Println("b: ", *b)

	// Se modifica el valor en la memoria
	*b = 100
	fmt.Println(a)
	fmt.Println(*b)

	myPc := pc{
		ram:   16,
		disk:  300,
		brand: "msi",
	}
	fmt.Println(myPc)
	myPc.ping()

	myPc.duplicateRam()
	fmt.Println(myPc)

	myPc.duplicateRam()
	fmt.Println(myPc)

	// New
	prt := new(int)
	*prt = 3
	fmt.Println("PTR: ", *prt)

}
