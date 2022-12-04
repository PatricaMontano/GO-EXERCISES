package main

import "fmt"

/*
Para BORRAR de un slice, usamos APPEND en conjunto con SLICING(dividir).
Para este ejercicio sigue los siguientes pasos:
Comienza con un slice
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
Usa APPEND & SLICING para obtener los siguientes valores el cual se los debes asignar
a una variable “y” y luego imprimir:
[42, 43, 44, 48, 49, 50, 51]

*/
func ejercicio5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := append(x[:3], x[6:]...)
	fmt.Println(y)
}
