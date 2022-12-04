package main

import "fmt"

/*
Asigna una función  a una variable, luego llama esa función
*/
func ejercicio7() {
	f := func(text string) {
		fmt.Println(text)
	}
	f("Ejercicio 7: Hello World")
}
