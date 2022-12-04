package main

import "fmt"

/*
Comenzando con este código https://go.dev/play/p/Odb27oVYQ4D
extrae los valores del canal usando un ciclo for range
*/
func ejercicio3() {
	c := gen()
	recibir(c)

	fmt.Println("A punto de finalizar.")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
func recibir(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
