package main

import "fmt"

func main() {
	numsCh := make(chan []int)
	sumCh := make(chan int)

	go SumWorker(numsCh, sumCh)
	numsCh <- []int{10, 20, 30}

	res := <-sumCh // 60
	fmt.Println(res)

}

func SumWorker(numCh chan []int, sumCh chan int) {
	for nums := range numCh {
		sumCh <- sum(nums)
	}
}

func sum(nums []int) int {
	temp := 0
	for _, val := range nums {
		temp += val
	}
	return temp
}
