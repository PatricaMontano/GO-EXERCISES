package main

import "fmt"

/*
Crea un slice para almacenar los nombres de todos los estados en los Estados Unidos
de América. ¿Cuál es la longitud del slice? ¿Cuál es la capacidad del Slice?
Imprime todos los valores con su  índice de posición, sin utilizar la cláusula range.
Aquí la lista de los estados:

` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`,
` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`,
` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`,
` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`,
` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`,
` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`,
` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`,
` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`,
` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
*/

func ejercicio6() {
	estados := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println("La longitud del slice es de ", len(estados))
	fmt.Println("La capacidad del slice es de ", cap(estados))

	for i := 0; i < len(estados); i++ {
		fmt.Println("Posicion", i, " valor = ", estados[i])
	}
}
