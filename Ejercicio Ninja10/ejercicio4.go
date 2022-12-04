package main

import "fmt"

/*
Comenzando con este código https://go.dev/play/p/9rDOEvNeBjc
extrae los valores del canal usando una declaración select
*/
func ejercicio4() {
	salir := make(chan int)
	c := gen2(salir)

	recibir2(c, salir)

	fmt.Println("A punto de finalizar.")
}

func gen2(salir chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		salir <- 1
		close(c)
	}()

	return c
}

func recibir2(c, salir <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-salir:
			return
		}
	}
}
