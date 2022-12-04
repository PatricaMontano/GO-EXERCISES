package main

import (
	"fmt"
)

/*
Usa la palabra clave “defer” para mostrar que una función
diferida corre después que la función que la contiene
finaliza o retorna.
*/
func ejercicio3() {
	defer printHello()
	fmt.Println("Main ejecucion Segunda linea")
}

func printHello() {
	defer func(text string) {
		fmt.Println(text)
	}("Hola mundo -- Defer")

	fmt.Println("Segunda linea - PrintHello")
}
