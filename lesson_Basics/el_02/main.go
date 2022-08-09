package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	one()
	fmt.Println("<------>")
	two()
	fmt.Println("<------>")
	three()

}

func one() {
	// i передается в общем скоупе, следовательно, когда горутины будут выполняться,
	// цикл уже закончится и i будет равно 5 +/-
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(100 * time.Millisecond)
}

func two() {
	// В данном случае нужно передать копию i
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
}

func three() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			//wg.Add(1)
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
