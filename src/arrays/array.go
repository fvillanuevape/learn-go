package main

import (
	"fmt"
)

func main() {

	// Arrays - Son inmutables
	var array [4]int
	array[0] = 13
	array[1] = 19
	array[2] = 23
	array[3] = 12
	fmt.Println(array, len(array), cap(array))

	// Slice - Son Mutables
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice, len(slice), cap(slice))

	//Metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Apppend
	slice = append(slice, 7)
	fmt.Println(slice)

	// Apend nueva lista
	newSlice := []int{11, 12, 13}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	fmt.Println("Slices Function: ", createSlices(10))

}
func createSlices(limit int) (sl []int){
	var slices []int

	for index := 0; index < limit; index++ {
		slices = append(slices, index)
	}
	return slices
}
