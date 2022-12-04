package main

import "fmt"

func foo(c, salir chan int) {
	x := 2
	for {
		select {
		case c <- x:
			x += x
			//fmt.Println("Value x = ", x)
		case <-salir:
			fmt.Println("Finalizando.")
			return
		}
	}
}

func main() {
	c := make(chan int)
	salir := make(chan int)

	go func() {
		for i := 0; i < 11; i++ {
			fmt.Println(<-c)
		}
		salir <- 0
		//close(salir)
	}()
	foo(c, salir)
}
