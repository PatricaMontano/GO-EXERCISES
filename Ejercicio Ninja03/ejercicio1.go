package main

import "fmt"

/*
  Imprime todos los números del 1 al 10,000
*/
func ejercicio1() {
	print(0, 100)
	print2(100)
}

func print(initial int, longitud int) {
	if initial <= longitud {
		fmt.Println(initial)
		print(initial+1, longitud)
	}

}

func print2(longitud int) {
	if longitud > -1 {
		print2(longitud - 1)
		fmt.Println(longitud)
	}

}
