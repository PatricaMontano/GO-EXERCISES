package main

import (
	"fmt"
	"time"
)

func recursividad() {

	start := time.Now()
	n := factorial(25)
	fmt.Println(n)
	end := time.Since(start)

	start2 := time.Now()
	n2 := loopFact(25)
	fmt.Println(n2)
	end2 := time.Since(start2)

	fmt.Println("Tiempo de Ejecucion con Recursividad", end)
	fmt.Println("Tiempo de Ejecucion con For", end2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
