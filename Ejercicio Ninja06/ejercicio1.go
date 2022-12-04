package main

import "fmt"

/*
crea una func con el identificador foo que retorne un int
crea una func con el identificador bar que retorne un int y un string
Llama ambas funciones
Imprime sus resultados

*/
func ejercicio1() {
	fmt.Println(foo())
	a, b := bar()
	fmt.Println(a, b)
}

func foo() int {
	return 2 * 5
}

func bar() (int, string) {
	a := 100
	b := "Hola mundo"
	return a, b
}
