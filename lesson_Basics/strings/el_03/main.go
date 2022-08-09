package main

import "fmt"

/*
rune — это алиас к int32. Как и байты, руны были созданы для отличия от встроенного типа данных.
Каждая руна представляет собой код символа стандарта Юникод.
Строка свободно преобразуется в []byte и []rune,
но эти 2 типа данных не конвертируются между собой напрямую:
*/
func main() {
	//one()
	two()
}

func one() {
	text := []rune("Привет")

	for i := 0; i < len(text); i++ {
		fmt.Println(text[i], string(text[i]))
	}
}

func two() {
	text := []rune("Руководство для всего")
	for _, ch := range text {
		fmt.Println(ch, string(ch))
	}
}
