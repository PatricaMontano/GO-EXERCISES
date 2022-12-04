package main

import "fmt"

/*
Haz que este código funcione: https://go.dev/play/p/j-EA6003P0
Usando una función literal, también conocida como,
función anónima autoejecutable

*/
func ejercicio1() {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
