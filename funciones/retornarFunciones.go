package main

import (
	"fmt"
)

func retornarFunciones() {
	a, b := bar()
	a()
	b(1516)

	bar2()(5000)

}

func bar() (func(), func(int)) {
	f := func() {
		fmt.Println("Mi primera expresión función")
	}

	g := func(x int) {
		fmt.Println("El año del descubrimiento de América fue:", x)
	}
	return f, g
}

func bar2() func(int) {
	return func(x int) {
		fmt.Println("El resultado es ", x)
	}
}
