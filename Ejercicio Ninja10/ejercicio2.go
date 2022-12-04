package main

import "fmt"

/*
Haz que este código funcione:
	1. https://play.golang.org/p/oB-p3KMiH6

	2. https://play.golang.org/p/_DBRueImEq
*/
func ejercicio2() {
	primerProblema()
	segundoProblema()
}

func primerProblema() {
	//1. https://play.golang.org/p/oB-p3KMiH6
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
func segundoProblema() {
	//2. https://play.golang.org/p/_DBRueImEq
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
