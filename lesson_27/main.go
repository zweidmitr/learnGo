package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	_ = sync.WaitGroup{}
	fmt.Println()
	_, c := context.WithTimeout(context.Background(), time.Millisecond*20)
	_ = c
	print("today is a GOod day")
}
