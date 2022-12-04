package main

import (
	"fmt"
	"math"
)

/*
Crea un tipo CUADRADO
Crea un tipo CÍRCULO
Adjunta un método a cada uno que calcule y retorne el ÁREA
	Área de un círculo= π r 2
	Área de un cuadrado = L * H
Crea un tipo FORMA que defina una interfaz como cualquier cosa que tenga el método ÁREA
Crea una func INFO el cuál toma un tipo forma e imprime el área de la misma.
Crea un valor de tipo cuadrado
Crea un valor de tipo círculo
Usa la func info para imprimir el área de un cuadrado
Usa la func info para imprimir el área de un círculo

*/

type cuadrado struct {
	lado float32
}

type circulo struct {
	radio float32
}

func (c cuadrado) area() float32 {
	return c.lado * c.lado
}

func (c circulo) area() float32 {
	return math.Pi * c.radio * c.radio
}

type forma interface {
	area() float32
}

func ejercicio5() {
	cu1 := cuadrado{
		lado: 5.75,
	}
	ci1 := circulo{
		radio: 6,
	}
	info(cu1)
	info(ci1)
}

func info(f forma) {
	switch f.(type) {
	case cuadrado:
		fmt.Println("Cuadrado")
	case circulo:
		fmt.Println("Circulo")
	}
	fmt.Println("El area es", f.area())
}
