package main

import (
	"fmt"
)

func main() {

	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(sum(2, 3))
	//deferValues()

	//runtime.GOMAXPROCS(1)         // количество горутин одновременно
	//fmt.Println(runtime.NumCPU()) // узнать количество логических ядер (количество доступных горутин)

	//go showNumbers(100)

	//runtime.Gosched()          // переключиться на другую горутину
	//time.Sleep(time.Second) // пока спит одна, планировщик переключается самостоятельно на свободную
	fmt.Println("exit")

	makePanic()
}

func showNumbers(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
}

func sum(x, y int) (sum int) {
	defer func() {
		sum *= 2
	}()
	sum = x + y
	return
}

func deferValues() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("first", i)
	}
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println("second", i)
		}()
	}

	for i := 0; i < 10; i++ {
		k := i
		defer func() {
			fmt.Println("third", k)
		}()
	}
	for i := 0; i < 10; i++ {
		defer func(k int) {
			fmt.Println("fourth", k)
		}(i)
	}

}

func makePanic() {
	defer func() {
		panicValue := recover()
		fmt.Println(panicValue)
	}()

	panic("some panic text")
	fmt.Println("Unreachable code")
}
