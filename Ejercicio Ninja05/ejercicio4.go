package main

import "fmt"

/*
Crea y usa un struct anónimo.
*/
func ejercicio4() {
	a1 := struct {
		nombre    string
		edad      int
		ciudad    string
		formacion []string
	}{
		nombre: "Omar",
		edad:   22,
		ciudad: "Loja",
		formacion: []string{
			"MQB",
			"NSR",
			"UNL",
		},
	}
	fmt.Println(a1)
	fmt.Println(a1.nombre, a1.edad, a1.ciudad)
	for i, v := range a1.formacion {
		fmt.Println("\t", i, " ", v)
	}
}
