package main

import (
	"fmt"
	"strings"
)

func main() {
	//one()
	two()
}

func one() {
	// проверяет наличие подстроки в строке
	fmt.Println(strings.Contains("Приветик", "и"))

	// разбивает строку по Юникод символам или по переданному разделителю
	fmt.Println(strings.Split("Удачи", ""))

	// склеивает строки из слайса с разделителем
	fmt.Println(strings.Join([]string{"Привет", "человек!"}, " "))

	// обрезает символы из второго аргумента в строке
	fmt.Println(strings.Trim(" тамТарам !", " "))
}

func two() {

	sb := &strings.Builder{}

	sb.WriteString("Сегодня")
	sb.WriteString(" ")
	sb.WriteString("хороший")
	sb.WriteString(" ")
	sb.WriteString("день")

	fmt.Println(sb.String())
}
