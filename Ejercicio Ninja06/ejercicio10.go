package main

import "fmt"

/*
Closure es cuando “encerramos” el scope de una variable en un bloque de código.
Para este ejercicio, crea una func el cual “encierra” el scope de una variable:
*/

func ejercicio10() {
	a := scopeFunc()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println("*******************")

	b := scopeFunc()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func scopeFunc() func() int {
	x := 100
	return func() int {
		x--
		return x
	}
}
