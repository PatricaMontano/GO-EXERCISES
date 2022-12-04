package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Arregla la race condition que creaste en el ejercicio anterior usando un mutex

	Tiene sentido eliminar runtime.Gosched()
*/
func ejercicio4() {
	fmt.Println("Número de CPUs:", runtime.NumCPU())
	fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())
	var wg sync.WaitGroup
	contador := 0
	const incremento = 50

	wg.Add(incremento)
	var m sync.Mutex

	for i := 0; i < incremento; i++ {
		go func() {
			defer wg.Done()

			m.Lock()
			aux := contador
			aux++
			contador = aux
			m.Unlock()
		}()
		fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("El contador final es", contador)
}
