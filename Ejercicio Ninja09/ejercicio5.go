package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
Arreglar la race condition que creaste en el ejercicio #4 usando el paquete atomic
*/
func ejercicio5() {
	fmt.Println("Número de CPUs:", runtime.NumCPU())
	fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	var contador int64
	const incremento = 50

	wg.Add(incremento)

	for i := 0; i < incremento; i++ {
		go func() {
			defer wg.Done()

			atomic.AddInt64(&contador, 1)
			//fmt.Println(atomic.LoadInt64(&contador))

		}()
		fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("El contador final es", contador)
}
