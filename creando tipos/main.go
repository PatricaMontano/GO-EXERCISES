package main

import "fmt"

type dinero int

func main() {
	a := 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	var b dinero
	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
