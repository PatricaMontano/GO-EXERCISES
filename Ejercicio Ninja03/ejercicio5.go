package main

import "fmt"

/*
Imprime el resto o módulo, el cual es resultado de
dividir entre 4 cada número en el rango de 10 y 100.
*/
func ejercicio5() {
	//printModulo()
	printRecursivo(10)
}

func printModulo() {
	for i := 10; i < 101; i++ {
		fmt.Printf("El resto o mudulo de %v dividido entre 4 es %v \n", i, i%4)
	}
}

func printRecursivo(begin int) {
	if begin < 101 {
		fmt.Printf("El resto o mudulo de %v dividido entre 4 es %v \n", begin, begin%4)
		printRecursivo(begin + 1)
	}
}
