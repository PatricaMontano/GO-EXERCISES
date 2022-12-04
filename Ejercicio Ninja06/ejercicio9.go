package main

import "fmt"

/*
Un “callback” es cuando le pasamos una función a otra función como argumento.
Para este ejercicio,
	Pasa una func a otra función como argumento
*/
func ejercicio9() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m, s := operacion(multiplicar, sumar, x...)
	fmt.Println("La multiplicacion es", m, "la suma es", s)
}

func multiplicar(x ...int) int {
	mul := 1
	for _, v := range x {
		mul *= v
	}
	return mul
}

func sumar(x ...int) int {
	sum := 1
	for _, v := range x {
		sum += v
	}
	return sum
}

func operacion(m func(x ...int) int, s func(y ...int) int, val ...int) (int, int) {
	return m(val...), s(val...)
}
