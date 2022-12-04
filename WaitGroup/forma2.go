package main

import (
	"fmt"
	"runtime"
	"sync"
)

func forma2() {
	wg := sync.WaitGroup{}

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo2(&wg)
	bar2()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Progan Exited")
}

func foo2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
}

func bar2() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
