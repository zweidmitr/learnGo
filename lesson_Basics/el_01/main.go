package main

import "fmt"

func main() {
	fmt.Println(MaxSum([]int{1, 2, 3}, []int{10, 20, 50}))
	fmt.Println(MaxSum([]int{3, 2, 1}, []int{1, 2, 3}))
}

func MaxSum(num1, num2 []int) []int {
	if sum(num1) >= sum(num2) {
		return num1
	}
	return num2
}
func sum(nums []int) int {
	temp := 0
	for _, val := range nums {
		temp += val
	}
	return temp
}
