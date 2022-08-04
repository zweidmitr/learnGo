package main

import "fmt"

func main() {
	//first, second := 1, 2
	//Greet()
	//PersonGreet("Dima")
	//FioGreet("John", "Smith")
	//sum := Sum(first, second)
	//fmt.Println(sum)
	//sum, multiply := SumAndMultiply(first, second)
	//fmt.Println(sum, multiply)
	//_, multiply64 := namedSumAndMultiply(first, second)
	//fmt.Println(multiply64)

	/* Func as values */
	//var multiplier func(x, y int) int
	//multiplier = func(x, y int) int { return x * y }
	//fmt.Println(multiplier(first, second))
	//
	//divider := func(x, y int) int { return x / y }
	//fmt.Println(divider(second, first))

	/* Pass func as argument */
	//sumFunc := func(x, y int) int { return x + y }
	//subtractFunc := func(x, y int) int { return x - y }
	//
	//fmt.Println(calculate(first, second, sumFunc))
	//fmt.Println(calculate(second, first, subtractFunc))

	/* Return new function */
	//divideBy2 := createDivider(2)
	//divideBy10 := createDivider(10)
	//
	//fmt.Println(divideBy2(100))
	//fmt.Println(divideBy10(100))

	/* НЕЙМИНГ как модификатор доступа к функциям */
	//fmt.getField()
	//fmt.Println()

	/* Closures */
	dollar := 30
	getDollarValue := func() int { return dollar }
	fmt.Println(getDollarValue())

	dollar = 70
	fmt.Println(getDollarValue())
}

func createDivider(divider int /* 2 */) func(x int) int {
	dividerFunc := func(x int) int {
		return x / divider /* 2 */
	}
	return dividerFunc
}

func calculate(x, y int, action func(x, y int) int) int {
	return action(x, y)
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
