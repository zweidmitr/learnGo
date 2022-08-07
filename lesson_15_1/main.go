package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	arrays()
}

func arrays() {
	var intArr [3]int
	fmt.Printf("Type: %T Value: %#v\n", intArr, intArr)

	intArr[0] = 5
	intArr[1] = 6
	fmt.Printf("Type: %T Value: %#v\n", intArr, intArr)

	people := [2]Person{
		{
			Age:  30,
			Name: "Katy",
		},
		{
			Age:  23,
			Name: "John",
		},
	}
	fmt.Printf("Type: %T Value: %#v\n", people, people)

	stringsArr := [...]string{"first", "second", "third", "fourth"}
	fmt.Printf("Type: %T Value: %#v\n", stringsArr, stringsArr)

	fmt.Printf("Length: %d Capacity: %d\n", len(stringsArr), cap(stringsArr))

	for i := 0; i < len(stringsArr); i++ {
		fmt.Printf("Index: %d Value: %s\n", i, stringsArr[i])
	}
	fmt.Println("<-------------->")

	for inx, value := range stringsArr {
		fmt.Printf("Index: %d Value: %s\n", inx, value)
	}
	fmt.Println("<-------------->")

	for _, num := range intArr {
		fmt.Printf("Value: %d\n", num)
	}
	fmt.Println("<-------------->")

	newIntArr := changeArray(intArr)
	fmt.Printf("Type: %T Value: %#v\n", intArr, intArr)
	fmt.Printf("Type: %T Value: %#v\n", newIntArr, newIntArr)

}
func changeArray(arr [3]int) [3]int {
	arr[2] = 3
	return arr
}
