package main

import (
	"fmt"
	"sort"
)

/*
Comenzando con este código, ordena el []int y []string.
*/
func ejercicio4() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"Una", "persona", "que", "nunca", "ha", "cometido", "un", "error", "nunca", "intenta", "nada", "nuevo"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)
}
