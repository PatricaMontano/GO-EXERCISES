package main

import (
	"ejercicio3/mymath"
	"fmt"
)

/*
Comienza con este código https://github.com/GoesToEleven/go-programming/tree/master/code_samples/010-ninja-level-thirteen/03/01-code-starting
Obtén el código listo para BET (agrega benchmarks, examples, tests).
Escribe una tabla de pruebas. Agrega documentación al código.
Corre lo siguiente en ese orden:
	tests
	benchmarks
	coverage
	coverage mostrado en el navegador
	examples mostrados en la documentación en el navegador
*/

func main() {
	xxi := gen()
	for _, v := range xxi {
		fmt.Println(mymath.CenteredAvg(v))
	}
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}
