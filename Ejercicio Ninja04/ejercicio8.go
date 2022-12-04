package main

import "fmt"

/*
Crea un mapa con una llave TIPO string el cual representa el “nombre_apellido”
de una persona y un valor de TIPO []string el cual almacena sus cosas favoritas.
Almacena tres registros en tu mapa. Imprime todos sus valores con su índice de
posición en el slice.

`eduar_tua`, `computadoras`, `montaña`, `playa`
`carlos_ramon`, `leer`, `comprar`, `música`
`juan_bimba`, `helado`, `pintar`, `bailar`

*/
func ejercicio8() {
	x := map[string][]string{
		`eduar_tua`:    []string{`computadoras`, `montaña`, `playa`},
		`carlos_ramon`: []string{`leer`, `comprar`, `música`},
		`juan_bimba`:   []string{`helado`, `pintar`, `bailar`},
	}

	for clave, valor := range x {
		fmt.Println("Persona :", clave)
		for i := 0; i < len(valor); i++ {
			fmt.Println("	Index :", i, "   Valor:", valor[i])
		}
	}
}
