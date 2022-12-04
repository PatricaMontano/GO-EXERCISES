package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 42
	//b := strconv.Itoa(a)
	c := strconv.FormatInt(int64(a), 10)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
