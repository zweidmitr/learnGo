package main

import "fmt"

func main() {
	//variadicFunctions()
	//convertToArrayPointer()
	//passToFunction()
	sliceWithNew()
}

func variadicFunctions() {
	showAllElements(1, 2, 3)
	showAllElements(1, 2, 3, 4, 5, 6, 7)

	firstSlice := []int{5, 6, 7, 8}
	secondSlice := []int{9, 3, 2, 1}
	showAllElements(firstSlice...) // (5,6,7,8)

	newSlice := append(firstSlice, secondSlice...) // (firstSlice, 9,3,2,1)
	fmt.Printf("type: %T value: %#v\n", newSlice, newSlice)
}

func showAllElements(values ...int) {
	for _, val := range values {
		fmt.Println("Value:", val)
	}
	fmt.Println()
}

func convertToArrayPointer() {
	initialSlice := []int{1, 2}
	fmt.Printf("type: %T value: %#v\n", initialSlice, initialSlice)
	fmt.Printf("length: %d capacity: %d\n\n", len(initialSlice), cap(initialSlice))

	intArray := (*[2]int)(initialSlice)
	fmt.Printf("type: %T value: %#v\n", intArray, intArray)
	fmt.Printf("length: %d capacity: %d\n\n", len(intArray), cap(intArray))
}

func passToFunction() {
	initialSlice := []int{1, 2}
	fmt.Printf("type: %T value: %#v\n", initialSlice, initialSlice)
	fmt.Printf("length: %d capacity: %d\n\n", len(initialSlice), cap(initialSlice))

	changeValue(initialSlice)
	fmt.Printf("type: %T value: %#v\n", initialSlice, initialSlice)
	fmt.Printf("length: %d capacity: %d\n\n", len(initialSlice), cap(initialSlice))
	fmt.Println("<--------------->")

	newSlice := append(initialSlice, 3)
	fmt.Printf("type: %T value: %#v\n", newSlice, newSlice)
	fmt.Printf("length: %d capacity: %d\n\n", len(newSlice), cap(newSlice))
	fmt.Println("<--------------->")

	newSlice2 := appendValue(newSlice)
	fmt.Printf("type: %T value: %#v\n", newSlice, newSlice)
	fmt.Printf("length: %d capacity: %d\n\n", len(newSlice), cap(newSlice))

	fmt.Printf("type: %T value: %#v\n", newSlice2, newSlice2)
	fmt.Printf("length: %d capacity: %d\n\n", len(newSlice2), cap(newSlice2))
}

func changeValue(slice []int) {
	slice[1] = 15
}

func appendValue(slice []int) []int {
	slice = append(slice, 4, 5)
	fmt.Printf("type: %T value: %#v\n", slice, slice)
	fmt.Printf("length: %d capacity: %d\n\n", len(slice), cap(slice))
	return slice
}

func sliceWithNew() {
	slicePointer := new([]int)

	fmt.Printf("type: %T value: %#v\n", slicePointer, slicePointer)
	fmt.Printf("length: %d capacity: %d\n\n", len(*slicePointer), cap(*slicePointer))

	newSlice3 := append(*slicePointer, 1)
	fmt.Printf("type: %T value: %#v\n", newSlice3, newSlice3)
	fmt.Printf("length: %d capacity: %d\n\n", len(newSlice3), cap(newSlice3))

}
