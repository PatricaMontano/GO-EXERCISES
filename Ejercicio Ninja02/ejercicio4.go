package main

import "fmt"

func main() {
	var a int = 5
	fmt.Printf("%d \t %b \t %#x \n", a, a, a)
	b := a << 1
	fmt.Printf("%d \t %b \t %#x \n", b, b, b)

}
