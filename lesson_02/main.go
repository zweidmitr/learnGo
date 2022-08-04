package main

import (
	"fmt"
	"unsafe"
)

var name string = "dMitr"

const nameLanguage = "Go"

func main() {
	fmt.Println("Hello Go...")
	name := "Zwei"

	var hello string = "hell0"
	fmt.Println(hello)
	hello = "hello"
	fmt.Printf("Type: %T  Value: %v\n", hello, hello)

	var ourBool bool
	fmt.Println(ourBool)
	ourBool = true
	fmt.Printf("Type: %T  Value: %v\n", ourBool, ourBool)

	ourBoolTwo := true
	fmt.Println(ourBoolTwo)

	fmt.Printf("Type: %T  Value: %v\n", name, name)
	fmt.Printf("Type: %T  Value: %v\n", nameLanguage, nameLanguage)

	one := "one"
	two := fmt.Sprintf(" dva/odin = two/%s", one)
	_ = two
	fmt.Println(two)

	fmt.Println("<=============>")
	var num1 int64 = 15
	var num2 int = 15
	fmt.Println(num1 + int64(num2))
	fmt.Println("<=============>")
	fmt.Println(unsafe.Sizeof(uint8(1)))
	fmt.Println(unsafe.Sizeof(int32(1)))
	fmt.Println(unsafe.Sizeof("hello"))
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println("<=============>")
	fmt.Println(unsafe.Sizeof(uint8(1)))
	var num3 uint64 = 1
	fmt.Println(unsafe.Sizeof(num3))
}
