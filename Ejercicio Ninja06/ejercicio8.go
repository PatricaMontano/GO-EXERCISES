package main

import "fmt"

/*
Crea una func el cual retorna una func
Asigna la func retornada a una variable
Llama la func retornada
*/
func ejercicio8() {
	f := peep()
	f()
}

func peep() func() {
	return func() {
		fmt.Println("Print Funcion Retorno")
	}
}
