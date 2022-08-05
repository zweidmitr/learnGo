package main

import "fmt"

func main() {
	age := 15

	//var firstVariable bool
	//fmt.Println(firstVariable)
	//
	//secondVariable := true
	//fmt.Println(secondVariable)

	/* Basic if */
	if age < 18 {
		//fmt.Println("You are too young")
	}

	/* Short syntax */
	if isChild := isChildren(age); isChild == true {
		//fmt.Println("You are very young (short)")
		fmt.Println(isChild)
	}

	/* If ... else */
	age = 20
	if age < 18 {
		fmt.Println("You are too young")
	} else {
		//fmt.Println("You are an adult")
	}

	/* && */
	if age >= 7 && age <= 18 {
		//fmt.Println("You are in school")
	}

	/* || */
	age = 40
	if age == 14 || age == 20 || age == 40 {
		//fmt.Println("You have to get new passport")
	}

	/* ! */
	if !isChildren(age) {
		fmt.Println("You are an adult")
	}
}

func isChildren(age int) bool {
	return age < 18
}
