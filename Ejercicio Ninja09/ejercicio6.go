package main

import (
	"fmt"
	"runtime"
)

/*
Crea un programa que imprima tu OS y ARCH. Usa los siguientes comandos para correrlo
go run
go build
go install
*/
func ejercicio6() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
