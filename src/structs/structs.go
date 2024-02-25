/**
Create Structs with Go
*/
package main

import (
	"fmt"
)

type Car struct{
	brand string
	year int
	price fl
}

// Struc Vertex
type Vertex struct{
	X int
	Y int
}

// Struct Person

type Person struct {
	name                 string
	lastName             string
	personIdentification PersonIdentification
	age                  int
}

// Struc Person Identification
type PersonIdentification struct {
	documentType       string
	documentCountry    string
	documentIdentifier string
	documentExpedition string
}
type pc struct{
	ram int
	brand string
	disk int
}

func (myPc pc) String() string{
	return fmt.Sprintf("Tengo %d GM RAM, %d GB Disco y es una %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {

	// Test my Struct
	v:= Vertex{1,2}
	v.X=10
	fmt.Println("X: ", v.X)

	// Using Pointer in Struct
	p:= &v
	p.X=200
	fmt.Println("p.X: ", p.X)


	personOne := Person{
		name:     "Fi",
		lastName: "Lao",
		personIdentification: PersonIdentification{
			documentType:       "PASSPORT",
			documentCountry:    "ES",
			documentIdentifier: "94275252",
			documentExpedition: "24/09/2016",
		},
		age: 23,
	}
	personTwo := Person{
		name:     "Jhon",
		lastName: "Lao",
		personIdentification: PersonIdentification{
			documentType:       "PASSPORT",
			documentCountry:    "ES",
			documentIdentifier: "252633622",
			documentExpedition: "24/12/2015",
		},
		age: 24,
	}
	myPc := pc{ram:16, brand: "MSI", disk: 250}
	fmt.Println(personOne, personTwo)
	fmt.Println(myPc)
}
