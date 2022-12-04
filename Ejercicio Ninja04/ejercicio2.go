package main

import "fmt"

/*
Usando un COMPOSITE LITERAL:
Crea un SLICE de TIPO int
Asigna 10 VALORES
Haz Range sobre el slice e imprime los valores.
Usando format para imprimir
Imprime el TIPO del slice

*/

func ejercicio2() {
	x := []int{1, 4, 8, 12, 16, 20, 24, 28, 32, 36}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("El tipo es de %T", x)
}
