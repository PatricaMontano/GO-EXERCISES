package main

import "fmt"

/*
Usa el código del ejemplo anterior y almacena los valores de tipo persona
en un mapa con la llave apellido. Accede a cada valor en el mapa. Imprime
los valores. Imprime también los valores haciendo range sobre el slice.
*/

func ejercicio2() {
	p1 := persona{
		Nombre:            "Omar",
		Apellido:          "Sanmartin",
		helados_favoritos: []string{"Oreo", "Mango", "Cafe"},
	}
	p2 := persona{
		Nombre:            "Domenica",
		Apellido:          "Vivanco",
		helados_favoritos: []string{"Vainilla", "Chocolate", "Cereza"},
	}

	x := map[string]persona{
		p1.Apellido: p1,
		p2.Apellido: p2,
	}

	for _, valPersona := range x {
		fmt.Println(valPersona.Nombre, valPersona.Apellido)
		for _, valueHelados := range valPersona.helados_favoritos {
			fmt.Println("\t", valueHelados)
		}
	}

}
