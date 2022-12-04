package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0

	c := make(chan int)

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			c <- 1
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	for i := 0; i < gs; i++ {
		counter += <-c
	}

	wg.Wait()
	close(c)
	fmt.Println("Total Count:", counter)
}
