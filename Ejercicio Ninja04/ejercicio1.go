package main

import "fmt"

/*
Usando un COMPOSITE LITERAL:
Crea un ARREGLO el cual tenga 5 VALORES de TIPO int
Asigna VALORES a cada posición dada por los índices.
Itera con Range sobre el arreglo e imprime los valores.
Usando el paquete fmt
Imprime el TIPO del arreglo

*/
func ejercicio1() {
	arreglo := [5]int{5, 10, 15, 20, 25}

	for i, v := range arreglo {
		fmt.Println("Posicion ", i, "con el valor", v)
	}

	fmt.Printf("El tipo del arreglo es %T", arreglo)
}
