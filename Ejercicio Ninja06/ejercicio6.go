package main

import "fmt"

/*
Crea y usa una func anónima
*/
func ejercicio6() {
	func(text string) {
		fmt.Println(text)
	}("Hola Mundo funcion anonima")
}
