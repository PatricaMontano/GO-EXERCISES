package main

import "fmt"

/*
Crea una función con el identificador foo que
	Tome un parámetro variable de tipo int
	Pásale un valor de tipo []int a la función (haz el despliegue del []int)
	Retorna la suma de todos los valores de tipo int pasados como argumento.
Crea una func con el identificador bar que
	Tome un parámetro de tipo []int
	Retorne la suma de todos los valores de tipo int pasados.
*/

func ejercicio2() {
	f := foo2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	b := bar2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	fmt.Println(f)
	fmt.Println(b)
}

func foo2(x ...int) int {
	suma := 0
	for _, v := range x {
		suma += v
	}
	return suma
}

func bar2(x []int) int {
	suma := 0
	for _, v := range x {
		suma += v
	}
	return suma
}
