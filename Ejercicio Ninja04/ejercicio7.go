package main

import "fmt"

/*
Crear un slice de slice de string ([][]string). Almacena los siguientes valores
en un slice multi-dimensional:

"Batman", "Jefe", "Vestido de negro"
"Robin", "Ayudante", "Vestido de colores"

Haz range sobre los registros, luego sobre los datos de cada registro.

*/
func ejercicio7() {
	x := [][]string{{"Batman", "Jefe", "Vestido de negro"}, {"Robin", "Ayudante", "Vestido de colores"}}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println("Registro:", i)
		for j, v2 := range v {
			fmt.Println("      Indice: ", j, "   Valor:", v2)
		}
	}
}
