package main

import "fmt"

/*
Crea tu propio tipo “persona” el cual tendrá un tipo subyacente tipo “struct”
de manera que pueda almacenar la siguiente data:
Nombre
Apellido
Sabores de helado favoritos
Crea dos VALORES de TIPO persona. Imprime los valores, usa range sobre los
elementos en el slice el cual almacena los valores de helados favoritos.

*/

type persona struct {
	Nombre            string
	Apellido          string
	helados_favoritos []string
}

func ejercicio1() {
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
	fmt.Println(p1.Nombre)
	fmt.Println(p1.Apellido)
	for _, v := range p1.helados_favoritos {
		fmt.Println("\t", v)
	}
	fmt.Println("---------------------------")
	fmt.Println(p2.Nombre)
	fmt.Println(p2.Apellido)
	for _, v := range p2.helados_favoritos {
		fmt.Println("\t", v)
	}
}
