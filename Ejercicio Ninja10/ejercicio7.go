package main

import "fmt"

/*
Escribe un programa que:
	Lance 10 gorutinas
		Cada gorutina agrega 10 números a un canal
	Extrae los números del canal e imprímelos
*/
func ejercicio7() {
	c := make(chan int)
	crearRotinas(c)
	for k := 0; k < 100; k++ {
		fmt.Println(<-c)
	}
	fmt.Println("Programa Terminado")
}

func crearRotinas(c chan<- int) {
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
		}()
	}
}
