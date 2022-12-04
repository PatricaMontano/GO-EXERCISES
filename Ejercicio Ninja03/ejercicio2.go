package main

import (
	"fmt"
	"time"
)

/*
 Imprime el rune code point de las letras del
 alfabeto en mayúsculas tres veces.
*/

func ejercicio2() {
	start := time.Now()
	iteracion()
	end := time.Since(start)

	start2 := time.Now()
	imprimirCode(65, 91)
	end2 := time.Since(start2)

	fmt.Println("Time Elapsed with For ", end)
	fmt.Println("Time Elapsed with Recursion ", end2)

}

func iteracion() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

func imprimirCode(begin int, end int) {
	if begin < end {
		fmt.Println(begin)
		fmt.Printf("\t%#U\n", begin)
		fmt.Printf("\t%#U\n", begin)
		fmt.Printf("\t%#U\n", begin)
		imprimirCode(begin+1, end)
	}

}
