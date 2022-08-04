package main

import "fmt"

func main() {
	first, second := 1, 2
	//Greet()
	//PersonGreet("Dima")
	//FioGreet("John", "Smith")
	//sum := Sum(first, second)
	//fmt.Println(sum)
	//sum, multiply := SumAndMultiply(first, second)
	//fmt.Println(sum, multiply)
	_, multiply64 := namedSumAndMultiply(first, second)
	fmt.Println(multiply64)
}

func namedSumAndMultiply(first, second int) (sum int64, multiply int64) {
	sum = int64(first + second)
	multiply = int64(first) * int64(second)
	return sum, multiply
}

func SumAndMultiply(first, second int) (int, int) {
	return first + second, first * second
}

func Sum(first, second int) int {
	return first + second
}

func FioGreet(name, surname string) {
	fmt.Printf("Hi %s %s\n", name, surname)
}

func Greet() {
	fmt.Println("hello guys")
}

func PersonGreet(name string) {
	fmt.Printf("Zdarova %s\n", name)
}
