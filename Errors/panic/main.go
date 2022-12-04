package main

import (
	"fmt"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		//		fmt.Println("un error ocurrió", err)
		//		log.Println("un error ocurrió", err)
		//		log.Fatalln(err)
		//		log.Panicln(err)
		panic(err)
	}
}

func foo() {
	fmt.Println("Hola Mundo")
}

// http://godoc.org/builtin#panic
