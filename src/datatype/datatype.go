package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// Numeros enteros
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807
	fmt.Println(integer8, integer16, integer32, integer64)
	var conversion int32 = int32(integer8) + integer32
	fmt.Println(conversion)

	// runes
	rune := 'G'
	fmt.Println("Rune : ",rune)

	// Numeros enteros positivos, empiezan desde 0
	var integeruint uint = 20
	var integeruint8 uint8 =23

	fmt.Println(integeruint)
	fmt.Println(integeruint8)

	// valores flotantes o decimales
	var float32 float32 = 2147483647
	var float64 float64 = 9223372036854775807
	fmt.Println(float32, float64)
	fmt.Println(math.MaxFloat32, math.MaxFloat64)

	/* conversiones
	 * el _ indica que no se usará dicho valor, ya que Atoi retorna dos valores
	 */
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)

	// Numeros complejos
	var complejo complex128 = 10 + 8i
	fmt.Println("Complejo: ", complejo)

}
