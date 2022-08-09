package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Строки в Go — это иммутабельные массивы байт.

	//one()
	//two()
	three()

}

func one() {
	s := "hey"
	fmt.Println(s[0], s[1], s[2])
	fmt.Println(string(s[0]), string(s[1]), string(s[2]))
}

func two() {
	s := "hello"
	bs := []byte(s)
	fmt.Println([]byte(s))
	fmt.Println(string(bs))
}
func three() {
	asciiCh := byte('Q')
	asciiChStr := string(asciiCh)

	fmt.Println(reflect.TypeOf(asciiCh), asciiCh)
	fmt.Println(reflect.TypeOf(asciiChStr), asciiChStr)
	fmt.Printf("Type: %T, Value: %#v", asciiChStr, asciiChStr)
}
