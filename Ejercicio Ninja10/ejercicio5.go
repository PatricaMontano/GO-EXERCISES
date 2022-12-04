package main

import "fmt"

/*
Demuestra el uso del idioma coma ok con este código.
https://go.dev/play/p/YHOMV9NYKK
*/
func ejercicio5() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
