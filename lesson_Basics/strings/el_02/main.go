package main

import "fmt"

func main() {
	//one()
	two()
}

func one() {
	// Таким способом можно обойти только строки, состоящие из ASCII символов.
	str := "World"
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}
}
func two() {
	// Вывод некорректен.
	str := "Возможности"
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}
}
