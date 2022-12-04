package main

import (
	"fmt"
	"sync"
)

/*
Además de la gorutina principal, lanza dos gorutinas adicionales

	Cada gorutina debe imprimir algo en pantalla

Usa waitgroups para asegurarte que cada gorutina finalize antes de
que el programa termine
*/
func ejercicio1() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Print Go Rutina 1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Print Go Rutina 2")
	}()

	wg.Wait()
	fmt.Println("Fin Ejecucion")
}
