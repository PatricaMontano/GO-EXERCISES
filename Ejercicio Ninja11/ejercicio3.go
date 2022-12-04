package main

import "fmt"

/*
Crea un struct “errorPer” el cual implemente la interfaz builtin.error.
Crea una función “foo” que tenga un valor de tipo error como parámetro.
Crea un valor de tipo “errorPer” y pásalo a “foo”.
Si necesitas un pista, aquí hay una https://go.dev/play/p/ESmKdz-T3At
*/

type errorPer struct {
	Tipo        string
	Descripcion string
}

func (ep errorPer) Error() string {
	return fmt.Sprintf("Tipo: %v Descripcin: %v", ep.Tipo, ep.Descripcion)
}

func ejercicio3() {
	e1 := errorPer{"Fuera Rango", "El error se produjo por una posicion extra a la longitud"}
	foo(e1)
}

func foo(e error) {
	fmt.Println(e)
	fmt.Println(e.(errorPer).Tipo)
}
