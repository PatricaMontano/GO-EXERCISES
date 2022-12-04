package main

import (
	"fmt"
	"testing"
)

func BenchmarkRuneCount(b testing.B) {
	t := sumaPares(suma, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...)
	fmt.Println(t)
}
func BenchmarkRuneCount2(b testing.B) {
	t2 := sumaPares(suma, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...)
	fmt.Println(t2)
}
