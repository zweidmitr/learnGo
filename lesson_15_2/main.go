package main

import "fmt"

func main() {
	slices()
}
func slices() {
	var defaultSlice []int
	fmt.Printf("Type: %T Value: %#v\n", defaultSlice, defaultSlice)
	fmt.Printf("Length: %d Capacity: %d\n", len(defaultSlice), cap(defaultSlice))
	fmt.Println("<-------------->")

	stringSliceLiteral := []string{"first", "second"}
	fmt.Printf("Type: %T Value: %#v\n", stringSliceLiteral, stringSliceLiteral)
	fmt.Printf("Length: %d Capacity: %d\n", len(stringSliceLiteral), cap(stringSliceLiteral))
	fmt.Println("<-------------->")

	sliceByMake := make([]int, 3, 5)
	fmt.Printf("Type: %T Value: %#v\n", sliceByMake, sliceByMake)
	fmt.Printf("Length: %d Capacity: %d\n", len(sliceByMake), cap(sliceByMake))
	fmt.Println("<-------------->")

	sliceByMake = make([]int, 0, 5)
	fmt.Printf("Type: %T Value: %#v\n", sliceByMake, sliceByMake)
	fmt.Printf("Length: %d Capacity: %d\n", len(sliceByMake), cap(sliceByMake))
	fmt.Println("<-------------->")

	sliceByMake = append(sliceByMake, 1, 2, 3, 4, 5)
	fmt.Printf("Type: %T Value: %#v\n", sliceByMake, sliceByMake)
	fmt.Printf("Length: %d Capacity: %d\n", len(sliceByMake), cap(sliceByMake))
	fmt.Println("<-------------->")

	sliceByMake = append(sliceByMake, 6)
	fmt.Printf("Type: %T Value: %#v\n", sliceByMake, sliceByMake)
	fmt.Printf("Length: %d Capacity: %d\n", len(sliceByMake), cap(sliceByMake))
	fmt.Println("<-------------->")

	for ind, value := range sliceByMake {
		fmt.Printf("index: %d value: %d\n", ind, value)
	}
}
