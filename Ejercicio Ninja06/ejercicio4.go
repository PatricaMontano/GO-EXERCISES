package main

import "fmt"

/*
Crea un struct con
	El identificador “persona”
	Los campos:
		Nombre
		Apellido
		Edad
Adjunta un método al tipo persona con
	El identificador “presentar”
	El método debe hacer que el tipo persona diga su nombre y edad
Crea un valor de tipo persona
Llama al método usando el valor tipo persona
*/
type persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func (p persona) presentar() {
	fmt.Println("Mi nombre es", p.Nombre, p.Apellido, "tengo", p.Edad, "años")
}

func ejercicio4() {
	p1 := persona{
		Nombre:   "Omar",
		Apellido: "Sanmartin",
		Edad:     22,
	}
	p1.presentar()
}
