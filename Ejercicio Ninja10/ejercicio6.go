package main

import "fmt"

/*
Escribe un programa que:
Ponga 100 números en un canal
Extraiga los números del canal y los imprima

*/
func ejercicio6() {

	c := make(chan int)
	ponerNumeros(c)
	extraerImprimir(c)
	fmt.Println("Ejecutado con Exito")
}

func ponerNumeros(c chan int) {
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

}
func extraerImprimir(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
