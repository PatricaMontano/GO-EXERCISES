package main

import "fmt"

const (
	a = iota
	b
	c
	d
)

func main() {
	s1 := "Hola mundo"

	for i := 0; i < len(s1); i++ {
		fmt.Printf("%#U ", s1[i])
	}
	fmt.Println("")
	for i, v := range s1 {
		fmt.Printf("El indices es %d el valor es %#x\n", i, v)
	}
	fmt.Println(a, b, c, d)
}
