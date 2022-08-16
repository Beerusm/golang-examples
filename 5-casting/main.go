package main

import (
	"fmt"
	"strconv"
	_ "time"
)

// Casting or conversion
// There is no implicit caste in go
func main() {

	var num1 int8 = 100

	//var num2 int16 = num1 // NOT OK

	// var num2 int16 = (int16)num1 // NOT OK

	var num2 int16 = int16(num1) // OK
	fmt.Println(num1, num2)

	var num3 float32 = 12.14

	var num4 int = int(num3)

	fmt.Println(num3, num4)

	var num5 float32 = 123.456

	var num6 float64 = float64(num5)

	fmt.Println(num5, num6)

	// type conversion

	var num7 uint32 = 19000
	var str string = string(num7) // NOT "19000"
	fmt.Println(str)

	str = fmt.Sprint(num7) // converts to "19000"
	fmt.Println(str)

	str1 := strconv.Itoa(int(num7)) // converts to "19000"
	fmt.Println(str1)

	// str to int
	str2 := '䨸'
	fmt.Println(int(str2))

	str3 := "17000"

	num8, _ := strconv.Atoi(str3) // _ blank identifier.
	fmt.Println(num8)

	// swap
	a, b, c := 10, 20, 30
	fmt.Println("a:", a, "b:", b, "c:", c)
	//a, b, c = c, a, b
	//fmt.Println("a:", a, "b:", b, "c:", c)
	t := a
	t2 := b
	a = c
	b = t
	c = t2
	fmt.Println("Beeru a:", a, "b:", b, "c:", c)

	a, b, c = 10, 20, 30
	a, b, c = swap3(a, b, c)
	fmt.Println("func swap a:", a, "b:", b, "c:", c)
	//a,b=b,a

}

func swap3(a int, b int, c int) (int, int, int) {
	return c, a, b
}

// a,b,c
// swap using manual way
