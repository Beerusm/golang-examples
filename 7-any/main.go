package main

import (
	"fmt"
	"reflect"
	//"github.com/golang/protobuf/ptypes/any"
)

// any is nothing but interface{}
// rune is nothing but int32
// byte is nothing but uint8

// empty interface
// type any = interface{}
func main() {

	//var val interface{} // this is called as empty interface. This is capable of holding any kind of value
	var val any // interface{}

	val = 100
	fmt.Println("Value:", val, "Type:", reflect.TypeOf(val))
	// type casting --> v = T(v)
	// type assertion-> v = v.(T)
	var num1 int = val.(int) // here type casting does not work. There is a concept i.e type assertion
	fmt.Println("Value of num1:", num1)

	val = false
	fmt.Println("Value:", val, "Type:", reflect.TypeOf(val))
	var bool1 bool = val.(bool)
	fmt.Println("Value of bool1:", bool1)

	val = 100.100
	fmt.Println("Value:", val, "Type:", reflect.TypeOf(val))

	val = "Hello World"
	fmt.Println("Value:", val, "Type:", reflect.TypeOf(val))

	val = 'A'
	fmt.Println("Value:", val, "Type:", reflect.TypeOf(val))

	val = byte(65) // byte is nothing but uint8
	fmt.Println("Value:", val, "Type:", reflect.TypeOf(val))

	//fmt.Println(val)

	var a any = 10
	var b int = 20
	var c float32 = 30

	d := a.(int) + b + int(c)
	e := a + b + c
	fmt.Println("d = a + b+ c", d)
}

// a,b,c
// a must be any
// b int
// c float32
//print (a+b+c)
