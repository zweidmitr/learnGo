package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	min = 1
	max = 5
)

func main() {
	rand.Seed(time.Now().UnixNano())

	value := rand.Intn(max-min) + min

	/* if example */
	fmt.Println("if example")
	if value == 1 {
		fmt.Println("Number is one")
	} else if value == 2 || value == 3 {
		fmt.Println("Number is two or three")
	} else if value == getFour() {
		fmt.Println("Number is four")
	} else {
		fmt.Println("Default case is shown")
	}
	fmt.Println("==============")
	fmt.Println("base switch ")

	/* base switch */
	switch value {
	case 1:
		fmt.Println("Number is one")
	case 2, 3:
		fmt.Println("Number is two or three")
	case getFour():
		fmt.Println("Number is four")
	default:
		fmt.Println("Default case is shown")
	}

	fmt.Println("==============")
	fmt.Println("switch with local variable definition")
	/* switch with local variable definition */
	switch num := rand.Intn(max-min) + min; num {
	case 1:
		fmt.Println("Number is one")
	case 2, 3:
		fmt.Println("Number is two or three")
	case getFour():
		fmt.Println("Number is four")
		fallthrough
	case 10:
		fmt.Println("Strange things happens")
	case 100:
		fmt.Println("ups")

	default:
		fmt.Println("Default case is shown")
	}

	fmt.Println("==============")
	fmt.Println("switch without condition")
	/* switch without condition */
	switch {
	case value > 2:
		fmt.Printf("Value %d greater than 2 \n", value)
	case value < 2:
		fmt.Printf("Value %d less than 2 \n", value)
	default:
		fmt.Println("Value equals 2")

	}

}

func getFour() int {
	fmt.Println("getFour called")
	return 4
}
