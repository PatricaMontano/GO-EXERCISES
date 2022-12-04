package main

import (
	"fmt"
	"time"
)

func callback() {
	n := 123456
	x := make([]int, n)
	for i := 1; i < n; i++ {
		x[i] = i
	}

	start := time.Now()
	t := sumaPares(suma, x...)
	fmt.Println(t)
	end := time.Since(start)

	start3 := time.Now()
	t3 := sumaPares2(sumaRecursiva2, x...)
	fmt.Println(t3)
	end3 := time.Since(start3)

	fmt.Println("Tiempo de Ejecucion con For", end)
	fmt.Println("Tiempo de Ejecucion con Recursividad2", end3)

}

func suma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func sumaRecursiva2(longitud int, x ...int) int {

	if longitud > 2 {
		return x[0] + x[1] + x[2] + sumaRecursiva2(longitud-3, x[3:]...)
	}
	if longitud == 2 {
		return x[0] + x[1]
	}
	if longitud == 1 {
		return x[0]
	}
	return 0

}

func sumaPares(f func(x ...int) int, y ...int) int {
	var xi []int
	for _, v := range y {
		if v%2 == 0 {
			xi = append(xi, v)
		}
	}
	total := f(xi...)
	return total
}

func sumaPares2(f func(a int, x ...int) int, y ...int) int {
	var xi []int
	for _, v := range y {
		if v%2 == 0 {
			xi = append(xi, v)
		}
	}
	total := f(len(xi), xi...)
	return total
}
