package main

import "fmt"

func ejercicio3() {
	yearB := 1999
	for {
		if yearB > 2022 {
			break
		}
		fmt.Println(yearB)
		yearB++
	}
}
