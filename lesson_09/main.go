package main

import (
	"fmt"
)

func main() {
	//Указатели это тип данных которые в качестве значения хранят фдрес ячейки
	//памяти значения, либо другого указателя (может быть nil)

	// default value
	var intPointer *int
	fmt.Printf("%T %#v \n", intPointer, intPointer)

	// nil pointer panic
	//fmt.Printf("%T %#v %#v \n", intPointer, intPointer, *intPointer)

	// Получение not nil указателей
	// variable
	var a int64 = 7
	fmt.Printf("%T %#v \n", a, a)
	// get variable pointer
	var pointerA *int64 = &a
	fmt.Printf("%T | %#v | %#v \n", pointerA, pointerA, *pointerA)

	//get pointer via new keyword
	var newPointer = new(float32)
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)
	*newPointer = 3
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)

	fmt.Println("\n ============== pointers usage ==============")
	//pointers usage
	//side effects
	num := 3
	square(num) // копируется значение и передаем в функцию
	fmt.Println(num)
	fmt.Println("==============")
	squarePointer(&num)
	fmt.Println(num)
	fmt.Println("==============")
	fmt.Println("==============")

	// empty value flag
	var wallet1 *int
	fmt.Println(hasWallet(wallet1)) // false

	wallet2 := 0
	fmt.Println(hasWallet(&wallet2)) // true

	wallet3 := 100
	fmt.Println(hasWallet(&wallet3)) // true
}

func square(num int) {
	num *= num
}

func squarePointer(num *int) {
	*num = *num * *num
}

func hasWallet(money *int) bool {
	return money != nil
}
