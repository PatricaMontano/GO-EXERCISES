package main

import "fmt"

/*
Crea un tipo struct persona
Adjunta el método hablar usando un receptor de tipo puntero
	*persona
Crea un tipo interfaz humano
	Para implícitamente implementar esa interfaz, un tipo humano debe tener el método hablar
Crea la función “diAlgo”
	Haz que tome un humano como parámetro
	Haz que llame al método hablar
Muestra lo siguiente en tu código
	PUEDES pasar un valor de tipo *persona a diAlgo
	NO puedes pasar un valor de tipo persona a diAlgo
Aquí hay una pista si necesitas ayuda
	https://play.golang.org/p/JQd6LsU_L-K
*/

type persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func (p *persona) hablar() {
	fmt.Println("Mi nombre es", p.Nombre, p.Apellido, "tengo", p.Edad, "años")
}

type humano interface {
	hablar()
}

func diAlgo(h humano) {
	h.hablar()
}

func ejercicio2() {
	p1 := persona{"Omar", "Sanmartin", 22}
	diAlgo(&p1)
	//diAlgo(p1)
}
